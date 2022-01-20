# ab压测

## ab安装

```
apt install apache2-utils
```

## 压测示例

```
 ab -n 20 -c 10 'http://localhost:28080/test-sinfo?action=getSidAuthState&tid=36569435&sid=2790311106'
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

Concurrency Level:      10
Time taken for tests:   0.006 seconds
Complete requests:      20
Failed requests:        0
Total transferred:      4140 bytes
HTML transferred:       1800 bytes
Requests per second:    3346.16 [#/sec] (mean)
Time per request:       2.988 [ms] (mean)
Time per request:       0.299 [ms] (mean, across all concurrent requests)
Transfer rate:          676.42 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       0
Processing:     1    3   1.7      4       5
Waiting:        1    3   1.7      4       5
Total:          1    3   1.7      4       5

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      4
  75%      5
  80%      5
  90%      5
  95%      5
  98%      5
  99%      5
 100%      5 (longest request)
```


# 参考链接

- [https://www.jianshu.com/p/43d04d8baaf7](https://www.jianshu.com/p/43d04d8baaf7)
