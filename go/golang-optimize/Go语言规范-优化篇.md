# Go语言规范-优化篇 

说明：本篇的意义是为开发提供一些经过验证的开发规则和建议，让开发在开发过程中避免低级错误，从而提高代码的质量保证和性能效率

## 4.1 质量保证
### 4.1.1 代码质量保证优先原则
【原则4.1.1】代码质量保证优先原则：
（1）正确性，指程序要实现设计要求的功能。
（2）简洁性，指程序易于理解并且易于实现。
（3）可维护性，指程序被修改的能力，包括纠错、改进、新需求或功能规格变化的适应能力。
（4）可靠性，指程序在给定时间间隔和环境条件下，按设计要求成功运行程序的概率。
（5）代码可测试性，指软件发现故障并隔离、定位故障的能力，以及在一定的时间和成本前提下，进行测试设计、测试执行的能力。
（6）代码性能高效，指是尽可能少地占用系统资源，包括内存和执行时间。
（7）可移植性，指为了在原来设计的特定环境之外运行，对系统进行修改的能力。

### 4.1.2 对外接口原则

【原则4.1.2】对于主要功能模块抽象模块接口，通过interface提供对外功能。

说明：Go语言其中一个特殊的功能就是interface，它让面向对象，内容组织实现非常的方便。正确的使用这个特性可以使模块的可测试性和可维护性得到很大的提升。对于主要功能包（模块），在package包主文件中通过interface对外提供功能。

示例：在buffer包的buffer.go中定义如下内容
```
	package buffer
	
	import (
	    "policy_engine/models"
	)
	
	//other code …
	type MetricsBuffer interface {
	    Store(metric *DataPoint) error
	    Get(dataRange models.MatchPolicyDataRange) (*MetricDataBuf, error)
	    Clear(redisKey string) error
	    Stop()
	    Stats() []MetrisBufferStat
	    GetByKey(metricKey string) []DataPoint
	}
```
使用buffer package的代码示例，通过interface定义，可以在不影响调用者使用的情况下替换package。基于这个特性，在测试过程中，也可以通过实现符合interface要求的类来打桩实现测试目的。
```
	package metrics
	
	import (
	...//other import
	    "policy_engine/worker/metrics/buffer"
	)
	
	type MetricsClient struct {
	    logger            lager.Logger
	    redisClient       *store.RedisClient
	    conf              *config.Config
	    metricsBuffer     buffer.MetricsBuffer //interface类型定义的成员
	    metricsStatClient *metricstat.MetricsStatClient
	    stopSignal        chan struct{}
	}
	
	func New(workerId string, redisClient *store.RedisClient, logger lager.Logger, conf *config.Config) *MetricsClient {
	    var metricsBuffer MetricsBuffer
	    if conf.MetricsBufferConfig.StoreType == config.METRICS_MEM_STORE {
	        //具有interface定义函数的package实现，通过内存保存数据
	        metricsBuffer = NewMemBuffer(logger, conf)  
	    } else if conf.MetricsBufferConfig.StoreType == config.METRICS_REDIS_STORE {
	        //具有interface定义函数的package实现，通过redis保存数据
	        metricsBuffer = NewRedisBuffer(redisClient, logger, conf) 
	    } else {
	      ... //other code
	    }
	    ... //other code
	}
```

### 4.1.3 值与指针（T/*T）的使用原则

关于接收者对指针和值的规则是这样的，值方法可以在指针和值上进行调用，而指针方法只能在指针上调用。这是因为指针方法可以修改接收者；使用拷贝的值来调用它们，将会导致那些修改会被丢弃。

对于使用T还是*T作为接收者，下面是一些建议：

【建议4.1.3.1】基本类型传递时，尽量使用值传递。

【建议4.1.3.2】如果传递字符串或者接口对象时，建议直接实例传递而不是指针传递。

【建议4.1.3.3】如果是map、func、chan，那么直接用T。

【建议4.1.3.4】如果是slice，method里面不重新reslice之类的就用T。

【建议4.1.3.5】如果想通过method改变里面的属性，那么请使用*T。

【建议4.1.3.6】如果是struct，并且里面包含了sync.Mutex之类的同步原语，那么请使用*T，避免copy。

【建议4.1.3.7】如果是一个大型的struct或者array，那么使用*T会比较轻量，效率更高。

