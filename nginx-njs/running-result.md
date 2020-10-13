# demo运行结果

```
curl http://127.0.0.1/hello

{"name":"nginx-upstream"}
```


```
curl http://127.0.0.1/foo -v
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 80 (#0)
> GET /foo HTTP/1.1
> Host: 127.0.0.1
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Server: openresty
< Date: Tue, 13 Oct 2020 04:41:54 GMT
< Content-Type: text/plain; charset=utf-8
< Connection: keep-alive
< foo: 1234
< Content-Length: 15
< X-Foo: foo
< 
* Connection #0 to host 127.0.0.1 left intact
nginxjavascript
```



```
curl 'http://127.0.0.1/summary?a=1&b=2222'

JS summary

Method: GET
HTTP version: 1.1
Host: 127.0.0.1
Remote Address: 127.0.0.1
URI: /summary
Headers:
  header 'Host' is '127.0.0.1'
  header 'User-Agent' is 'curl/7.47.0'
  header 'Accept' is '*/*'
Args:
  arg 'a' is '1'
  arg 'b' is '2222'
```

用lua也完全可以实现如上功能；只是javascript更加通用的语言。

