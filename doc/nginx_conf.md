

# 下载速率限制

```
location / {
            proxy_pass http://localhost:8999/;    #转发到http://localhost:8999
            limit_conn one 10;    #限制并发连接数
            limit_rate 200k;    #限制最高下载速度
            limit_rate_after 1000k;    #下载到指定的文件大小之后开始限速
        }
```


# proxy代理超时配置

proxy_connect_timeout :后端服务器连接的超时时间_发起握手等候响应超时时间

proxy_read_timeout:连接成功后_等候后端服务器响应时间_其实已经进入后端的排队之中等候处理（也可以说是后端服务器处理请求的时间）

proxy_send_timeout :后端服务器数据回传时间_就是在规定时间之内后端服务器必须传完所有的数据


设置代理超时10s

```
proxy_read_timeout 10s;
```

默认超时是xxs。-- 为了解决后端响应时延较大的应急处理。


# nginx ip白名单

```
allow 127.0.0.1;
deny all;
```


# 参考链接

- [nginx timeout 配置](https://www.cnblogs.com/derekchen/archive/2012/04/20/2459106.html)