【建议4.1.3.8】如果是struct、slice、array里面的元素是一个指针类型，然后调用函数又会改变这个数据，那么对于读者来说采用*T比较容易懂。

【建议4.1.3.9】其它情况下，建议采用*T。

参考：https://github.com/golang/go/wiki/CodeReviewComments#pass-values

### 4.1.4 init的使用原则

每个源文件可以定义自己的不带参数的init函数，来设置它所需的状态。init是在程序包中所有变量声明都被初始化，以及所有被导入的程序包中的变量初始化之后才被调用。

除了用于无法通过声明来表示的初始化以外，init函数的一个常用法是在真正执行之前进行验证或者修复程序状态的正确性。

【规则4.1.4.1】一个文件只定义一个init函数。

【规则4.1.4.2】一个包内的如果存在多个init函数，不能有任何的依赖关系。

注意如果包内有多个init，每个init的执行顺序是不确定的。

### 4.1.5 defer的使用原则

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

推荐做法：如果能够明确函数退出的位置，可以选择不使用defer处理。保证功能不变的情况下，性能明显提升，是耗时时使用defer的1/3。
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
  
### 4.1.6 Goroutine使用原则

【规则4.1.6.1】确保每个goroutine都能退出。

说明：Goroutine是Go并行设计的核心，在实现功能时不可避免会使用到，执行goroutine时会占用一定的栈内存。

启动goroutine就相当于启动了一个线程，如果不设置线程退出的条件就相当于这个线程失去了控制，占用的资源将无法回收，导致内存泄露。

错误示例：示例中ready()启动了一个goroutine循环打印信息到屏幕上，这个goroutine无法终止退出。
```
	package main
	
	import (
	    "fmt"
	    "time"
	)
	
	func ready(w string, sec int) {    
	    go func() { // ## 【错误】goroutine启动之后无法终止
	        for {
	            time.Sleep(time.Duration(sec) * time.Second)
	            fmt.Println(w, "is ready! ")
	        }
	    }()
	}
	
	func main() {
	    ready("Tea", 2) 
	    ready("Coffee", 1)
	    fmt.Println("I'm waiting")
	    time.Sleep(5 * time.Second)
	}
```

推荐做法：对于每个goroutine都需要有退出机制，能够通过控制goroutine的退出，从而回收资源。通常退出的方式有：
 使用标志位的方式；
 信号量；
 通过channel通道通知；

注意：channel是一个消息队列，一个goroutine获取signal后，另一个goroutine将无法获取signal，以下场景下每个channel对应一个goroutine
```
	package main
	
	import (
	    "fmt"
	    "time"
	)
	
	func ready(w string, sec int, signal chan struct{}) {
	    go func() {
	        for {
	            select {
	            case <-time.Tick(time.Duration(sec) * time.Second):
	                fmt.Println(w, "is ready! ")
	            case <-signal: // 对每个goroutie增加一个退出选项 
	                fmt.Println(w, "is close goroutine!")
	                return
	            }
	        }
	    }()
	}
	
	func main() {
	    signal1 := make(chan struct{}) // 增加一个signal
	    ready("Tea", 2, signal1)
	
	    signal2 := make(chan struct{}) // 增加一个signal
	    ready("Coffee", 1, signal2)
	
	    fmt.Println("I'm waiting")
	    time.Sleep(4 * time.Second)
	    signal1 <- struct{}{}
	    signal2 <- struct{}{}
	    time.Sleep(4 * time.Second)
	}
 ```
 
【规则4.1.6.2】`禁止在闭包中直接引用闭包外部的循环变量。`

说明：Go语言的特性决定了它会出现其它语言不存在的一些问题，比如在循环中启动协程，当协程中使用到了循环的索引值，往往会出现意想不到的问题，通常需要程序员显式地进行变量调用。
```
	for i := 0; i < limit; i++ {
	    go func() { DoSomething(i) }()        //错误做法
	    go func(i int) { DoSomething(i)}(i)   //正确做法
	}
```
参考：http://golang.org/doc/articles/race_detector.html#Race_on_loop_counter

### 4.1.7 Channel使用原则

【规则4.1.7.1】传递channel类型的参数时应该区分其职责。

