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

### 最小作用域

```
if err := DoSomething(); err != nil {
    return err
}
```
尽量减少作用域, GC 比较友好

【注】经常被调用的函数用了`取小作用域`优化，内存减少非常明显。-- 在某项目中内存减少5~10M。

### slice

* slice 缩容时，被缩掉对象如果不置 nil，是不会释放的

a = a[:1]，如果后面的元素都是指针，都指向了 500MB 的一个大 buffer，没法释放，GC 认为你还是持有引用的。这种情况需要自己先把后面的元素全置为 nil，再缩容。

### 减少逃逸的手段

尽量少用 fmt.Print、fmt.Sprint 系列的函数。

设计函数签名时，参数尽量少用 interface

少用闭包，被闭包引用的局部变量会逃逸到堆上。

### 使用strconv包替代fmt.Sprintf的格式化方式

### sync.Mutex

如果是struct，并且里面包含了sync.Mutex之类的同步原语，那么请使用*T，避免copy。

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

- [深入Go语言内部做技术优化](https://mp.weixin.qq.com/s/HuGudn8ViKXAz0mIawquyQ)

- [go1.14基于netpoll优化timer定时器实现原理](https://xiaorui.cc/archives/6483?share_token=70d6bd36-1f68-4c66-b042-525fced1e0c5)

- [Golang性能优化](https://blog.csdn.net/yonggeit/article/details/122393354?share_token=f71127c8-846f-4a7a-a7be-58ca7047684f)

- [golang服务器优化之旅](https://www.cnblogs.com/huangliang-hb/p/1115golang gc 优化思路以及实例分析398.html?share_token=4a4139ff-e0c5-4c46-a125-321ade2a62b6)

- [golang服务优化](https://blog.csdn.net/qq_27290011/article/details/121041355?share_token=024b50a3-9d1d-42a3-948f-edbe03dd976c)

- [golang gc 优化思路以及实例分析](https://www.cnblogs.com/gao88/p/9850235.html?share_token=ea789f77-71be-47a2-bedf-f4dbb9dddc15)

- [怎样优化一个 Go 服务以减少 40% 的 CPU 使用率](https://www.toutiao.com/article/6837035024143876611/?app=news_article&timestamp=1668239461&use_new_style=1&req_id=202211121551000101512042280D56B2F8&group_id=6837035024143876611&wxshare_count=1&tt_from=weixin&utm_source=weixin&utm_medium=toutiao_android&utm_campaign=client_share&share_token=2ed44c70-c62c-4101-8f3c-023132bdf6f4&source=m_redirect&wid=1668356410287)

- [Golang中常用的代码优化点](https://mp.weixin.qq.com/s/QONfbKioFf6VqJE2OwP7Kw)
