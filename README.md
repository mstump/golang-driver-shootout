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
These are the results from single threaded inserts from my macbook air to my macbook air. Don't use these results as a reference for absolute throughput. The results are only useful for a relative throughput comparison between drivers on the same (constrained) hardware and can be used as a proxy to judge efficency.


### C++
Modified the ```examples/perf.c``` code to log at ERROR and not insert into the set.

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

### GoLang Wrapping C++

```
➜  golang-driver git:(master) ✗ go run examples/perf.go
rate stats (requests/second): mean 5655 1m 13872.4 5m 13872.4 10m 13872.4 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5109 1m 14014.609004140286 5m 13901.797672232871 10m 13882.25371433849 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5046 1m 13940.473465453288 5m 13888.336862699804 10m 13877.850103947183 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 4845 1m 13845.272471537637 5m 13869.518503588879 10m 13871.600534913223 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 4921 1m 13797.021476820757 5m 13859.143229315285 10m 13868.111344223205 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5060 1m 13307.291799233586 5m 13756.87874673728 10m 13833.783945597625 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5033 1m 13098.136639159511 5m 13706.210839712983 10m 13816.37466091204 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5122 1m 13112.645648398842 5m 13699.159582748589 10m 13813.400847703468 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5129 1m 13124.331505135584 5m 13691.881079050036 10m 13810.328274710577 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5183 1m 13109.609162857285 5m 13679.456883623541 10m 13805.507631813443 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5186 1m 13123.23286198283 5m 13672.854442076297 10m 13802.596239658611 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5212 1m 12915.025890178515 5m 13620.729118999778 10m 13784.405725430752 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5219 1m 12927.752749305148 5m 13611.69578721856 10m 13780.471077764072 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5251 1m 12904.633371972914 5m 13595.611928563529 10m 13774.144938434447 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5272 1m 12783.801823085118 5m 13559.212567007682 10m 13760.955249613948 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5296 1m 12751.867416522231 5m 13539.794623242951 10m 13753.32892430143 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5308 1m 12724.453251529665 5m 13521.104232094414 10m 13745.881137729386 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5320 1m 12742.119178139359 5m 13511.588678109152 10m 13741.446350557151 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5330 1m 12548.265328008794 5m 13458.796688749793 10m 13722.477721367184 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5336 1m 12566.617907070717 5m 13447.540803532467 10m 13717.244063204422 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5344 1m 12566.617907070717 5m 13447.540803532467 10m 13717.244063204422 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5353 1m 12262.573900602973 5m 13355.173979082663 10m 13683.213338348509 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5356 1m 12007.869474638712 5m 13284.461982910241 10m 13657.694216332186 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5368 1m 11593.949845406387 5m 13177.795638515274 10m 13619.873292749613 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5374 1m 11284.077988883684 5m 13087.559765598015 10m 13587.178233712737 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5392 1m 11100.701613747728 5m 13019.842976921353 10m 13561.712491640585 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5395 1m 10991.969884174001 5m 12965.645163656407 10m 13540.544055676906 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5408 1m 11037.41904682719 5m 12942.418504076963 10m 13529.573756636257 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5406 1m 11093.802202730374 5m 12922.587248526648 10m 13519.673650413231 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
rate stats (requests/second): mean 5417 1m 10962.850768889683 5m 12865.289601107983 10m 13497.160280731447 connection_timeouts 0 pending_request_timeouts 0 request_timeouts 0
final stats (microseconds): min 147 max 242687 median 2591 75th 5839 95th 14143 98th 0 99th 77695 99.9th 157567
DONE.
```