在只发送的功能中,传递channel类型限定为: c chan<- int
在只接收的功能中,传递channel类型限定为: c <-chan int

【规则4.1.7.2】确保对channel是否关闭做检查。

说明：在调用方法时不能想当然地认为它们都会执行成功，当错误发生时往往会出现意想不到的行为，因此必须严格校验并合适处理函数的返回值。例如：channel在关闭后仍然支持读操作，如果channel中的数据已经被读取，再次读取时会立即返回0值与一个channel关闭指示。如果不对channel关闭指示进行判断，可能会误认为收到一个合法的值。因此在使用channel时，需要判断channel是否已经关闭。

错误示例：下面代码中若cc已被关闭，如果不对cc是否关闭做检查，则会产生死循环。
```
	package main
	import (
	    "errors"
	    "fmt"
	    "time"
	)
	
	func main() {
	    var cc = make(chan int)
	    go client(cc)
	
	    for {
	        select {
	            case <-cc: //## 【错误】当channel cc被关闭后如果不做检查则造成死循环
	            fmt.Println("continue")
	            case <-time.After(5 * time.Second):
	            fmt.Println("timeout")
	        }
	    }
	}
	
	func client(c chan int) {
	    defer close(c)
	
	    for {
	        err := processBusiness()
	        if err != nil {
	            c <- 0
	            return
	        }
	        c <- 1
	    }
	}
	
	func processBusiness() error {
	    return errors.New("domo")
	}
```

推荐做法：对通道增加关闭判断。
```
	// 前面代码略……
	for {
	    select {
	    case _, ok := <-cc:
	        // 增加对chnnel关闭的判断，防止死循环
	        if ok == false {
	            fmt.Println("channel closed")
	            return
	        }
	        fmt.Println("continue")
	    case <-time.After(5 * time.Second):
	        fmt.Println("timeout")
	    }
	}
	// 后面代码略……
 ```
 
【规则4.1.7.3】禁止重复释放channel。

说明：重复释放channel会触发run-time panic，导致程序异常退出。重复释放一般存在于异常流程判断中，如果恶意攻击者能够构造成异常条件，则会利用程序的重复释放漏洞实施DoS攻击。

错误示例：
```
	func client(c chan int) {
	    defer close(c)
	    for {
	        err := processBusiness()

	        if err != nil {
	            c <- 0
	            close(c) // ## 【错误】可能会产生双重释放
	            return
	        }
	        c <- 1
	    }
	}
```
推荐做法：确保创建的channel只释放一次。
```
	func client(c chan int) {
	    defer close(c)
	
	    for {
	        err := processBusiness()
	        if err != nil {
	            c <- 0     // ## 【修改】使用defer延迟close后，不再单独进行close
	            return
	        }
	        c <- 1
	    }
	}
```

### 4.1.8 其它
【建议4.1.8.1】使用go vet --shadow检查变量覆盖，以避免无意的变量覆盖。

GO的变量赋值和声明可以通过”:=”同时完成，但是由于Go可以初始化多个变量，所以这个语法容易引发错误。下面的例子是一个典型的变量覆盖引起的错误，第二个val的作用域只限于for循环内部，赋值没有影响到之前的val。
```
	package main
	
	import "fmt"
	import "strconv"
	
	func main() {
	    var val int64
	
	    if val, err := strconv.ParseInt("FF", 16, 64); nil != err {
	        fmt.Printf("parse int failed with error %v\n", err)
	    } else {
	        fmt.Printf("inside  : val is %d\n", val)
	    }
	    fmt.Printf("outside : val is %d \n", val)
	}
```

	执行结果：
```
	inside  : val is 255
	outside : val is 0
```

正确的做法：
```
	package main
	
	import "fmt"
	import "strconv"
	
	func main() {
	    var val int64
	    var err error
	
	    if val, err = strconv.ParseInt("FF", 16, 64); nil != err {
	        fmt.Printf("parse int failed with error %v\n", err)
	    } else {
	        fmt.Printf("inside  : val is %d\n", val)
	    }
	    fmt.Printf("outside : val is %d \n", val)
	}
```	
	执行结果：
```
	inside  : val is 255
	outside : val is 255
```

【建议4.1.8.2】GO的结构体中控制使用Slice和Map。

