# Golang优雅保持main函数不退出的办法

我们有时会希望我们的程序保持执行，但是有一种情况是：我们的代码全部塞入go routine时，主函数会立刻退出，本文将和大家分享如何让main函数优雅地保持执行。


问题演示：

```
func main() {
    go func() {
        for i := 0; i<10000;i ++ {
            fmt.Println(i)
        }
    }()
}
```
此时我们可以看到，控制台几乎不会输出任何内容。究其原因，是主函数在go routine执行前就已经结束，也就是说go routine不会阻塞主函数。

可能有些读者会想到，我直接加个死循环在下面，让主函数不退出不就行啦？博主表示十分赞同，因为博主就是采用这个方法，导致服务器跑满CPU从而不停的告警。

那么解决办法是：让死循环慢一点执行，即添加以下内容：
```
for {
    time.Sleep(time.Second)
}
```
但是在博主的完美主义光环加持下，还是希望我们的代码能更加优雅，下面将介绍另外三种比较优雅的保持main函数的办法。


解决办法演示

## 操作系统信号阻塞

先上代码：
```
func main() {
    c := make(chan os.Signal)
    signal.Notify(c)
    go func() {
        fmt.Println("Go routine running")
        time.Sleep(3*time.Second)
        fmt.Println("Go routine done")
    }()
    <-c
    fmt.Println("bye")
}
```
官网机翻：signal.Notify()方法使信号将传入c。如果没有提供信号，所有传入的信号将被中继到c。

这里我们创建了一个os.Signal类型的管道。当管道为空的时候，读管道操作“<-”会阻塞住，直到我们向进程发送一个信号（例如 Ctrl+C），才会继续执行该操作后面的代码。

## 上下文操作阻塞

再上代码：
```
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    go func() {
        fmt.Println("Go routine running")
        time.Sleep(3 * time.Second)
        fmt.Println("Go routine done")
        cancel()
    }()
    <-ctx.Done()
    fmt.Println("bye")
}
```
官网机翻：CancelFunc() 通知操作放弃其（当前的）工作。CancelFunc() 不会等待工作停止。

这也是一个十分优雅的办法，我们创建一个可以终止的上下文——context.WithCancel()，并在go routine执行完毕时调用其返回的CancelFunc() 方法，即表示该上下文已经结束了。而在这之前，我们会使用<-ctx.Done()来一直等待上下文的结束，也就是说main函数被成功阻塞，并等待go routine执行完毕并执行了cancel()方法后优雅退出。

## WaitGroup阻塞

然后上代码：
```
func main() {
    wg := &sync.WaitGroup{}
    wg.Add(2)
    go func() {
        time.Sleep(3*time.Second)
        fmt.Println("3 second passed")
        wg.Done()
    }()
    go func() {
        time.Sleep(5*time.Second)
        fmt.Println("5 second passed")
        wg.Done()
    }()
    wg.Wait()
    fmt.Println("bye")
}
```
官网机翻：WaitGroup 等待一组 go routine 完成。主 go routine 调用 Add() 来设置要等待的 go routine 的数量。

我们首先创建一个WaitGroup{}对象，然后调用Add()方法，在里面传入我们接下来会创建的go routine的数量，每当我们执行完一个go routine时，调用一次Done()方法，使得正执行的go routine的数量减一，当减到0时，Wait()方法将不再等待（阻塞），使main函数继续向下执行。

# 小结

以上就是我们告别for {}或者select {}，并优雅地阻塞主函数的三种办法，也是博主作为新手时对Go语言特性的入门级体验。
