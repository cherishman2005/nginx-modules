# defer的使用原则

【建议4.1.5.1】`如果函数存在多个返回的地方，则采用defer来完成如关闭资源、解锁等清理操作。`

说明：Go的defer语句用来调度一个函数调用（被延期的函数），在函数即将返回之前defer才被运行。这是一种不寻常但又很有效的方法，用于处理类似于不管函数通过哪个执行路径返回，资源都必须要被释放的情况。典型的例子是对一个互斥解锁，或者关闭一个文件。

【建议4.1.5.2】`defer会消耗更多的系统资源，不建议用于频繁调用的方法中。`

【建议4.1.5.3】`避免在for循环中使用defer。`

说明：`一个完整defer过程要处理缓存对象、参数拷贝，以及多次函数调用，要比直接函数调用慢得多。`

错误示例：实现一个加解锁函数，解锁过程使用defer处理。这是一个非常小的函数，并且能够预知解锁的位置，使用defer编译后会使处理产生很多无用的过程导致性能下降。
```
	var lock sync.Mutex
	func testdefer() {
	    lock.Lock()
	    defer lock.Unlock()
	}
	
	func BenchmarkTestDefer(b *testing.B) {
	    for i := 0; i < b.N; i++ {
	        testdefer()
	    }
	}
```
```
	// 耗时结果
	BenchmarkTestDefer 10000000 211 ns/op
```

推荐做法：`如果能够明确函数退出的位置，可以选择不使用defer处理。保证功能不变的情况下，性能明显提升，是耗时时使用defer的1/3。`
```
	var lock sync.Mutex
	func testdefer() {
	    lock.Lock()
	    lock.Unlock() // ## 【修改】去除defer
	}
	
	func BenchmarkTestDefer(b *testing.B) {
	    for i := 0; i < b.N; i++ {
	        testdefer()
	    }
	}
```

```
	// 耗时结果
	BenchmarkTest" 30000000 43.5 ns/op
```