GO的slice和map等变量在赋值时，传递的是引用。从结果上看，是浅拷贝，会导致复制前后的两个变量指向同一片数据。这一点和Go的数组、C/C++的数组行为不同，很容易出错。
```
	package main
	import "fmt"
	
	type Student struct {
	    Name     string
	    Subjects []string
	}
	
	func main() {
	    sam := Student{
	        Name: "Sam", Subjects: []string{"Math", "Music"},
	    }
	    clark := sam //clark.Subject和sam.Subject是同一个Slice的引用！
	    clark.Name = "Clark"
	    clark.Subjects[1] = "Philosophy" //sam.Subject[1]也变了！
	    fmt.Printf("Sam : %v\n", sam)
	    fmt.Printf("Clark : %v\n", clark)
	}
```	
	执行结果：
```
	Sam : {Sam [Math Philosophy]}
	Clark : {Clark [Math Philosophy]}
```
作为对比，请看作为Array定义的Subjects的行为：
```
	package main
	import "fmt"
	
	type Student struct {
	    Name     string
	    Subjects [2]string
	}
	
	func main() {
	    var clark Student
	    sam := Student{
	        Name: "Sam", Subjects: [2]string{"Math", "Music"},
	    }
	
	    clark = sam //clark.Subject和sam.Subject不同的Array
	    clark.Name = "Clark"
	    clark.Subjects[1] = "Philosophy" //sam.Subject不受影响！
	    fmt.Printf("Sam : %v\n", sam)
	    fmt.Printf("Clark : %v\n", clark)
	}
```

执行结果：
```  
	Sam : {Sam [Math Music]}
	Clark : {Clark [Math Philosophy]}
```

编写代码时，建议这样规避上述问题：
 结构体内尽可能不定义Slice、Maps成员；
 如果结构体有Slice、Maps成员，尽可能以小写开头、控制其访问；
 结构体的赋值和复制，尽可能通过自定义的深度拷贝函数进行；

【规则4.1.8.3】避免在循环引用调用 runtime.SetFinalizer。

说明：指针构成的 "循环引用" 加上 runtime.SetFinalizer 会导致内存泄露。

runtime.SetFinalizer用于在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中。在对象被 GC 进程选中并从内存中移除以前，SetFinalizer 都不会执行，即使程序正常结束或者发生错误。

错误示例：垃圾回收器能正确处理 "指针循环引用"，但无法确定 Finalizer 依赖次序，也就无法调用Finalizer 函数，这会导致目标对象无法变成不可达状态，其所占用内存无法被回收。
```
	package main
	
	import (
	    "fmt"
	    "runtime"
	    "time"
	)
	
	type Data struct {
	    d [1024 * 100]byte
	    o *Data
	}
	
	func test() {
	    var a, b Data
	    a.o = &b
	    b.o = &a
	
	    // ## 【错误】循环和SetFinalize同时使用
	    runtime.SetFinalizer(&a, func(d *Data) { fmt.Printf("a %p final.\n", d) })
	    runtime.SetFinalizer(&b, func(d *Data) { fmt.Printf("b %p final.\n", d) })
	}
	
	func main() {    
	    for { // ## 【错误】循环和SetFinalize同时使用
	        test()
	        time.Sleep(time.Millisecond)
	    }
	}
```
通过跟踪GC的处理过程，可以看到如上代码内存在不断的泄露：
go build -gcflags "-N -l" && GODEBUG="gctrace=1" ./test
gc11(1): 2+0+0 ms, 104 -> 104 MB 1127 -> 1127 (1180-53) objects
gc12(1): 4+0+0 ms, 208 -> 208 MB 2151 -> 2151 (2226-75) objects
gc13(1): 8+0+1 ms, 416 -> 416 MB 4198 -> 4198 (4307-109) objects
以上结果标红的部分代表对象数量，我们在代码中申请的对象都是局部变量，在正常处理过程中GC会持续的回收局部变量占用的内存。但是在当前的处理过程中，内存无法被GC回收，目标对象无法变成不可达状态。

推荐做法：需要避免内存指针的循环引用以及runtime.SetFinalizer同时使用。

【规则4.1.8.4】避免在for循环中使用time.Tick()函数。

如果在for循环中使用time.Tick()，它会每次创建一个新的对象返回，应该在for循环之外初始化一个ticker后，再在循环中使用：
```
	ticker := time.Tick(time.Second)
	for {
	    select {
	        case <-ticker:
	        // …
	    }
	}
```

