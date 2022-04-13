# Go中map全局变量

现阶段的go标准库里的map是动态扩展的。const语义只支持简单的原生类型：例如int，string。这点和c++里的const语义不同。如果你想让map的内容保持const，只能自己定义一个结构来实现类似的操作。

## Go中map全局变量使用原则

原则一：如果没有写操作，map可以直接定义成为全局变量即可。
原则二：如果存在读写操作，
解决方案如下：
* 1、使用sync.Map
* 2、使用map+sync.Mutex
* 3、使用concurrent-map

sync.Map和concurrent-map对比
sync.Map适用于读多写少的场景，和sync.Map使用的两个map的实现机制有关，因为某些情况下会操作dirty map。
concurrent-map实现方式是通过将大的map拆分成小map，将大锁拆分成小锁的解决方案，从而降低了锁粒度，减少锁冲突。
关于map的不堪往事：
之前在公司将ruby的代码迁移到Go实现，ruby中有很多的字典格式的变量，迁移时没有弄明白map的使用方式，最终将变量迁移到Go代码中时将字典变量定义成了数组，导致了后面出现了很多的数组越界的报错。因为其他微服务可能增加了变量的类型而自己的代码没有兼容，就panic了。
于是之前直接使用map的全局数据结构即可。

# 参考链接

- [https://juejin.cn/post/6944516385682767880](https://juejin.cn/post/6944516385682767880)
- [https://www.zhihu.com/question/25953192](https://www.zhihu.com/question/25953192)
