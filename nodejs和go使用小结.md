# nodejs与go使用小结

## nodejs和go交互时针对json中的bigint怎么处理？

在互联网系统中，许多地方会用到唯一ID。使用int64类型相比字符串的优点是方便索引，这样比较高效。

一个go服务给前端提供了一个接口，返回json格式数据，其中Int64字段会超出javascript Number可表示的最大的Int值会丢精度，可以通过返回string类型值来解决问题。

**【解决方法1】**
Go服务再返回json的时候 以字符串形式返回值 即可解决这个问题。

**【解决方法2】**
采用json-bigint，bignumber.js/long.js等库

**【解决方法3】**

应用层协议设计：

如果要做到各端互通，js与其他语言rpc互通；最好采用protobuf设计应用层协议，如果一定要使用json，尽量采用int32，针对int64的id数据，尽量设计为字符串。



# 参考链接

- [https://www.mscto.com/go/10253.html](https://www.mscto.com/go/10253.html)

- [https://morioh.com/p/70a2ab2b1947](https://morioh.com/p/70a2ab2b1947)

- [https://github.com/gobwas/ws](https://github.com/gobwas/ws)

- [https://shockerli.net/post/go-awesome/](https://shockerli.net/post/go-awesome/)

- [https://github.com/catmullet/go-workers](https://github.com/catmullet/go-workers)
   

   * https://github.com/catmullet/go-workers/blob/master/workers.go
     
     go-workers - 安全地并发运行一组 worker，通过 channel 进行输入输出