## 4.2 性能效率
4.2.1 Memory优化
【建议4.2.1.1】将多次分配小对象组合为一次分配大对象。

比如, 将 *bytes.Buffer 结构体成员替换为bytes。缓冲区 (你可以预分配然后通过调用bytes.Buffer.Grow为写做准备) 。这将减少很多内存分配(更快)并且减缓垃圾回收器的压力(更快的垃圾回收) 。

【建议4.2.1.2】`将多个不同的小对象绑成一个大结构，可以减少内存分配的次数。`

比如：将
```
	for k, v := range m {
	   k, v := k, v   // copy for capturing by the goroutine
	   go func() {
	     // use k and v
	   }()
	}
```
替换为:
```
	for k, v := range m {
	   x := struct{ k, v string }{k, v}   // copy for capturing by the goroutine

	   go func() {

	       // use x.k and x.v
	   }()
	}
```
这就将多次内存分配（分别为k、v分配内存）替换为了一次（为x分配内存）。然而，这样的优化方式会影响代码的可读性，因此要合理地使用它。

【建议4.2.1.3】组合内存分配的一个特殊情形是对分片数组进行预分配。

如果清楚一个特定的分片的大小，可以对数组进行预分配：
```
	type X struct {
	    buf      []byte
	    bufArray [16]byte // Buf usually does not grow beyond 16 bytes.
	}
	
	
	func MakeX() *X {
	    x := &X{}
	    // Preinitialize buf with the backing array.
	    x.buf = x.bufArray[:0]
	    return x
	}
 ```
 
【建议4.2.1.4】尽可能使用小数据类型，并尽可能满足硬件流水线（Pipeline）的操作，如对齐数据预取边界。

说明：不包含任何指针的对象(注意 strings,slices,maps 和 chans 包含隐含指针)不会被垃圾回收器扫描到。

比如，1GB 的分片实际上不会影响垃圾回收时间。因此如果你删除被频繁使用的对象指针，它会对垃圾回收时间造成影响。一些建议：使用索引替换指针，将对象分割为其中之一不含指针的两部分。

【建议4.2.1.5】使用对象池来重用临时对象，减少内存分配。

标准库包含的sync.Pool类型可以实现垃圾回收期间多次重用同一个对象。然而需要注意的是，对于任何手动内存管理的方案来说，不正确地使用sync.Pool会导致 use-after-free bug。

###4.2.2 GC 优化

【建议4.2.2.1】设置GOMAXPROCS为CPU的核心数目，或者稍高的数值。

GC是并行的，而且一般在并行硬件上具有良好可扩展性。所以给 GOMAXPROCS 设置较高的值是有意义的，就算是对连续的程序来说也能够提高垃圾回收速度。但是，要注意，目前垃圾回收器线程的数量被限制在 8 个以内。

【建议4.2.2.2】避免频繁创建对象导致GC处理性能问题。

说明：尽可能少的申请内存，减少内存增量，可以减少甚至避免GC的性能冲击，提升性能。
Go语言申请的临时局部变量（对象）内存，都会受GC（垃圾回收）控制内存的回收，其实我们在编程实现功能时申请的大部分内存都属于局部变量，所以与GC有很大的关系。

Go在GC的时候会发生Stop the world，整个程序会暂停，然后去标记整个内存里面可以被回收的变量，标记完成之后再恢复程序执行，最后异步地去回收内存。（暂停的时间主要取决于需要标记的临时变量个数，临时变量数量越多，时间越长。Go 1.7以上的版本大幅优化了GC的停顿时间， Go 1.8下，通常的GC停顿的时间<100μs）

目前GC的优化方式原则就是尽可能少的声明临时变量：
 局部变量尽量利用
 如果局部变量过多，可以把这些变量放到一个大结构体内，这样扫描的时候可以只扫描一个变量，回收掉它包含的很多内存

本规则所说的创建对象包含：
 &obj{}
 new(abc{})
 make()

我们在编程实现功能时申请的大部分内存都属于局部变量，下面这个例子说明的是我们实现功能时需要注意的一个问题，适当的调整可以减少GC的性能消耗。

