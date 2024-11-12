# k8s-cache缓存机制

K8S里的几种缓存

* lruCache

可以设置一个固定的大小, 每次新增的时候会淘汰最晚未被使用的元素

* singleCache

* LRUExpireCache

底层基于lruCache实现的, 添加元素的时候会计算超时时间, 获取元素的时候判断元素有没有过期, 和guava cache类似

* client-go cache

ThreadSafeStore : 内部是一个有读写锁的map, 可以根据对象计算索引key, 然后根据索引key找到缓存的对象.

* scheduler cache


# 参考链接

- [https://blog.csdn.net/levena/article/details/127298786](https://blog.csdn.net/levena/article/details/127298786)
