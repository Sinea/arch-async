##### Benchmark with operations sent to a background worker
```
Concurrency Level:      100
Time taken for tests:   7.242 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      11600000 bytes
Total body sent:        13200000
HTML transferred:       0 bytes
Requests per second:    13808.08 [#/sec] (mean)
Time per request:       7.242 [ms] (mean)
Time per request:       0.072 [ms] (mean, across all concurrent requests)
Transfer rate:          1564.20 [Kbytes/sec] received
                        1779.95 kb/s sent
                        3344.14 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   1.5      2      13
Processing:     0    5   9.9      4     244
Waiting:        0    5   9.9      4     244
Total:          0    7   9.9      6     247

Percentage of the requests served within a certain time (ms)
  50%      6
  66%      7
  75%      8
  80%      8
  90%     10
  95%     11
  98%     13
  99%     15
 100%    247 (longest request)
 ```

##### Benchmark with operations done inside the request
```
Concurrency Level:      100
Time taken for tests:   51.053 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      11600000 bytes
Total body sent:        13200000
HTML transferred:       0 bytes
Requests per second:    1958.75 [#/sec] (mean)
Time per request:       51.053 [ms] (mean)
Time per request:       0.511 [ms] (mean, across all concurrent requests)
Transfer rate:          221.89 [Kbytes/sec] received
                        252.49 kb/s sent
                        474.38 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       3
Processing:    50   51   0.4     51      61
Waiting:       49   51   0.4     51      61
Total:         50   51   0.6     51      64

Percentage of the requests served within a certain time (ms)
  50%     51
  66%     51
  75%     51
  80%     51
  90%     52
  95%     52
  98%     53
  99%     53
 100%     64 (longest request)
```
