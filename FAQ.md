# FAQ

## http Content-Length

nginx http请求头的Content-Length = -1时判定客户端不合法，nginx返回400错误码。

# 多机器集群

nodejs-cluster 适用于在单台机器上，如果应用的流量巨大，多机器是必然的。这时，反向代理就派上用场了，可以用 node 来写反向代理的服务（比如用 http-proxy），好处是可以保持工程师技术栈的统一，不过生产环境，更多的还是nginx。