错误示例：
代码中定义了一个tables对象，每个tables对象里面有一堆类似tableA和tableC这样的一对一的数据，也有一堆类似tableB这样的一对多的数据。假设有1万个玩家，每个玩家都有一条tableA和一条tableC的数据，又各有10条tableB的数据，那么将总的产生1w (tables) + 1w (tableA) + 1w (tableC) + 10w (tableB)的对象。

不好的例子：
```
	// 对象数据表的集合
	type tables struct {
	    tableA *tableA
	    tableB *tableB
	    tableC *tableC
	    // 此处省略一些表
	}
	
	// 每个对象只会有一条tableA记录
	type tableA struct {
	    fieldA int
	    fieldB string
	}
	
	// 每个对象有多条tableB记录
	type tableB struct {
	    city string
	    code int
	    next *tableB // 指向下一条记录
	}
	
	// 每个对象只有一条tableC记录
	type tableC struct {
	    id    int
	    value int64
	}
```
建议一对一表用结构体，一对多表用slice，每个表都加一个_is_nil的字段，用来表示当前的数据是否是有用的数据，这样修改的结果是，一万个玩家，产生的对象总量是1w（tables）+1w([]tablesB)，跟前面的差别很明显：
```
	// 对象数据表的集合
	type tables struct {
	        tableA tableA
	        tableB []tableB
	        tableC tableC
	    // 此处省略一些表
	}
	
	// 每个对象只会有一条tableA记录
	type tableA struct {
	    _is_nil bool 
	    fieldA  int
	    fieldB  string
	}
	

	// 每个对象有多条tableB记录
	type tableB struct {
	    _is_nil bool 
	    city    string
	    code    int
	    next *tableB // 指向下一条记录
	}
	
	// 每个对象只有一条tableC记录
	type tableC struct {
	    _is_nil bool
	    id      int
	    value   int64
	}
```

4.2.3 其它优化建议
【建议4.2.3.1】减少[]byte和string之间的转换，尽量使用[]byte来处理字符。

说明：Go里面string类型是immutable类型，而[]byte是切片类型，是可以修改的，所以Go为了保证语法上面没有二义性，在string和[]byte之间进行转换的时候是一个实实在在的值copy，所以我们要尽量的减少不必要的这个转变。

下面这个例子展示了传递slice但是进行了string的转化，
```
	func PrefixForBytes(b []byte) string {
	        return "Hello" + string(b)
	}
```
所以我们可以有两种方式，一种是保持全部的都是slice的操作，如下：
```
	func PrefixForBytes(b []byte) []byte {
	    return append([]byte(“Hello”,b…))
	}
```
还有一种就是全部是string的操作方式
```
	func PrefixForBytes(str string) string {
	        return "Hello" + str
	}
```

推荐阅读：https://blog.golang.org/strings

【建议4.2.3.2】make申请slice/map时，根据预估大小来申请合适内存。

说明：map和数组不同，可以根据新增的<key,value>对动态的伸缩，因此它不存在固定长度或者最大限制。

map的空间扩展是一个相对复杂的过程，每次扩容会增加到上次大小的两倍。它的结构体中有一个buckets和oldbuckets，用来实现增量扩容，正常情况下直接使用buckets，oldbuckets为空，如果当前哈希表正在扩容，则oldbuckets不为空，且buckets大小是oldbuckets大小的两倍。对于大的map或者会快速扩张的map，即便只是大概知道容量，也最好先标明。

slice是一个C语言动态数组的实现，在对slice进行append等操作时，可能会造成slice的自动扩容，其扩容规则：
 如果新的大小是当前大小2倍以上，则大小增长为新大小
 否则循环以下操作：如果当前大小小于1024，按每次2倍增长，否则每次按当前大小1/4增长，直到增长的大小超过或者等于新大小

推荐做法：在初始化map时指明map的容量。
```
map := make(map[string]float, 100)
```
【建议4.2.3.3】字符串拼接优先考虑bytes.Buffer。

Golang字符串拼接常见有如下方式：
 fmt.Sprintf
 strings.Join
 string +
 bytes.Buffer

fmt.Sprintf会动态解析参数，效率通常是最差的，而string是只读的，string+会导致多次对象分配与值拷贝，而bytes.Buffer在预设大小情况下，通常只会有一次拷贝和分配，不会重复拷贝和复制，故效率是最佳的。

