# 浅谈 golang 代码规范, 性能优化和需要注意的坑

## 编码规范

[强制] 声明slice
申明 slice 最好使用

var t []int
而不是使用
```
t := make([]int, 0)
```
因为 var 并没有初始化，但是 make 初始化了。

但是如果要指定 slice 的长度或者 cap，可以使用 make

最小作用域
```
if err := DoSomething(); err != nil {
    return err
}
```
尽量减少作用域, GC 比较友好

赋值规范
声明一个对象有4种方式：make, new(), var, :=

比如:
```
t := make([]int, 0)
u := new(User)
var t []int
u := &User{}
```
var 声明但是不立刻初始化
:= 声明并立刻使用
尽量减少使用 new() 因为他不会初始化值, 使用 u := User{} 更好
接口命名
单个功能使用 er 结尾或者名词
```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
2 个功能
```
type ReaderWriter interface {
    Reader
    Writer
}
```
3 个及以上功能
```
type Car interface {
    Drive()
    Stop()
    Recover()
}
```
命名规范
代码风格
[强制] go 文件使用下划线命名

[强制] 常量使用下划线或者驼峰命名, 表达清除不要嫌名字太长

[推荐] 不要在名字中携带类型信息

// 反例
userMap := map[string]User{}
// 正例
users := map[string]User{}
[推荐] 方法的参数要能表达含义

// 反例
func CopyFile(a, b string) error
// 正例
func CopyFile(src, dst string) error
[推荐] 包名一律使用小写字母, 不要加下划线或者中划线

[推荐] 如果使用了设计模式, 名称中体现设计模式的含义

type AppFactory interface {
    CreateApp() App
}
[推荐] 如果变量名是 bool 类型, 如果字段名不能表达 bool 类型, 可以使用 is 或者 has 前缀

var isDpdk bool
[强制] 一个变量只能有一个功能, 并和名称一致, 不要把一个变量作为多种用途

// 反例
length := len(userList)
length := len(orderList)
// 正例
userNum := len(userList)
orderNum := len(orderList)
[推荐] 如果变量名是 bool 类型, 如果字段名不能表达 bool 类型, 可以使用 is 或者 has 前缀

golang 基本规范
包设计
[强制] 包的设计满足单一职责

说明: 在 SRP (Single Response Principle) 模式中, 单一职责原则是指一个类只负责一项功能, 并且不能负责多个功能. 将包设计的非常内聚, 减少包之间的 api

[强制] 包的设计遵循最小可见性原则

说明: 仅在包内调用的函数, 或者变量, 或者结构体, 或者接口, 或者类型, 或者函数等等, 需要小写开头, 不可以可以被外部包访问

[强制] 代码需要可测试性, 使用接口和依赖注入替代硬编码

[强制] 单元测试文件放到代码文件同级目录, 便于 golang 工具使用

比如: vscode 在方法上右键可以直接生成测试代码和测试覆盖率并可视化展示执行情况

布局
[推荐] 程序实体之间使用空行区分, 增加可读性

说明: 比如函数中各个模块功能使用空行区分, 增加可读性

[推荐] 每个文件末尾应该有且仅有一个空行

[推荐] 一元操作符不要加空格, 二元操作符的才需要

注释
[推荐] 可导出的方法, 变量, 结构体等都需要注释

表达式和语句
[推荐] if 或者循环的嵌套层数不宜大于 3

[推荐] 对于 for 遍历, 优先使用 range 而不是显式'的下标, 如果 value 占用内存大的话可以使用显式下标

说明: range 可以让代码更加整洁, 特别是多层 for 嵌套的时候, 但是 range 非拷贝值, 如果 value 不是指针类型, 而且占用内存较大会有性能损耗.

函数
[强制] 命名不要暴露实现细节, 一般以"做什么"来命名而不是"怎么做"

[推荐] 短小精悍, 尽量控制到 20 行左右

说明: 函数的粒度越小, 可复用的概率越大, 而且函数越多, 高层函数调用起来就是代码可读性很高, 读起来就像一系列解释

[推荐] 单一职责, 函数只做好一件事情, 只做一件事情

[强制] 不要设置多功能函数

例如: 一个函数既修改了状态, 又返回了状态, 应该拆分

[推荐] 为简单的功能编写函数

说明: 为 1,2 行代码编写函数也是必要的, 增加代码的可复用性, 增加高层函数的可读性, 可维护性, 可测试性

参数
[推荐] 参数个数限制在 3 个以内, 如果超过了, 可以使用配置类或者 Options 设计模式

说明: 函数的最理想的参数的个数首先是0, 然后是 1, 然后是 2, 3 就会很差了. 因为参数有很多概念性, 可读性差, 而且让测试十分复杂

[推荐] 函数不能含有表示参数

说明: 标识参数丑陋不堪, 函数往往根据标识参数走不同的逻辑, 这个和单一职责违背

[强制] struct 作为参数传递的时候, 使用指针

说明: 函数的执行就是压栈, struct 如果有多个字段将会被多次压栈, 有性能损失, 指针只会被压栈一次

[推荐] 在 api(controller) 层对传入的参数进行检查, 而不是每一层都检查一次

[推荐] 当 chan 作为函数的参数的时候, 根据最小权限原则, 使用单向 chan

// 读取单向chan
func Parse(ch <-chan struct{}) {
	for v := range ch {
		println(v)
	}
}

// 写入单向chan
func Do(down chan<- struct{}) {
	time.Sleep(time.Second)
	down <- struct{}{}
}
返回值
[推荐] 返回值的个数不要大于 3

[强制] 统一定义错误, 不要随便抛出错误

说明: 比如记录不存在可能有多种错误

"record not exits"
"record not exited"
"record not exited!!"
上层函数要处理底层的错误的话, 要知道所有的抛出情况, 这个是不现实的, 需要处理的错误应该使用统一文件定义错误码

[强制] 没有失败原因的时候, 不要使用 error

// 正例
func IsPhone() bool
// 反例
func IsPhone() error
[推荐] 当多重试几次可以避免失败的时候, 不要返回 error

错误是偶然发生的, 应该给一个机会重试, 可以避免大多数的偶然问题

[推荐] 上层函数不关心 error 的时候,  不要返回 error

比如 Close(), Clear() 抛出了 error, 上层函数大概率不知道怎么处理

异常设计
[推荐] 程序的开发阶段, 坚持速错, 让异常程序崩溃

说明: 速错的本质逻辑就是 "让它挂", 只有挂了你才第一时间知道错误, panic 能让 bug 尽快被修复

[强制] 程序部署后, 应该避免终止

是否 recover 应该根据配置文件确定, 默认需要 recover

注意: 有时候需要在延迟函数中释放资源, 比如 panic 之前 read 了 channel, 但是还没有 write 就 panic , 需要在 deffer 函数中做好处理, 防止 channel 阻塞.

[推荐] 当入参不合法的时候, panic

说明: 当入参不合法的时候, panic, 可以让上层函数知道错误, 而不是继续执行(api 应该提前做好参数检查)

整洁测试
[强制] 不要为了测试对代码进行入侵式的修改, 应该 mock

说明: 禁止为了测试在函数中增加条件分支和测试变量

[推荐] 测试的三要数, 可读性, 可读性, 可读性

生产代码的可靠性由测试代码来保证, 测试代码的可靠性由最简单的可读性来保证, 逻辑需要简单到没有 bug

REFERENCE
bilibili  go 规范

uber go-guide https://github.com/xxjwxc/uber_go_guide_cn

golang 性能优化
内存优化
小对象合并
小对象在堆内存上频繁的创建和销毁, 会导致内存碎片, 一般会才使用内存池

golang 的内存机制也是内存池, 每个 span 大小为 4KB, 同时维护一个 cache, cache 有一个 list 数组

数组里面储存的是链表, 就像 HashMap 的拉链法, 数组的每个格子代表的内存大小是不一样的, 64 位的机器是 8 byte 为基础, 比如下标 0 是 8 byte 大小的链表节点, 下标 1 是 16 byte 的链表节点, 每个下标的内存不一样, 使用的是按需分配最近的内存, 比如一个结构体的内存实际上算下来是 31 byte, 分配的时候会分配 32 byte.

一个下标的一条链表的每个 Node 储存的内存是一致的.

所以建议将小对象合并为一个 struct
```
for k, v := range m {
    x := struct {k , v string} {k, v} // copy for capturing by the goroutine
    gofunc() {
        // using x.k & x.v
    }()
}
```

使用 buf 缓存
协议编码的时候需要频繁的操作 buf, 可以使用 bytes.Buffer 作为缓存区对象, 它会一次性分配足够大的内存, 避免内存不够的时候动态申请内存, 减少内存分配次数, 而且, buf 可以被复用(建议复用)

slice 和 map 创建的时候, 预估大小指定的容量
预先分配内存, 可以减少动态扩容带来的开销

t := make([]int, 0, 100)
m := make(map[string]int, 100)
如果不确定 slice 会不会初始化, 使用 var 这样不会分配内存, make([]int,0) 会分配内存空间
```
var t []int
```
拓展:

slice 容量在 1024 前扩容是倍增, 1024 后是1/4

map 的扩容机制比较复杂, 每次扩容是 2 倍数, 结构体中有一个 bucket 和 oldBuckets 实现增量扩容

长调用栈避免申请较多的临时对象
说明: goroutine 默认的 栈的大小是 4K, 1.7 改为 2K, 它采用的是连续栈的机制, 当栈空间不够的时候, goroutine 会不断扩容, 每次扩容就先 slice 的扩容一样, 设计新的栈空间申请和旧栈空间的拷贝, 如果 GC 发现现在的空间只有之前的 1/4 又会缩容, 频繁的内存申请和拷贝会带来开销

建议: 控制函数调用栈帧的复杂度, 避免创建过多的临时对象, 如果确实需要比较长的调用栈或者 job 类型的代码, 可以考虑将 goroutine 池化

避免频繁创建临时变量
说明: GC STW 的时间已经优化到最糟糕 1ms 内了, 但是还是有混合写屏障会降低性能, 如果临时变量个数太多, GC 性能损耗就高.

建议: 降低变量的作用域, 使用局部变量, 最小可见性, 将多个变量合并为一个 struct 数组(降低扫描次数)

大的 struct 使用指针传递
golang 都是值拷贝, 特别是 struct 入栈帧的时候会将变量一个一个入栈, 频繁申请内存, 可以使用指针传递来优化性能

并发优化
goroutine 池化
go 虽然轻量, 但是对于高并发的轻量级任务, 比如高并发的 job 类型的代码, 可以考虑使用 goroutine 池化, 减少 goroutine 的创建和销毁, 减少 goroutine 的创建和销毁的开销

减少系统调用
goroutine 的实现是通过同步模拟异步操作, 比如下面的操作并不会阻塞, runtime 的线程调度

网络IO
channel
time.Sleep
基于底层异步的 SysCall
下面的阻塞会创建新的线程调度

本地 IO
基于底层同步的 SysCall
CGO 调用 IO 或者其他阻塞
建议将同步调用: 隔离到可控 goroutine 中, 而不是直接高并 goroutine 调用

减少锁, 减少大锁
Go 推荐使用 channel 的方式调用而不是共享内存, channel 之间存在大锁, 可以将锁的力度降低

拓展: channel

channel 不要传递大数据, 会有值拷贝
channel 的底层是链表 + 锁

不要用 channel 传递图片等数据, 任何的队列的性能都很低, 可以尝试指针优化大对象

合并请求 singleflight
参考: singleflight

协议压缩 protobuf
protobuf 比 json 的储存效率和解析效率更高, 推荐在持久化或者数据传输的时候使用 protobuf 替代 json

批量协议
对数据访问接口提供批量协议, 比如门面设计模式或者 pipeline, 可以减少非常多的 IO, QPS, 和拆包解包的开销

并行请求 errgroup
对于网关接口, 通常需要聚合多个模块的数据, 当这些业务模块数据之间没有依赖的时候, 可以并行请求, 减少耗时
```
ctxTimeout, cf := context.WithTimeout(context.Background(), time.Second)
defer cf()
g, ctx := errgroup.WithContext(ctxTimeout)
var urls = []string{
    "http://www.golang.org/",
    "http://www.google.com/",
    "http://www.somestupidname.com/",
}
for _, url := range urls {
    // Launch a goroutine to fetch the URL.
    url := url // https://golang.org/doc/faq#closures_and_goroutines
    g.Go(func() error {
        // Fetch the URL.
        resp, err := http.Get(url)
        if err == nil {
            resp.Body.Close()
        }
        return err
    })
}
// Wait for all HTTP fetches to complete.
if err := g.Wait(); err == nil {
    fmt.Println("Successfully fetched all URLs.")
}
select {
case <-ctx.Done():
    fmt.Println("Context canceled")
default:
    fmt.Println("Context not canceled")
}
```
其他优化
需要注意的坑
channel 之坑
如何优雅的关闭 channel
参考: 如何优雅的关闭 channel

关闭 channel 的坑
关闭已经关闭的 channel 会导致 panic
给关闭的 channel 发送数据会导致 panic
从关闭的 channel 中读取数据是初始值默认值
CCP 原则
CCP: Channel Close Principle (关闭通道原则)

不要从接收端关闭 channel
不要关闭有多个发送端的 channel
当发送端只有一个且后面不会再发送数据才可以关闭 channel
有缓存的 channel 不一定有序
defer 之坑
defer 中的变量
参数传递是在调用的时候
```
i := 1
deferprintln("defer", i)
i++
// defer 1
```

非参数的闭包
```
i := 1
deferfunc() {
    println("defer", i)
}()
i++
// defer 2
```
有名返回同理闭包, 并且会修改有名返回的返回值
```
func main(){
	fmt.Printf("main: %v\n", getNum())
	// defer 2
	// main: 2
}

