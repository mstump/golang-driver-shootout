## Schema
```
  CREATE KEYSPACE examples WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

  CREATE TABLE examples.songs (
        id uuid PRIMARY KEY,
        title text,
        album text,
        artist text,
        tags set<text>
    );
```

## Results

### C++
Modified to log at ERROR and not insert into the set.

```
#define NUM_THREADS 1
#define NUM_IO_WORKER_THREADS 1
#define NUM_CONCURRENT_REQUESTS 20000
#define NUM_ITERATIONS 100

#define DO_SELECTS 0
#define USE_PREPARED 1
```

```
rate stats (requests/second): mean 16551.712695 1m 16642.200000 5m 16642.200000 10m 16642.200000 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 17538.921087 1m 16749.260529 5m 16664.331723 10m 16649.618264 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 17969.699572 1m 16952.934547 5m 16707.839291 10m 16664.282945 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 17980.691545 1m 17050.549559 5m 16732.069490 10m 16672.645899 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 17881.257925 1m 17085.190352 5m 16744.494503 10m 16677.139816 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 17663.567314 1m 16995.401001 5m 16731.564312 10m 16673.178940 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 17477.250162 1m 16983.567494 5m 16733.478908 10m 16674.144152 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 17351.811528 1m 16958.272146 5m 16732.383414 10m 16674.105680 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 17147.957447 1m 16837.645382 5m 16711.180869 10m 16667.321732 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 17018.114508 1m 16757.462292 5m 16696.695567 10m 16662.709435 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 16986.793147 1m 16727.154145 5m 16691.434604 10m 16661.134318 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 16923.720388 1m 16700.964361 5m 16686.611002 10m 16659.685378 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 16762.171107 1m 16583.544438 5m 16662.575009 10m 16651.778000 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 16558.685887 1m 16359.929099 5m 16615.042629 10m 16635.905586 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 16383.540289 1m 16148.308325 5m 16567.079389 10m 16619.713353 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 16123.749962 1m 15850.033348 5m 16498.497833 10m 16596.434115 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 15993.935014 1m 15707.965598 5m 16458.411185 10m 16582.455014 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
final stats (microseconds): min 246 max 180095 median 8367 75th 12359 95th 22623 98th 34463 99th 59871 99.9th 116735
```

### Java

```
/Users/mstump/dse/resources/cassandra/tools/bin/cassandra-stress user profile=/Users/mstump/src/golang-driver-shootout/stress.yaml ops\(insert=1\) -rate threads=1
```