### GoLang
```
23 Jun 2015 08:18:56 | seconds : 00000005 | total : 0000039059 | succes : 1.00 | failure : 0.00 | P50 : 381.00 µs | P95 : 1860.00 µs | P99 : 404411.00 µs
23 Jun 2015 08:19:01 | seconds : 00000010 | total : 0000084151 | succes : 1.00 | failure : 0.00 | P50 : 354.00 µs | P95 : 1474.00 µs | P99 : 240257.00 µs
23 Jun 2015 08:19:06 | seconds : 00000015 | total : 0000126869 | succes : 1.00 | failure : 0.00 | P50 : 355.00 µs | P95 : 1197.00 µs | P99 : 4147.00 µs
23 Jun 2015 08:19:11 | seconds : 00000020 | total : 0000169758 | succes : 1.00 | failure : 0.00 | P50 : 358.00 µs | P95 : 1433.00 µs | P99 : 2894.00 µs
23 Jun 2015 08:19:16 | seconds : 00000025 | total : 0000212608 | succes : 1.00 | failure : 0.00 | P50 : 363.00 µs | P95 : 1191.00 µs | P99 : 3314.00 µs
23 Jun 2015 08:19:21 | seconds : 00000030 | total : 0000256136 | succes : 1.00 | failure : 0.00 | P50 : 362.00 µs | P95 : 1749.00 µs | P99 : 207879.00 µs
23 Jun 2015 08:19:26 | seconds : 00000035 | total : 0000296384 | succes : 1.00 | failure : 0.00 | P50 : 360.00 µs | P95 : 1491.00 µs | P99 : 4355.00 µs
23 Jun 2015 08:19:31 | seconds : 00000040 | total : 0000340350 | succes : 1.00 | failure : 0.00 | P50 : 360.00 µs | P95 : 1143.00 µs | P99 : 2697.00 µs
23 Jun 2015 08:19:36 | seconds : 00000045 | total : 0000380378 | succes : 1.00 | failure : 0.00 | P50 : 379.00 µs | P95 : 1996.00 µs | P99 : 183584.00 µs
23 Jun 2015 08:19:41 | seconds : 00000050 | total : 0000420323 | succes : 1.00 | failure : 0.00 | P50 : 379.00 µs | P95 : 1201.00 µs | P99 : 3648.00 µs
23 Jun 2015 08:19:46 | seconds : 00000055 | total : 0000457529 | succes : 1.00 | failure : 0.00 | P50 : 382.00 µs | P95 : 2014.00 µs | P99 : 4199.00 µs
23 Jun 2015 08:19:51 | seconds : 00000060 | total : 0000490185 | succes : 1.00 | failure : 0.00 | P50 : 401.00 µs | P95 : 2246.00 µs | P99 : 149739.00 µs
23 Jun 2015 08:19:56 | seconds : 00000065 | total : 0000525514 | succes : 1.00 | failure : 0.00 | P50 : 394.00 µs | P95 : 1659.00 µs | P99 : 4007.00 µs
23 Jun 2015 08:20:01 | seconds : 00000070 | total : 0000564621 | succes : 1.00 | failure : 0.00 | P50 : 382.00 µs | P95 : 1572.00 µs | P99 : 100618.00 µs
23 Jun 2015 08:20:06 | seconds : 00000075 | total : 0000604250 | succes : 1.00 | failure : 0.00 | P50 : 384.00 µs | P95 : 1912.00 µs | P99 : 97317.00 µs
23 Jun 2015 08:20:11 | seconds : 00000080 | total : 0000644039 | succes : 1.00 | failure : 0.00 | P50 : 384.00 µs | P95 : 1244.00 µs | P99 : 91405.00 µs
23 Jun 2015 08:20:16 | seconds : 00000085 | total : 0000683354 | succes : 1.00 | failure : 0.00 | P50 : 384.00 µs | P95 : 1381.00 µs | P99 : 4135.00 µs
23 Jun 2015 08:20:21 | seconds : 00000090 | total : 0000721868 | succes : 1.00 | failure : 0.00 | P50 : 374.00 µs | P95 : 1919.00 µs | P99 : 9356.00 µs
23 Jun 2015 08:20:26 | seconds : 00000095 | total : 0000760253 | succes : 1.00 | failure : 0.00 | P50 : 375.00 µs | P95 : 1781.00 µs | P99 : 79226.00 µs
23 Jun 2015 08:20:31 | seconds : 00000100 | total : 0000802587 | succes : 1.00 | failure : 0.00 | P50 : 342.00 µs | P95 : 1100.00 µs | P99 : 171078.00 µs
23 Jun 2015 08:20:36 | seconds : 00000105 | total : 0000848679 | succes : 1.00 | failure : 0.00 | P50 : 336.00 µs | P95 : 1312.00 µs | P99 : 73368.00 µs
23 Jun 2015 08:20:41 | seconds : 00000110 | total : 0000892637 | succes : 1.00 | failure : 0.00 | P50 : 333.00 µs | P95 : 951.00 µs | P99 : 4889.00 µs
23 Jun 2015 08:20:46 | seconds : 00000115 | total : 0000935251 | succes : 1.00 | failure : 0.00 | P50 : 347.00 µs | P95 : 1433.00 µs | P99 : 112801.00 µs
23 Jun 2015 08:20:51 | seconds : 00000120 | total : 0000977797 | succes : 1.00 | failure : 0.00 | P50 : 343.00 µs | P95 : 1460.00 µs | P99 : 141141.00 µs
23 Jun 2015 08:20:56 | seconds : 00000125 | total : 0001015397 | succes : 1.00 | failure : 0.00 | P50 : 381.00 µs | P95 : 2313.00 µs | P99 : 5095.00 µs
23 Jun 2015 08:21:01 | seconds : 00000130 | total : 0001054495 | succes : 1.00 | failure : 0.00 | P50 : 361.00 µs | P95 : 1550.00 µs | P99 : 3406.00 µs
23 Jun 2015 08:21:06 | seconds : 00000135 | total : 0001096463 | succes : 1.00 | failure : 0.00 | P50 : 360.00 µs | P95 : 1458.00 µs | P99 : 3586.00 µs
23 Jun 2015 08:21:11 | seconds : 00000140 | total : 0001136174 | succes : 1.00 | failure : 0.00 | P50 : 382.00 µs | P95 : 1719.00 µs | P99 : 55946.00 µs
23 Jun 2015 08:21:16 | seconds : 00000145 | total : 0001177025 | succes : 1.00 | failure : 0.00 | P50 : 380.00 µs | P95 : 1347.00 µs | P99 : 3107.00 µs
23 Jun 2015 08:21:21 | seconds : 00000150 | total : 0001222443 | succes : 1.00 | failure : 0.00 | P50 : 336.00 µs | P95 : 1091.00 µs | P99 : 3202.00 µs
23 Jun 2015 08:21:26 | seconds : 00000155 | total : 0001263349 | succes : 1.00 | failure : 0.00 | P50 : 373.00 µs | P95 : 1359.00 µs | P99 : 3650.00 µs
23 Jun 2015 08:21:31 | seconds : 00000160 | total : 0001305960 | succes : 1.00 | failure : 0.00 | P50 : 346.00 µs | P95 : 1210.00 µs | P99 : 134373.00 µs
23 Jun 2015 08:21:36 | seconds : 00000165 | total : 0001346037 | succes : 1.00 | failure : 0.00 | P50 : 369.00 µs | P95 : 1449.00 µs | P99 : 290587.00 µs
23 Jun 2015 08:21:41 | seconds : 00000170 | total : 0001375162 | succes : 1.00 | failure : 0.00 | P50 : 395.00 µs | P95 : 2019.00 µs | P99 : 142994.00 µs
23 Jun 2015 08:21:46 | seconds : 00000175 | total : 0001408714 | succes : 1.00 | failure : 0.00 | P50 : 357.00 µs | P95 : 1993.00 µs | P99 : 134888.00 µs
23 Jun 2015 08:21:51 | seconds : 00000180 | total : 0001444775 | succes : 1.00 | failure : 0.00 | P50 : 368.00 µs | P95 : 1566.00 µs | P99 : 125319.00 µs
23 Jun 2015 08:21:56 | seconds : 00000185 | total : 0001477511 | succes : 1.00 | failure : 0.00 | P50 : 370.00 µs | P95 : 1747.00 µs | P99 : 6255.00 µs
23 Jun 2015 08:22:01 | seconds : 00000190 | total : 0001508650 | succes : 1.00 | failure : 0.00 | P50 : 397.00 µs | P95 : 2946.00 µs | P99 : 5383.00 µs
23 Jun 2015 08:22:06 | seconds : 00000195 | total : 0001542626 | succes : 1.00 | failure : 0.00 | P50 : 371.00 µs | P95 : 1462.00 µs | P99 : 4782.00 µs
23 Jun 2015 08:22:11 | seconds : 00000200 | total : 0001571563 | succes : 1.00 | failure : 0.00 | P50 : 388.00 µs | P95 : 1726.00 µs | P99 : 7619.00 µs
23 Jun 2015 08:22:16 | seconds : 00000205 | total : 0001603905 | succes : 1.00 | failure : 0.00 | P50 : 380.00 µs | P95 : 1444.00 µs | P99 : 5804.00 µs
23 Jun 2015 08:22:21 | seconds : 00000210 | total : 0001642614 | succes : 1.00 | failure : 0.00 | P50 : 345.00 µs | P95 : 2670.00 µs | P99 : 6730.00 µs
23 Jun 2015 08:22:26 | seconds : 00000215 | total : 0001678948 | succes : 1.00 | failure : 0.00 | P50 : 362.00 µs | P95 : 1358.00 µs | P99 : 4136.00 µs
23 Jun 2015 08:22:31 | seconds : 00000220 | total : 0001713969 | succes : 1.00 | failure : 0.00 | P50 : 305.00 µs | P95 : 1149.00 µs | P99 : 124368.00 µs
23 Jun 2015 08:22:36 | seconds : 00000225 | total : 0001745546 | succes : 1.00 | failure : 0.00 | P50 : 316.00 µs | P95 : 906.00 µs | P99 : 1842.00 µs
23 Jun 2015 08:22:41 | seconds : 00000230 | total : 0001777178 | succes : 1.00 | failure : 0.00 | P50 : 269.00 µs | P95 : 529.00 µs | P99 : 1060.00 µs
23 Jun 2015 08:22:46 | seconds : 00000235 | total : 0001801597 | succes : 1.00 | failure : 0.00 | P50 : 264.00 µs | P95 : 546.00 µs | P99 : 126325.00 µs
23 Jun 2015 08:22:51 | seconds : 00000240 | total : 0001819936 | succes : 1.00 | failure : 0.00 | P50 : 236.00 µs | P95 : 351.00 µs | P99 : 32772.00 µs
23 Jun 2015 08:22:56 | seconds : 00000245 | total : 0001838717 | succes : 1.00 | failure : 0.00 | P50 : 230.00 µs | P95 : 333.00 µs | P99 : 655.00 µs
23 Jun 2015 08:23:01 | seconds : 00000250 | total : 0001857910 | succes : 1.00 | failure : 0.00 | P50 : 233.00 µs | P95 : 316.00 µs | P99 : 42094.00 µs
23 Jun 2015 08:23:06 | seconds : 00000255 | total : 0001877508 | succes : 1.00 | failure : 0.00 | P50 : 233.00 µs | P95 : 359.00 µs | P99 : 41191.00 µs
23 Jun 2015 08:23:11 | seconds : 00000260 | total : 0001897128 | succes : 1.00 | failure : 0.00 | P50 : 228.00 µs | P95 : 311.00 µs | P99 : 42751.00 µs
23 Jun 2015 08:23:16 | seconds : 00000265 | total : 0001915519 | succes : 1.00 | failure : 0.00 | P50 : 238.00 µs | P95 : 330.00 µs | P99 : 591.00 µs
23 Jun 2015 08:23:21 | seconds : 00000270 | total : 0001934868 | succes : 1.00 | failure : 0.00 | P50 : 234.00 µs | P95 : 310.00 µs | P99 : 54155.00 µs
23 Jun 2015 08:23:26 | seconds : 00000275 | total : 0001954897 | succes : 1.00 | failure : 0.00 | P50 : 236.00 µs | P95 : 304.00 µs | P99 : 44528.00 µs
23 Jun 2015 08:23:31 | seconds : 00000280 | total : 0001975225 | succes : 1.00 | failure : 0.00 | P50 : 227.00 µs | P95 : 289.00 µs | P99 : 389.00 µs
23 Jun 2015 08:23:36 | seconds : 00000285 | total : 0001993170 | succes : 1.00 | failure : 0.00 | P50 : 237.00 µs | P95 : 356.00 µs | P99 : 61007.00 µs
23 Jun 2015 08:23:38 | seconds : 00000287 | total : 0002000000 | succes : 1.00 | failure : 0.00 | P50 : 229.00 µs | P95 : 299.00 µs | P99 : 440.00 µs
```

