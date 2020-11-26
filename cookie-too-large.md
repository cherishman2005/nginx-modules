
# cookie too large

【问题描述】
400 bad request

![](/img/cookie-too-large.png)


Cookie 是小甜饼的意思。顾名思义，cookie 确实非常小，它的大小限制为4KB左右。

## 浏览器本地storage



## 分析

到底是chrome浏览器cookie存储大小限制？ 还是后端服务对cookie大小或请求头大小的限制？

* 第一步

分析cookie大小增大的原因，分析是哪个组件（或sdk）的cookie引起？

* 第二步

分析nginx的cookie或请求头限制，并进行验证。是否能增大cookie或header的配置？
