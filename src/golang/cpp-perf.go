package main

import (
	"fmt"
	"golang-driver/cassandra"
	// "os"
	"time"
)

func printError(future *cassandra.Future) {
	// fmt.Fprintf(os.Stderr, "Error: %v %v\n", future.ErrorCode(), future.ErrorMessage())
}

func main() {
	iterations := 100
	concurrentRequests := uint(20000)
	reportInterval := 5
	numThreads := uint(1)

	bigString := "0123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567012345670123456701234567"
	uuidGen := cassandra.NewUuidGenerator()
	defer uuidGen.Finalize()

	futures := make([]*cassandra.Future, concurrentRequests)
	quit := make(chan bool)

	cassandra.SetLogLevel(cassandra.CASS_LOG_ERROR)
	cluster := cassandra.NewCluster()
	defer cluster.Finalize()

	cluster.SetContactPoints("127.0.0.1")
	cluster.SetQueueSizeIo(concurrentRequests)
	cluster.SetNumThreadsIo(numThreads)
	cluster.SetPendingRequestsLowWaterMark(5000)
	cluster.SetPendingRequestsHighWaterMark(10000)
	cluster.SetCoreConnectionsPerHost(1)
	cluster.SetMaxConnectionsPerHost(2)

	session := cassandra.NewSession()
	defer session.Finalize()

	sessfuture := cluster.SessionConnect(session)
	defer sessfuture.Finalize()
	sessfuture.Wait()
	if sessfuture.ErrorCode() != cassandra.CASS_OK {
		printError(sessfuture)
		return
	}

	go func(quit chan bool) {
		tick := time.NewTicker(time.Duration(reportInterval) * time.Second)
		for {
			select {
			case <-tick.C:
				metrics := session.Metrics()

				fmt.Printf("rate stats (requests/second): mean %v 1m %v 5m %v 10m %v connection_timeouts %v pending_request_timeouts %v request_timeouts %v\n",
					metrics.Requests.Mean,
					metrics.Requests.OneMinuteRate,
					metrics.Requests.FiveMinuteRate,
					metrics.Requests.FifteenMinuteRate,
					metrics.Errors.ConnectionTimeouts,
					metrics.Errors.PendingRequestTimeouts,
					metrics.Errors.RequestTimeouts)

			case <-quit:
				return
			}
		}
	}(quit)

	for z := 0; z < iterations; z++ {

		for i := uint(0); i < concurrentRequests; i++ {
			uuid := uuidGen.GenRandom()
			statement := cassandra.NewStatement("INSERT INTO examples.songs (id, title, album, artist) VALUES (?, ?, ?, ?);", 4)
			statement.Bind(uuid, bigString, bigString, bigString)
			defer statement.Finalize()
			futures[i] = session.Execute(statement)
		}

		for i := uint(0); i < concurrentRequests; i++ {
			future := futures[i]
			future.Wait()
			if future.ErrorCode() != cassandra.CASS_OK {
				printError(future)
			}
			result := future.Result()
			defer result.Finalize()
			defer future.Finalize()
		}
	}

	close(quit)
	metrics := session.Metrics()
	fmt.Printf("final stats (microseconds): min %llu max %llu median %llu 75th %llu 95th %llu 98th %llu 99th %llu 99.9th %llu\n",
		metrics.Requests.Min,
		metrics.Requests.Max,
		metrics.Requests.Median,
		metrics.Requests.Percentile75th,
		metrics.Requests.Percentile95th,
		metrics.Requests.Percentile98th,
		metrics.Requests.Percentile99th,
		metrics.Requests.Percentile999th)

	fmt.Printf("DONE.\r\n")
}
