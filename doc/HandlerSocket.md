# HandlerSocket

HandlerSocket是针对Mysql的一个NoSQL插件，它作为一个守护进程工作在mysqld进程里面,接收tcp连接，并处理来自客户端的请求。HandlerSocket不支持SQL查询，作为替代，它支持表的简单的CRUD操作。
由于下面的原因，在某些情况下HandlerSocket比mysqld/libmysql对儿更快速：
* HandlerSocket 处理数据不需要解析SQL，由于这个原因使得其占用少量CPU资源。
* HandlerSocket 从客户端批量读取多个请求并处理他们，这使得其占用更少的CPU和磁盘使用率。
* HandlerSocket 客户端/服务器协议比mysql/libmysql对儿更简洁，这使得其占用更少的网络使用率。