推荐做法：优先使用bytes.Buffer，非关键路径，若考虑简洁，可考虑其它方式，比如错误日志拼接使用fmt.Sprintf，但接口日志使用就不合适。

【建议4.2.3.4】避免使用CGO或者减少跨CGO调用次数。

说明：Go可以调用C库函数，但是Go带有垃圾收集器且Go的栈是可变长，跟C实际是不能直接对接的，Go的环境转入C代码执行前，必须为C新创建一个新的调用栈，把栈变量赋值给C调用栈，调用结束后再拷贝回来，这个调用开销非常大，相比直接GO语言调用，单纯的调用开销，可能有2个甚至3个数量级以上，且Go目前还存在版本兼容性问题。

推荐做法：尽量避免使用CGO，无法避免时，要减少跨CGO调用次数。

【建议4.2.3.5】避免高并发调用同步系统接口。

说明：编程世界同步场景更普遍，GO提供了轻量级的routine，用同步来模拟异步操作，故在高并发下的，相比线程，同步模拟代价比较小，可以轻易创建数万个并发调用。然而有些API是系统函数，而这些系统函数未提供异步实现，程序中最常见的posix规范的文件读写都是同步，epoll异步可解决网络IO，而对regular file是无法工作的。Go的运行时环境不可能提供超越操作系统API的能力，它依赖于系统syscall文件中暴露的api能力，而1.6版本还是多线程模拟，线程创建切换的代价也非常巨大，开源库中有filepoller来模拟异步其实也基于这两种思路，效率上也会大打折扣。

推荐做法：把诸如写文件这样的同步系统调用，要隔离到可控的routine中，而不是直接高并发调用。

【建议4.2.3.6】高并发时避免共享对象互斥。

说明：在Go中，可以轻易创建10000个routine而对系统资源通常就是100M的内存要求，但是并发数多了，在多线程中，当并发冲突在4个到8个线程间时，性能可能就开始出现拐点，急剧下降，这同样适应于Go，Go可以轻易创建routine，但对并发冲突的风险必须要做实现的处理。

推荐做法：routine需要是独立的，无冲突的执行，若routine间有并发冲突，则必须控制可能发生冲突的并发routine个数，避免出现性能恶化拐点。

【建议4.2.3.7】长调用链或在函数中避免申明较多较大临时变量。

routine的调用栈默认大小1.7版本已修改为2K，当栈大小不够时，Go运行时环境会做扩栈处理，创建10000个routine占用空间才20M，所以routine非常轻量级，可以创建大量的并发执行逻辑。而线程栈默认大小是1M，当然也可以设置到8K（有些系统可以设置4K），一般不会这么做，因为线程栈大小是固定的，不能随需而变大，不过实际CPU核一般都在100以内，线程数是足够的。

routine是怎么实现可变长栈呢？当栈大小不够时，它会新创建一个栈，通常是2倍大小增长，然后把栈赋值过来，而栈中的指针变量需要搜索出来重新指向新的栈地址，好处不是随便有的，这里就明显有性能开销，而且这个开销不小。

说明：频繁创建的routine，要注意栈生长带来的性能风险，比如栈最终是2M大小，极端情况下就会有数10次扩栈操作，从而让性能急剧下降。所以必须控制调用栈和函数的复杂度，routine就意味着轻量级。

对于比较稳定的routine，也要注意它的栈生长后会导致内存飙升。

【建议4.2.3.8】为高并发的轻量级任务处理创建routine池。

说明：Routine是轻量级的，但对于高并发的轻量级任务处理，频繁创建routine来执行，执行效率也是非常低效率的。

推荐做法：高并发的轻量级任务处理，需要使用routine池，避免对调度和GC带来冲击。

【建议4.2.3.9】建议版本提供性能/内存监控的功能，并动态开启关闭，但不要长期开启pprof提供的CPU与MEM profile功能。

Go提供了pprof工具包，可以运行时开启CPU与内存的profile信息，便于定位热点函数的性能问题，而MEM的profile可以定位内存分配和泄漏相关问题。开启相关统计，跟GC一样，也会严重干扰性能，因而不要长期开启。

推荐做法：做测试和问题定位时短暂开启，现网运行，可以开启短暂时间收集相关信息，同时要确保能够自动关闭掉，避免长期打开。