func getNum() (i int) {
	deferfunc() {
		i++
		println("defer", i)
	}()
	i++
	return
}
```

`不要 for 循环中调用 defer`

`因为 deffer 只会在函数 return 之后执行, 这样会累积大量的 defer 而且极其容易出错`

建议: 将 for 循环需要 defer 的代码逻辑封装为一个函数

HTTP 之坑
request 超时时间
golang 的 http 默认的 request 没有超时, 这是一个大坑, 因为如果服务器没有响应, 也没有断开, 客户端会一直等待, 导致客户端阻塞, 量一上来就崩溃了

关闭 HTTP 的 response
http 请求框架的 response 一定要通过 Close 方法关闭, 不然有可能内存泄露

## interface 之坑

interface 到底什么才等于 nil?
说明: interface{}和接口类型 不同于 struct,  接口底层有 2 个成员, 一个是 type 一个是 value, 只有当 type 和 value 都为 nil 时, interface{} 才等于 nil
```
var u interface{} = (*interface{})(nil)
if u == nil {
    t.Log("u is nil")
} else {
    t.Log("u is not nil")
}
// u is not nil
```
接口
```
var u Car = (Car)(nil)
if u == nil {
    t.Log("u is nil")
} else {
    t.Log("u is not nil")
}
// u is nil
```
自定义的 struct
```
var u *user = (*user)(nil)
if u == nil {
    t.Log("u is nil")
} else {
    t.Log("u is not nil")
}
// u is nil
```

### map 之坑

#### map 并发读写

`map 并发读写会 panic, 需要加锁或者使用 sync.Map`

map 不能直接更新 value 的某一个字段
```
type User struct{
	name string
}
func TestMap(t *testing.T) {
	m := make(map[string]User)
	m["1"] = User{name:"1"}
	m["1"].name = "2"
	// 编译失败，不能直接修改map的一个字段值
}
```
需要单独拿出来
```
func TestMap(t *testing.T) {
	m := make(map[string]User)
	m["1"] = User{name: "1"}
	u1 := m["1"]
	u1.name = "2"
}
```

### 切片之坑

数组是值类型, 切片是引用类型(指针)
```
func TestArray(t *testing.T) {
	a := [1]int{}
	setArray(a)
	println(a[0])
	// 0
}
func setArray(a [1]int) {
	a[0] = 1
}
func TestSlice(t *testing.T) {
	a := []int{
		1,
	}
	setSlice(a)
	println(a[0])
	// 1
}
func setSlice(a []int) {
	a[0] = 1
}
```

range 遍历
range 会给每一个元素创建一个副本, 会有值拷贝, 如果数组存的是大的结构体可以用 index 遍历或者指针优化

因为 value 是副本, 所以不能修改原有的值

append 会改变地址
slice 类型的本质是一个结构体

type slice struct {
	array unsafe.Pointer
	lenint
	capint
}
函数的值拷贝会导致修改失效
```
func TestAppend1(t *testing.T) {
	var a []int
	add(a)
	println(len(a))
	// 0
}

