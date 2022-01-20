# ab压测

## ab安装

```
apt install apache2-utils
```

## 压测示例

```
# ab -n 20 -c 1 'http://localhost:28080/test-sinfo?action=getSidAuthState&tid=36569435&sid=2790311106' 
```

其中－n表示请求数，－c表示并发数

运行结果：
```
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient).....done


Server Software:        
Server Hostname:        localhost
Server Port:            28080

Document Path:          /test-sinfo?action=getSidAuthState&tid=36569435&sid=2790311106
Document Length:        90 bytes

Concurrency Level:      1
Time taken for tests:   0.017 seconds
Complete requests:      20
Failed requests:        0
Total transferred:      4140 bytes
HTML transferred:       1800 bytes
Requests per second:    1148.83 [#/sec] (mean)
Time per request:       0.870 [ms] (mean)
Time per request:       0.870 [ms] (mean, across all concurrent requests)
Transfer rate:          232.23 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     0    1   0.9      1       4
Waiting:        0    1   0.9      1       4
Total:          0    1   0.9      1       4

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      3
  95%      4
  98%      4
  99%      4
 100%      4 (longest request)
```


# 参考链接

- [https://www.jianshu.com/p/43d04d8baaf7](https://www.jianshu.com/p/43d04d8baaf7)
