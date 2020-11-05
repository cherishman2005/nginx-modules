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

![c-module动态负载均衡](/img/c-module-upsync-upstream.png)

## nginx-lua-upstream动态负载均衡

![lua动态负载均衡](/img/lua-upsync-upstream.png)

## nginx-njs动态负载均衡

![njs动态负载均衡](/img/njs-upsync-upstream.png)

* nginx-njs没有挖掘http-upstream模块，主要在js-http-module采用js_set回调函数实现

# 小结


# Author

zhangbiwu

欢迎交流nginx二次开发技术。
