# nginx-modules

## nginx-consul完美组合：

	（1）服务发现API网关；
	（2）后端微服务架构：负载均衡；
	
## nginx
	（1）http/https API网关：限流，黑白名单，接入控制；—— 成熟的lua插件，C module；
	（2）http2;
	（3）http CDN代理缓存；
	（4）upsteram静态负载均衡；（4层，7层负载均衡）；
	（5）c module插件开发，做点播/直播的转码转封装框架；

## consul
	（1）microservice服务发现；
	（2）key/value配置中心；
	（3）health健康检测；
	（4）与nginx结合，实现（4层/7层）动态负载均衡；


# C++11

- [c++11中std::thread的join用法](./doc/c++11中std::thread的join用法.md)

- [c++11 cast标准写法](./c++11_cast.md)
- [使用stl中的advance和distance方法来进行iterator的加减](/cplus/使用stl中的advance和distance方法来进行iterator的加减.md)

# thread

## pthread_kill

函数pthread_kill()向同一进程下的另一线程发送信号。因为仅在同一进程中可保证线程ID的唯一性，所以无法调用pthread_kill向其他进程中的线程发送信号。

# go

- [https://go.dev/](https://go.dev/)

- [http://c.biancheng.net/view/94.html](http://c.biancheng.net/view/94.html)

- [go-redis-cluster读写分离](./go-redis-cluster读写分离.md)

- [go-mysql](./go-mysql.md)

- [go-sync.Once](./go-sync.Once.md)

- [Golang结构体类型的深浅拷贝](./Golang结构体类型的深浅拷贝.md)

# ab压测

- [ab压测](doc/ab-perf.md)

# 参考链接

- [https://en.wikipedia.org/wiki/Reserved_IP_addresses](https://en.wikipedia.org/wiki/Reserved_IP_addresses)