func add(a []int) {
	a = append(a, 1)
}
```

### 闭包之坑

并发下 go 函数闭包问题
```
for i := 0; i < 3; i++ {
    gofunc() {
        println(i)
    }()
}
time.Sleep(time.Second)
// 2
// 2
// 2
```

说明: 因为闭包导致 i 变量逃逸到堆空间, 所有的 go 共用了 i 变量, 导致并发问题

解决方法1: 局部变量
```
for i := 0; i < 3; i++ {
    ii := i
    gofunc() {
        println(ii)
    }()
}
time.Sleep(time.Second)
// 2
// 0
// 1
```

解决方法2: 参数传递

```
for i := 0; i < 3; i++ {
    gofunc(ii int) {
        println(ii)
    }(i)
}
time.Sleep(time.Second)
// 2
// 0
// 1
```

### buffer 之坑

buffer 对象池
buffer 对象池一定要用完才还回去, 不然buffer在多处复用导致底层的 []byte 内容不一致

参考: golang-buffer-pool

我们的一个 httpClient 返回处理使用了 sync.pool 缓存 buffer, 测试是内存优化了6-8倍

后面测试的时候发现, 获取的内容会偶尔不一致, review 代码发现可能是并发时候 buffer 指针放回去了还在使用, 导致和buffer pool 里面不一致 首先考虑就是将 buffer 的 bytes 读取出来, 然后再 put 回池子里面 然后 bytes 是一个切片, 底层还是和 buffer 共用一个 []byte, buffer 再次修改的时候底层的 []byte 也会被修改, 导致状态不一致 这些理论上是并发问题, 但是我们测试发现, 单线程调用 httpClient 时候, 有时候会有问题, 有时候又没有问题 官方的 http client 做请求的时候会开一个协程, sync pool在同一个协程下面复用对象是一致的, 但是多协程就会新建, 会尝试通过协程的id获取与之对应的对象, 没有才去新建. 串行执行请求也会产生多个协程, 所以偶尔会触发新建 sync 的buffer, 如果新建就不会报错, 如果不新建就会报错.

### select 之坑

for select default 之坑

for 中的 default 在 select 一定会执行, CPU 一直被占用不会让出, 导致 CPU 空转

示例代码
```
func TestForSelect(t *testing.T) {
	for {
		select {
		case <-time.After(time.Second * 1):
			println("hello")
		default:
			if math.Pow10(100) == math.Pow(10, 100) {
				println("equal")
			}
		}
	}
}
```

top CPU 跑满了
```
top - 15:00:50 up 1 day, 15:55,  0 users,  load average: 1.36, 0.85, 0.35
  PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND   
28632 root      20   0 2168296   1.4g   2244 S 252.8  11.7   1:04.15 __debug_bin   
```
