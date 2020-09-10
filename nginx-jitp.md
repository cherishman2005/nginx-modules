# nginx-jitp即时转封装系统设计

IPTV转OTT系统设计

![nginx-jitp](/img/nginx-jitp.png)


## nginx-jitp

nginx-jitp即时转封装功能： 基于nginx框架的C module开发。

需要的nginx技术栈：

* subrequest子请求（C module） + upstream

  subreqeust包含了 content-handle

* header-filter/body-filter（C module）

# FAQ

## 设计中遇到的坑

* 如果不用subrequest子请求，就会出现body-filter是range（偏移）的数据；所以必须用subrequest

* http range 是获取整体中的部分数据；

# Author

zhangbiwu 

【注】
* 设计过程中借鉴了openresty作者章亦春的nginx echo c module。
* 如有兴趣欢迎联系交流

# 参考链接

* https://github.com/cherishman2005/ngx_http_jitp_module

* https://github.com/openresty/echo-nginx-module