```
Connected to cluster: Test Cluster
Datatacenter: Solr; Host: localhost/127.0.0.1; Rack: rack1
Created schema. Sleeping 1s for propagation.
Sleeping 2s...
Warming up insert with 50000 iterations...
Generating batches with [1..1] partitions and [1..1] rows (of [1..1] total rows in the partitions)
Running [insert] with 1 threads until stderr of mean < 0.02
type,      total ops,    op/s,    pk/s,   row/s,    mean,     med,     .95,     .99,    .999,     max,   time,   stderr, errors,  gc: #,  max ms,  sum ms,  sdv ms,      mb
total,          3416,    3416,    3416,    3416,     0.3,     0.3,     0.4,     0.5,     1.5,     6.7,    1.0,  0.00000,      0,      0,       0,       0,       0,       0
total,          6940,    3516,    3516,    3516,     0.3,     0.3,     0.3,     0.4,     1.4,     7.4,    2.0,  0.01046,      0,      0,       0,       0,       0,       0
total,         10456,    3496,    3496,    3496,     0.3,     0.3,     0.3,     0.4,     1.3,     7.3,    3.0,  0.00737,      0,      0,       0,       0,       0,       0
total,         13934,    3471,    3471,    3471,     0.3,     0.3,     0.3,     0.4,     1.3,     6.7,    4.0,  0.00555,      0,      0,       0,       0,       0,       0
total,         17094,    3152,    3152,    3152,     0.3,     0.3,     0.3,     0.4,     1.4,    98.2,    5.0,  0.00445,      0,      1,      96,      96,       0,     291
total,         20558,    3456,    3456,    3456,     0.3,     0.3,     0.3,     0.6,     1.5,     1.6,    6.0,  0.00405,      0,      0,       0,       0,       0,       0
total,         24068,    3494,    3494,    3494,     0.3,     0.3,     0.3,     0.4,     1.2,     1.2,    7.0,  0.00348,      0,      0,       0,       0,       0,       0
total,         27558,    3480,    3480,    3480,     0.3,     0.3,     0.3,     0.4,     1.2,     1.5,    8.0,  0.00305,      0,      0,       0,       0,       0,       0
total,         31115,    3547,    3547,    3547,     0.3,     0.3,     0.3,     0.4,     1.3,     1.6,    9.0,  0.00326,      0,      0,       0,       0,       0,       0
total,         34300,    3167,    3167,    3167,     0.3,     0.3,     0.3,     0.5,     1.6,    89.6,   10.0,  0.00300,      0,      1,      87,      87,       0,     298
total,         37724,    3414,    3414,    3414,     0.3,     0.3,     0.3,     0.5,     1.8,     4.3,   11.0,  0.00321,      0,      0,       0,       0,       0,       0
total,         41038,    3286,    3286,    3286,     0.3,     0.3,     0.4,     0.7,     1.9,     5.2,   12.0,  0.00522,      0,      0,       0,       0,       0,       0
total,         44466,    3417,    3417,    3417,     0.3,     0.3,     0.4,     0.5,     1.5,     1.9,   13.0,  0.00495,      0,      0,       0,       0,       0,       0
total,         47981,    3500,    3500,    3500,     0.3,     0.3,     0.3,     0.4,     1.3,     1.4,   14.0,  0.00464,      0,      0,       0,       0,       0,       0
total,         51509,    3516,    3516,    3516,     0.3,     0.3,     0.3,     0.6,     1.0,     1.8,   15.1,  0.00442,      0,      0,       0,       0,       0,       0
total,         54819,    3298,    3298,    3298,     0.3,     0.3,     0.3,     0.4,     0.6,    93.2,   16.1,  0.00497,      0,      1,      91,      91,       0,     296
total,         58428,    3585,    3585,    3585,     0.3,     0.3,     0.3,     0.4,     1.0,     3.3,   17.1,  0.00501,      0,      0,       0,       0,       0,       0
total,         61902,    3450,    3450,    3450,     0.3,     0.3,     0.4,     0.5,     1.8,     3.3,   18.1,  0.00476,      0,      0,       0,       0,       0,       0
total,         65557,    3632,    3632,    3632,     0.3,     0.3,     0.3,     0.4,     0.6,     0.8,   19.1,  0.00496,      0,      0,       0,       0,       0,       0
total,         69212,    3641,    3641,    3641,     0.3,     0.3,     0.3,     0.4,     0.7,     1.8,   20.1,  0.00514,      0,      0,       0,       0,       0,       0
total,         72538,    3314,    3314,    3314,     0.3,     0.3,     0.3,     0.4,     1.6,    89.1,   21.1,  0.00518,      0,      1,      87,      87,       0,     298
total,         76134,    3575,    3575,    3575,     0.3,     0.3,     0.3,     0.4,     0.7,     4.4,   22.1,  0.00504,      0,      0,       0,       0,       0,       0
total,         79501,    3341,    3341,    3341,     0.3,     0.3,     0.3,     0.4,     4.8,    77.0,   23.1,  0.00497,      0,      1,      75,      75,       0,     304
total,         82863,    3346,    3346,    3346,     0.3,     0.2,     0.4,     0.7,     6.6,    49.2,   24.1,  0.00476,      0,      1,      48,      48,       0,     322
total,         86046,    3159,    3159,    3159,     0.3,     0.3,     0.5,     1.1,     4.3,    25.1,   25.1,  0.00555,      0,      1,      23,      23,       0,     320
total,         89568,    3487,    3487,    3487,     0.3,     0.3,     0.3,     0.6,     2.6,    26.6,   26.1,  0.00539,      0,      1,      24,      24,       0,     313
total,         93218,    3633,    3633,    3633,     0.3,     0.3,     0.3,     0.5,     1.3,     2.7,   27.1,  0.00536,      0,      0,       0,       0,       0,       0
total,         96240,    3006,    3006,    3006,     0.3,     0.3,     0.5,     0.8,     3.1,    45.5,   28.1,  0.00636,      0,      1,      37,      37,       0,     312
total,         99545,    3282,    3282,    3282,     0.3,     0.3,     0.4,     0.6,     1.1,    30.0,   29.1,  0.00626,      0,      1,      28,      28,       0,     314
total,        102904,    3346,    3346,    3346,     0.3,     0.3,     0.4,     0.6,     1.5,     2.0,   30.1,  0.00621,      0,      0,       0,       0,       0,       0


Results:
op rate                   : 3414 [insert:3414]
partition rate            : 3414 [insert:3414]
row rate                  : 3414 [insert:3414]
latency mean              : 0.3 [insert:0.3]
latency median            : 0.3 [insert:0.3]
latency 95th percentile   : 0.3 [insert:0.3]
latency 99th percentile   : 0.5 [insert:0.5]
latency 99.9th percentile : 1.3 [insert:1.3]
latency max               : 98.2 [insert:98.2]
Total partitions          : 102904 [insert:102904]
Total errors              : 0 [insert:0]
total gc count            : 10
total gc mb               : 3068
total gc time (s)         : 1
avg gc time(ms)           : 60
stdev gc time(ms)         : 29
Total operation time      : 00:00:30
END
```
