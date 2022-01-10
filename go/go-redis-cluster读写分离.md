# go-redis-cluster读写分离

![0F5ECF682B3D9FE9CE52B9F2B7A5EBF9](https://user-images.githubusercontent.com/17688273/148764159-9ed4efcf-7aaa-4171-a70b-d13557248b5b.jpg)

![BEC927C8778AD360FE7DCD54AB6B6A0E](https://user-images.githubusercontent.com/17688273/148764139-34d8e13b-f3ea-4284-ba45-f6d46b4af857.jpg)

go-redis-cluster 设置了 ReadyOnly后，读写是分离的。读只请求slave从节点。


* 若是ReadOnly = true，只选择Slave Node

* 若是ReadOnly = true 且 RouteByLatency = true 将从slot对应的Master Node 和 Slave Node选择，选择策略为: 选择PING 延迟最低的节点

* 若是ReadOnly = true 且 RouteRandomly = true 将从slot对应的Master Node 和 Slave Node选择，选择策略为:随机选择

![76EA1D40184563030A123E0418D0381A](https://user-images.githubusercontent.com/17688273/148764206-5611c890-40db-43f3-8900-edd9c5898a62.jpg)


# 参考链接

- [https://github.com/go-redis/redis](https://github.com/go-redis/redis)

- [https://www.shangmayuan.com/a/dcb403dcb4e94c22ab6aab5e.html](https://www.shangmayuan.com/a/dcb403dcb4e94c22ab6aab5e.html)
