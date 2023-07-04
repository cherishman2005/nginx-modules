# go-channel与锁

不同 goroutine 之间如何通讯？有两种方案：

1. 全局变量的互斥锁

2. 使用管道 channel 来解决

因为没有对全局变量 m 加锁，因此会出现资源争夺问题，代码会出现错误，提示 concurrent map　writes

```golang
var (
    myMap = make(map[int]int, 10)
    lock sync.Mutex
)

func test(n int) {
    res := 1
    for i := 1; i <= n; i++ {
        res *= i
    }
    //这里我们将 res 放入到 myMap
    lock.Lock()
    myMap[n] = res //concurrent map writes?
    lock.Unlock()
}

func main()  {
    // 我们这里开启多个协程完成这个任务[200 个]
    for i := 1; i <= 200; i++ {
        go test(i)
    }
    //休眠 10 秒钟【第二个问题 】
    time.Sleep(time.Second * 10)
    lock.Lock()
    for i, v := range myMap {
        fmt.Printf("map[%d]=%d\n", i, v)
    }
    lock.Unlock()
}
```

1) 前面使用全局变量加锁同步来解决 goroutine 的通讯，但不完美

2) 主线程在等待所有 goroutine 全部完成的时间很难确定，我们这里设置 10 秒，仅仅是估算。如果主线程休眠时间长了，会加长等待时间，如果等待时间短了，可能还有 goroutine 处于工作
状态，这时也会随主线程的退出而销毁

3) 通过全局变量加锁同步来实现通讯，也并不利于多个协程对全局变量的读写操作。

要解决上面的不足，引入一个新的通讯机制 - channel

1) channel 是一个队列，那么就有先进先出的特点

2) 线程安全，多 goroutine 访问时，不需要加锁，说明 channel 本身就是线程安全的

3) channel 有类型的，一个 int 的 channel 只能存放 string 数据类型

var 变量名 chan 数据类型
var intChan chan
var mapChan chan map [int] string (mapChan 用于存 map [int] string 类型)

channel 是引用类型

channel 必须初始化才能写入数据，即 make 后才能使用

管道是有类型的，intChan 只能写入 整数 int

```golang
func main() {
    //演示一下管道的使用
    //1. 创建一个可以存放 3 个 int 类型的管道
    var intChan chan int
    intChan = make(chan int, 3)
    //2. 看看 intChan 是什么
    fmt.Printf("intChan 的值=%v intChan 本身的地址=%p\n", intChan, &intChan)
    //3. 向管道写入数据
    intChan<- 10
    num := 211
    intChan<- num
    intChan<- 50
    // intChan<- 98//注意点, 当我们给管写入数据时,不能超过其容量
    //4. 看看管道的长度和 cap(容量)
    fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3
    //5. 从管道中读取数据
    var num2 int
    num2 = <-intChan
    fmt.Println("num2=", num2)
    fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan))
    // 2, 3
    //6. 在没有使用协程的情况下,如果我们的管道数据已经全部取出,再取就会报告 deadlock
    num3 := <-intChan
    num4 := <-intChan
    num5 := <-intChan
    fmt.Println("num3=", num3, "num4=", num4, "num5=", num5)
}
```

运行结果：
```
intChan 的值=0xc0000ba000 intChan 本身的地址=0xc0000b4018
channel len= 3 cap=3 
num2= 10
channel len= 2 cap=3 
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan receive]:
main.main()
        /home/malina/gotest/src/test3/main.go:29 +0x395
Process finished with the exit code 2
```

1) channel 中只能存放指定的数据类型

2) channle 的数据放满后，就不能再放入了

3) 如果从 channel 取出数据后，可以继续放入

4) 在没有使用协程的情况下，如果 channel 数据取完了，再取，就会报 dead lock
