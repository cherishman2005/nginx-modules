# nginx动态负载均衡

Nginx负载均衡经历的阶段：

静态负载均衡  ==》 动态负载均衡


根据服务发现的成熟度：

（1）在配置中心配置key-value。 Nginx定时动态拉去key-value值。

（2）nginx去服务中心拉取服务IP+Port列表。

以上方案均不用nginx -s reload，减少重启nginx的性能损耗 和 抖动。

代码实现方式:
（1）nginx c/c++ module

（2）openresty lua实现

（3）nginx njs实现


## nginx-c-module-upstream动态负载均衡

![c-module动态负载均衡](/img/upsync-upstream/c-module-upsync-upstream.png)

## nginx-lua-upstream动态负载均衡

![lua动态负载均衡](/img/upsync-upstream/lua-upsync-upstream.png)

## nginx-njs动态负载均衡

![njs动态负载均衡](/img/upsync-upstream/njs-upsync-upstream.png)

* nginx-njs没有挖掘http-upstream模块，主要在js-http-module采用js_set回调函数实现

* njs的js模块暂时不支持共享内存功能；

# 小结

| nginx动态负载均衡       |    | 
| ---------- | ------ |
| c module upstream      | 开发门槛要求高 | 
| lua-upstream     | 开发效率高，扩展个性化upstream均衡算法方便 |
| njs-upstream     | 开发效率高，可扩展个性化均衡算法，暂没有lua-upstream成熟 | 

* lua实现upstream的关键在upstream模块；njs实现upstream在http-js-module，充分利用js的回调函数特性；


# Author

zhangbiwu

欢迎交流nginx二次开发技术。
