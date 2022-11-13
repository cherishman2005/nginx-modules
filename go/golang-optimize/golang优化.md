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


# 参考链接

- []
