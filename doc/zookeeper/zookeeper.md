# zookeeper

![image](https://github.com/user-attachments/assets/3190c15a-9918-46a0-8748-7fa265cfa8dc)


## ZooKeeper集群脑裂问题处理

### ZooKeeper是如何解决“脑裂”问题的？

要解决Split-Brain脑裂的问题，一般有下面几种方法：
* Quorums（法定人数）方式：比如3个节点的集群，Quorums = 2，也就是说集群可以容忍1个节点失效，这时候还能选举出1个lead，集群还可用。比如4个节点的集群，它的Quorums = 3，Quorums要超过3，相当于集群的容忍度还是1，如果2个节点失效，那么整个集群还是无效的。这是ZooKeeper防止“脑裂”默认采用的方法。
* 采用Redundant communications（冗余通信）方式：集群中采用多种通信方式，防止一种通信方式失效导致集群中的节点无法通信。
* Fencing（共享资源）方式：比如能看到共享资源就表示在集群中，能够获得共享资源的锁的就是Leader，看不到共享资源的，就不在集群中。
* 仲裁机制方式。
* 启动磁盘锁定方式。

## 小结

* 为了保证一致性，宁可不可用，也不要污染了数据。


# 参考链接

- [Zookeeper](https://www.runoob.com/w3cnote/zookeeper-tutorial.html)

- [ZooKeeper集群脑裂问题处理](https://cloud.tencent.com/developer/article/1758883)


