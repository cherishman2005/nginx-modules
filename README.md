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

- [c++11编程注意事项](/cplusplus/README.md)
- [c++11中thread的join用法](./doc/c++11中thread的join用法.md)

- [c++11 cast标准写法](./c++11_cast.md)
- [使用stl中的advance和distance方法来进行iterator的加减](/cplusplus/使用stl中的advance和distance方法来进行iterator的加减.md)

## json

* c++ json解析

- [https://github.com/nlohmann/json](https://github.com/nlohmann/json)

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

# git

- [git合并分支到master](./doc/git合并分支到master.md)
- [git_submodule](./doc/git_submodule.md)


# linux工具

- [/proc/stat解析](/doc/proc_stat解析.md)

# 小结

* 后端开发技术
  在开发需求的角度，如果有一定工作经验，比较简单。
  
  更重要的是多做积累和开发调试工具，方便监控和定位问题。

* 多借鉴开源技术
 
* 读代码可以读到人心

  通过一部分项目代码就可以看出：（如一个段代码的嵌套层数，1个function代码行数等）
  * 是否经过比较正规的代码规范编写
  * 或是否进过大公司
  * 或是否读过优秀的代码（或开源代码）；

* 不背锅，不甩锅

## 技术总结

1. http-api网关： openresty(nginx+lua)  不管是大体量 还是小体量业务。都或多或少可能面临被攻击的风险。 采用nginx网关开发一些lua插件，开发效率高，性能好。
2. golang/nodejs开发业务效率高；
3. rtc音视频： C++是必备条件，音视频传输，音视频处理； 音视频sdk。
4. 积极拥抱开源，在开源的基础上做二次开发。

## 对AI技术的理解

对AI技术要有一个比较正确的理解，在很大程度上，AI技术只是一个辅助手段。

1. 在直播公司，AI瘦脸技术，抠图等要多实践。-- 价值比较高。
   * 这些技术是在直播系统比较成熟、稳定的基础上去实践。

2. 挂机，低质识别，优质推荐。
   * 这些用AI技术存在一定的忽悠成分，非常烧钱。不要被一些只会demo的AI算法人员（没有工程实践经验）主导。
   * 这些功能 多采用统计学的方法 去实践，效果更好，成本更低。
   * 另外，采用音视频基础技术去做一些低质识别（如检测主播端的设备、画质、码率、分辨率等）。 -- 更加利用直播技术良性发展。

# FAQ

## 浏览器访问跨域问题

- [浏览器访问跨域问题](doc/浏览器访问跨域问题.md)

# 参考链接

- [https://en.wikipedia.org/wiki/Reserved_IP_addresses](https://en.wikipedia.org/wiki/Reserved_IP_addresses)
- [https://github.com/apache/arrow](https://github.com/apache/arrow)
