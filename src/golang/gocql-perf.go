package main

import (
	"fmt"
	"github.com/bmizerany/perks/quantile"
	"github.com/gocql/gocql"
	"os"
	"runtime"
	"sync"
	"time"
)

type Result struct {
	start    time.Time
	duration time.Duration
	err      error
}

func (this *Result) Reset() {
	this.start = time.Unix(0, 0)
	this.duration = 0
	this.err = nil
}

type Metrics struct {
	startTime time.Time
	total     int64
	success   int64
	failed    int64
	quantiles *quantile.Stream
}

func NewMetrics() *Metrics {
	return &Metrics{
		startTime: time.Now(),
		total:     0,
		success:   0,
		failed:    0,
		quantiles: quantile.NewTargeted(0.50, 0.95, 0.99),
	}
}

func (this *Metrics) AddResult(result *Result) {
	this.total++
	if result.err != nil {
		this.failed++
	} else {
		this.success++
	}
	this.quantiles.Insert(float64(result.duration.Nanoseconds() / int64(time.Microsecond)))
}

func (this *Metrics) P50() float64 {
	return this.quantiles.Query(0.50)
}

func (this *Metrics) P95() float64 {
	return this.quantiles.Query(0.95)
}

func (this *Metrics) P99() float64 {
	return this.quantiles.Query(0.99)
}

func (this *Metrics) ResetQuantiles() {
	this.quantiles.Reset()
}

func metricsUpdater(metrics *Metrics, quit chan bool, results chan Result) {
	tick := time.NewTicker(time.Duration(5) * time.Second)

	for {
		select {
		case <-tick.C:
			printStatus(metrics)
			metrics.ResetQuantiles()
		case <-quit:
			return
		case r := <-results:
			metrics.AddResult(&r)
		}
	}
}

func printStatus(metrics *Metrics) {
	successRate := 0.0
	failRate := 0.0

	if metrics.total != 0 {
		successRate = float64(metrics.success / metrics.total)
		failRate = float64(metrics.failed / metrics.total)
	}

	now := time.Now()

	fmt.Printf("%s | seconds : %08.0f | total : %010d | succes : %1.2f | failure : %1.2f | P50 : %2.2f µs | P95 : %2.2f µs | P99 : %2.2f µs\n",
		now.Format("02 Jan 2006 15:04:05"),
		now.Sub(metrics.startTime).Seconds(),
		metrics.total,
		successRate,
		failRate,
		metrics.P50(),
		metrics.P95(),
		metrics.P99())
}

func client(iterations int, quit chan bool, results chan Result, done *sync.WaitGroup) {
	bigString := "0123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567"

	defer func() {
		done.Done()
	}()

	cluster := gocql.NewCluster("127.0.0.1")
	session, _ := cluster.CreateSession()
	defer session.Close()

	for z := 0; z < iterations; z++ {
		var result Result
		result.Reset()

		uuid, _ := gocql.RandomUUID()
		startTime := time.Now()
		result.err = session.Query("INSERT INTO examples.songs (id, title, album, artist) VALUES (?, ?, ?, ?);", uuid, bigString, bigString, bigString).Exec()
		result.duration = time.Now().Sub(startTime)
		result.start = startTime
		results <- result
	}
}

func main() {
	concurrentRequests := 2000000
	numberOfClients := 5
	requestsPerClient := concurrentRequests / numberOfClients

	// set the max number of process
	goMaxProcs := os.Getenv("GOMAXPROCS")
	if goMaxProcs == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	metrics := NewMetrics()
	quit := make(chan bool)

	// spawn the metric updater
	results := make(chan Result, 100000)
	go metricsUpdater(metrics, quit, results)

	// spawn the clients
	var clientGroup sync.WaitGroup
	clientGroup.Add(numberOfClients)
	for i := 0; i < numberOfClients; i++ {
		go client(requestsPerClient, quit, results, &clientGroup)
	}

	clientGroup.Wait()
	printStatus(metrics)
}
