# nginx的ip-hash机制

【ip-hash机制】
nginx中的ip-hash技术能够将某个ip 的请求定向到同一台后端web机器中，这样一来这个ip下的客户端和某个后端web机器就能建立起稳固的session。

【典型应用场景】
ip-hash机制能够让某一客户机在相当长的一段时间内只访问固定的后端的某台真实的web服务器，这样会话就会得以保持，在网站页面进行login的时候就不会在后面的web服务器之间跳来跳去了，也不会出现登录一次的网站又提醒重新登录的情况。

ip-hash是在upstream配置中定义的:
```
upstream backend {
    server 192.168.74.235:80;
    server 192.168.74.236:80;
    ip_hash;
}

server {
    listen 80;

    location / {
        proxy_pass http://backend;
    }
}
```

# ip-hash机制缺陷

（1）nginx不是最前端的服务器

ip-hash要求nginx一定是最前端的服务器，否则nginx得不到正确ip，就不能根据ip作hash. 例如：使用的是squid为最前端，那么nginx取ip时只能得到squid的服务器ip地址，用这个地址来作分流肯定是错乱的。


（2）nginx的后端还有其它负载均衡

假如nginx后端还有其它负载均衡,将请求又通过另外的方式分流了,那么某个客户端的请求肯定不能定位到同一台session应用服务器上，这么算起来，nginx后端只能直接指向应用服务器。
* 需要通过业务路由proxy-server，指向应用服务器。
* 其他办法：用location作一次分流，将需要session的部分请求通过ip-hash分流，剩下的走其它后端去。

# 小结

ip-hash机制缺陷的解决方法：

（1）nginx的前面还有proxy-server，采用其他的hash-key，如对args或请求头字段做hash；

（2）nginx的后端还有其它负载均衡； 
     * 这个已经是整体的设计方面的优化问题，有很多解决方法。

ip-hash只是nginx普通hash中的一种形式，以remoteAddress ip为key值。系统设计时有很多hash算法均衡算法，可以c-module、nginx-lua，nginx-njs实现。
