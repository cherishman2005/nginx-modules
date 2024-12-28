# Channel

## 设计原理

### 多线程使用共享内存传递数据

最直观的线程中传递数据是使用共享内存，为了防止读写冲突，需要加锁。

![image](https://github.com/user-attachments/assets/f093a49a-90af-4873-8b69-0d61fdd86020)

### Goroutine 使用 Channel 传递数据

goroutine协程采用channel传递数据。


![image](https://github.com/user-attachments/assets/ebb7eecb-1190-47d9-bd38-1d8944d4d3b2)

上图中的两个 Goroutine，一个会向 Channel 中发送数据，另一个会从 Channel 中接收数据，它们两者能够独立运行并不存在直接关联，但是能通过 Channel 间接完成通信。

#### 先入先出

目前的 Channel 收发操作均遵循了先进先出的设计，具体规则如下：
* 先从 Channel 读取数据的 Goroutine 会先接收到数据；
* 先向 Channel 发送数据的 Goroutine 会得到先发送数据的权利；


# channel优缺点

1. 优点：

* 保证数据有序；

2. 缺点：

即使是带有缓存的channel，也会出现阻塞，channel大小初始化后没法改变。

虽然出现阻塞后，write端 采用timeout机制，变成不阻塞。但是没有写成功的数据怎么处理？ 出现了read端 数据 不符合预期（预期是希望有序/不丢失）。

`因为有这样的缺点，在有些场景是不满足要求的`。例如做proxy转发功能 或 k8s调度功能时，必须要有序，数据不丢失。

## 怎么解决这个问题

* 根据自己的需求，设计自己的数据结构。


# 小结

Channel 是 Go 语言能够提供强大并发能力的原因之一。一般情况下优先使用channel，使用起来简单，方便。

# 参考链接

- [https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-channel/](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-channel/)

