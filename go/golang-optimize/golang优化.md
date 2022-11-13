# golang优化

常见优化方法：异步、去锁、复用、零拷贝、批量，另外要说避免过早优化、业务逻辑层面的优化要先行

## cpu耗时优化

* make时提前预估size

* 临时的map、slice采用sync.Pool

* 大于32Kb也可用sync.Pool

* 不滥用goroutine，减少gc压力

* 不滥用mutex，减少上下文切换

* []byte与string临时变量转换用unsafe

* 减少reflect、defer使用

* atomic无锁使用

## 网络io性能优化

# 批量接口支持

# http 长连接

* redis pipeline

* db、redis连接池

* 增加缓存

* 大量数据压缩传输

## golang优化技巧

### array与slice

对于一些初学者，自知道 Go 里面的 array 以 pass-by-value 方式传递后，就莫名地引起 “恐慌”。外加诸多文章未作说明，就建议用 slice 代替 array，企图避免数据拷贝，提升性能。实际上，此做法有待商榷。某些时候怕会适得其反，倒造成不必要的性能损失。

### defer

编译器通过 runtime.deferproc “注册” 延迟调用，除目标函数地址外，还会复制相关参数（包括 receiver）。在函数返回前，执行runtime.deferreturn 提取相关信息执行延迟调用。这其中的代价自然不是普通函数调用一条 CALL 指令所能比拟的。

![image](https://user-images.githubusercontent.com/17688273/201530043-eec035c6-9af1-4b0a-b832-af78c2d2aa97.png)

解决方法么，要么去掉 f.close 前的 defer，要么将内层处理逻辑重构为独立函数（比如匿名函数调用）。

# 参考链接

- [Go性能优化技巧](https://blog.csdn.net/zhonglinzhang/article/details/71107168?share_token=7c9f28b2-f504-4bf2-a52b-6f034b02a4f9)

- [golang：快来抓住让我内存泄漏的“真凶”](https://mp.weixin.qq.com/s/FyHEiaa-UfyLStMKl2VFGA)

- [golang如何优化编译、逃逸分析、内联优化](https://mp.weixin.qq.com/s/tddRxcbzC1mB08C62br38Q)
