# golang优化

 * Go 1.13 改进了 sync 包中的 Pool，在 gc 运行时不会清除 pool。重写了逃逸分析，减少了 Go 程序中堆上的内存申请的空间。 
 * Go 1.14 进一步提升 defer 性能、页分配器更高效，同时 timer 也更高效。

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

### 字符串的拼接优先考虑bytes.Buffer

由于string类型是一个不可变类型，但拼接会创建新的string。GO中字符串拼接常见有如下几种方式：

string + 操作 ：导致多次对象的分配与值拷贝

fmt.Sprintf ：会动态解析参数，效率好不哪去

strings.Join ：内部是[]byte的append

bytes.Buffer ：可以预先分配大小，减少对象分配与拷贝

建议：对于高性能要求，优先考虑bytes.Buffer，预先分配大小。非关键路径，视简洁使用。fmt.Sprintf可以简化不同类型转换与拼接。

## 最小作用域

```
if err := DoSomething(); err != nil {
    return err
}
```
`尽量减少作用域, GC 比较友好`

## 减少内存拷贝

对于slice或者map等结构，如果不指定初始长度，使用类似与append的方法，系统会根据需要动态的增长内存容量，这样会导致内存的重新分配，增大gc压力，因此在可以预估容量时，可以考虑初始化固定长度，避免内存拷贝造成的开销。



# 参考链接

- [通过 Exit Code 定位 Pod 异常退出原因](https://cloud.tencent.com/document/product/457/43125)
