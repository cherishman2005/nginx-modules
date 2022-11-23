# golang优化

- [pprof命令](/go/golang-optimize/pprof.md)

## defer的使用原则

【建议4.1.5.1】`如果函数存在多个返回的地方，则采用defer来完成如关闭资源、解锁等清理操作。`

说明：Go的defer语句用来调度一个函数调用（被延期的函数），在函数即将返回之前defer才被运行。这是一种不寻常但又很有效的方法，用于处理类似于不管函数通过哪个执行路径返回，资源都必须要被释放的情况。典型的例子是对一个互斥解锁，或者关闭一个文件。

【建议4.1.5.2】`defer会消耗更多的系统资源，不建议用于频繁调用的方法中。`

【建议4.1.5.3】`避免在for循环中使用defer。`

说明：`一个完整defer过程要处理缓存对象、参数拷贝，以及多次函数调用，要比直接函数调用慢得多。`

## strings.Builder使用

* string：只读变量，对string操作会出现分配空间，拼接，拷贝。

* 采用strings.Builder提升性能


# 参考链接

- [通过 Exit Code 定位 Pod 异常退出原因](https://cloud.tencent.com/document/product/457/43125)