### Java

```
/Users/mstump/dse/resources/cassandra/tools/bin/cassandra-stress user profile=/Users/mstump/src/golang-driver-shootout/stress.yaml ops\(insert=1\) -rate threads=15
```

```
Connected to cluster: Test Cluster
Datatacenter: Cassandra; Host: localhost/127.0.0.1; Rack: rack1
Created schema. Sleeping 1s for propagation.
Sleeping 2s...
Warming up insert with 50000 iterations...
Generating batches with [1..1] partitions and [1..1] rows (of [1..1] total rows in the partitions)
Running [insert] with 15 threads until stderr of mean < 0.02
type,      total ops,    op/s,    pk/s,   row/s,    mean,     med,     .95,     .99,    .999,     max,   time,   stderr, errors,  gc: #,  max ms,  sum ms,  sdv ms,      mb
total,         11176,   11145,   11145,   11145,     1.3,     1.1,     2.3,     4.3,    59.2,    59.6,    1.0,  0.00000,      0,      1,      57,      57,       0,     292
total,         24805,   13518,   13518,   13518,     1.1,     1.0,     1.7,     2.3,    55.9,    56.7,    2.0,  0.06664,      0,      1,      53,      53,       0,     303
total,         37683,   12822,   12822,   12822,     1.2,     1.0,     1.8,     3.1,    75.5,    76.2,    3.0,  0.04642,      0,      1,      73,      73,       0,     282
total,         49309,   11527,   11527,   11527,     1.3,     1.0,     2.1,     3.8,    86.6,    87.1,    4.0,  0.03729,      0,      1,      84,      84,       0,     286
total,         60624,   11245,   11245,   11245,     1.3,     1.1,     1.8,     2.7,    75.6,    76.1,    5.0,  0.03328,      0,      2,      74,     137,       6,     595
total,         72402,   11715,   11715,   11715,     1.3,     1.1,     2.0,     3.4,    56.2,    56.6,    6.0,  0.02868,      0,      1,      53,      53,       0,     293
total,         84985,   12480,   12480,   12480,     1.2,     1.1,     1.8,     2.7,    55.8,    57.6,    7.0,  0.02474,      0,      1,      54,      54,       0,     294
total,         97808,   12709,   12709,   12709,     1.2,     1.0,     1.8,     3.0,    79.0,    79.7,    8.1,  0.02286,      0,      1,      77,      77,       0,     289
total,        109174,   11304,   11304,   11304,     1.3,     1.1,     1.9,     2.7,    81.8,    82.3,    9.1,  0.02127,      0,      1,      79,      79,       0,     286
total,        120284,   11025,   11025,   11025,     1.4,     1.2,     2.1,     3.7,    64.5,    64.9,   10.1,  0.02115,      0,      1,      61,      61,       0,     292
total,        131728,   11361,   11361,   11361,     1.3,     1.1,     2.1,     3.0,    54.9,    55.3,   11.1,  0.02011,      0,      1,      52,      52,       0,     306
total,        142656,   10861,   10861,   10861,     1.4,     1.1,     2.0,     3.0,    64.8,    65.6,   12.1,  0.01994,      0,      2,      62,     114,       5,     588
total,        155760,   13017,   13017,   13017,     1.1,     1.0,     1.8,     3.1,    63.2,    63.7,   13.1,  0.01962,      0,      1,      60,      60,       0,     289
total,        167853,   12000,   12000,   12000,     1.2,     1.1,     1.9,     2.6,    58.7,    59.3,   14.1,  0.01822,      0,      1,      57,      57,       0,     287
total,        178746,   10834,   10834,   10834,     1.4,     1.2,     2.2,     3.7,    53.1,    53.4,   15.1,  0.01838,      0,      1,      51,      51,       0,     309
total,        189346,   10522,   10522,   10522,     1.4,     1.2,     2.3,     3.8,    57.6,    78.7,   16.1,  0.01881,      0,      2,      56,      84,      14,     289
total,        201504,   12080,   12080,   12080,     1.2,     1.1,     2.0,     3.0,    30.4,    31.1,   17.1,  0.01772,      0,      1,      27,      27,       0,     305
total,        214087,   12458,   12458,   12458,     1.2,     1.0,     2.0,     2.9,    72.0,    72.6,   18.1,  0.01707,      0,      1,      70,      70,       0,     290
total,        224728,   11129,   11129,   11129,     1.3,     1.1,     1.9,     3.0,    86.2,    86.6,   19.1,  0.01627,      0,      1,      84,      84,       0,     287
total,        233899,    8133,    8133,    8133,     1.8,     1.3,     2.8,     5.3,   123.3,   124.6,   20.2,  0.02075,      0,      2,      94,     158,      15,     598
total,        243239,    9264,    9264,    9264,     1.6,     1.1,     2.3,     3.8,   134.2,   135.3,   21.2,  0.02093,      0,      1,      65,      65,       0,     298
total,        253669,   10348,   10348,   10348,     1.4,     1.2,     2.7,     5.1,    49.8,    50.1,   22.2,  0.02076,      0,      1,      47,      47,       0,     295
total,        266189,   12410,   12410,   12410,     1.2,     1.1,     1.7,     2.5,    62.7,    63.0,   23.2,  0.02007,      0,      1,      60,      60,       0,     288
total,        277820,   11559,   11559,   11559,     1.3,     1.1,     2.0,     3.1,    72.1,    72.5,   24.2,  0.01923,      0,      2,     133,     201,      33,    1000
total,        289786,   11885,   11885,   11885,     1.3,     1.1,     1.9,     2.8,    55.7,    56.0,   25.2,  0.01847,      0,      1,      54,      54,       0,     289
total,        300768,   10865,   10865,   10865,     1.4,     1.1,     2.2,     3.5,    47.3,    80.3,   26.3,  0.01803,      0,      2,      45,      76,       7,     299
total,        311464,   10636,   10636,   10636,     1.4,     1.2,     2.4,     3.9,    25.9,    26.5,   27.3,  0.01791,      0,      1,      24,      24,       0,     313
total,        324614,   13034,   13034,   13034,     1.1,     1.0,     1.6,     2.5,    50.7,    50.8,   28.3,  0.01771,      0,      1,      48,      48,       0,     292
total,        336037,   11335,   11335,   11335,     1.3,     1.1,     2.3,     3.4,    59.0,    59.4,   29.3,  0.01712,      0,      1,      57,      57,       0,     285
total,        346561,   10438,   10438,   10438,     1.4,     1.2,     2.0,     3.4,    68.3,    68.9,   30.3,  0.01686,      0,      2,      66,     113,      10,     584


Results:
op rate                   : 11443 [insert:11443]
partition rate            : 11443 [insert:11443]
row rate                  : 11443 [insert:11443]
latency mean              : 1.3 [insert:1.3]
latency median            : 1.1 [insert:1.1]
latency 95th percentile   : 2.0 [insert:2.0]
latency 99th percentile   : 2.9 [insert:2.9]
latency 99.9th percentile : 7.2 [insert:7.2]
latency max               : 135.3 [insert:135.3]
Total partitions          : 346561 [insert:346561]
Total errors              : 0 [insert:0]
total gc count            : 37
total gc mb               : 10705
total gc time (s)         : 2
avg gc time(ms)           : 60
stdev gc time(ms)         : 20
Total operation time      : 00:00:30
END```
