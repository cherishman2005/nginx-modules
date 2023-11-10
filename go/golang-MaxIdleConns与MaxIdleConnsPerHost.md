# golang-MaxIdleConns与MaxIdleConnsPerHost

在业务上使用http访问，需要初始化httpClient，其中在高并发场景下，MaxIdleConns与MaxIdleConnsPerHost的配置会影响业务的请求效率。

Client.Transports属性中包含：
* MaxIdleConns  所有host的连接池最大连接数量，默认无穷大
* MaxIdleConnsPerHost  每个host的连接池最大空闲连接数,默认2
* MaxConnsPerHost 对每个host的最大连接数量，0表示不限制

相对应的代码解释(/src/net/http/transport.go:193:198:210)
```
 // MaxIdleConns controls the maximum number of idle (keep-alive)

 // connections across all hosts. Zero means no limit.

 MaxIdleConns int

 // MaxIdleConnsPerHost, if non-zero, controls the maximum idle

 // (keep-alive) connections to keep per-host. If zero,

 // DefaultMaxIdleConnsPerHost is used.

 MaxIdleConnsPerHost int

 // MaxConnsPerHost optionally limits the total number of

 // connections per host, including connections in the dialing,

 // active, and idle states. On limit violation, dials will block.

 //

 // Zero means no limit.

 //

 // For HTTP/2, this currently only controls the number of new

 // connections being created at a time, instead of the total

 // number. In practice, hosts using HTTP/2 only have about one

 // idle connection, though.

 MaxConnsPerHost int
```
测试：

如果MaxConnsPerHost=1，则只有一个http client被创建.

如果MaxIdleConnsPerHost=1，则会缓存一个http client.

golang 中默认http.client
```
 newClient := &http.Client{

        Timeout: time.Minute * 1, //设置超时时间

        Transport: &http.Transport{

            Dial: (&net.Dialer{

                Timeout:   30 * time.Second, //限制建立TCP连接的时间

                KeepAlive: 30 * time.Second,

            }).Dial,

            TLSHandshakeTimeout:   10 * time.Second, //限制 TLS握手的时间

            ResponseHeaderTimeout: 10 * time.Second, //限制读取response header的时间,默认 timeout + 5*time.Second

            ExpectContinueTimeout: 1 * time.Second,  //限制client在发送包含 Expect: 100-continue的header到收到继续发送body的response之间的时间等待。

            MaxIdleConns:          2,                //所有host的连接池最大连接数量，默认无穷大

            MaxIdleConnsPerHost:   1,                //每个host的连接池最大空闲连接数,默认2

            MaxConnsPerHost:       1,                //每个host的最大连接数量

            IdleConnTimeout:       3 * time.Minute,  //how long an idle connection is kept in the connection pool.

        },

    }
    ```
    
