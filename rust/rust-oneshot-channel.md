# rust-oneshot-channel

以下是一个使用 Rust 的 oneshot 通道实现异步消息传递的完整示例。oneshot 通道是一种特殊的消息通道，用于一次性发送和接收单个值，非常适合实现请求 - 响应模式或跨线程通知。

```rust
use futures::executor::block_on;
use tokio::sync::oneshot;

#[tokio::main]
async fn main() {
    // 创建一个 oneshot 通道
    let (sender, receiver) = oneshot::channel();

    // 启动一个异步任务发送消息
    tokio::spawn(async move {
        // 模拟一些耗时操作
        tokio::time::sleep(tokio::time::Duration::from_secs(1)).await;
        
        // 发送消息到通道
        let result = sender.send("Hello from the other task!");
        
        // 处理发送结果
        if let Err(msg) = result {
            println!("发送失败: 接收端已关闭，消息为 '{}'", msg);
        } else {
            println!("消息发送成功");
        }
    });

    // 在主任务中接收消息
    println!("等待消息...");
    match receiver.await {
        Ok(message) => println!("收到消息: {}", message),
        Err(e) => println!("接收失败: {}", e),
    }
    
    println!("程序结束");
}
```

## 代码说明：

1. 通道创建：
```rust
let (sender, receiver) = oneshot::channel();
```
创建一个 oneshot 通道，返回发送端 sender 和接收端 receiver。

2. 发送消息：
```rust
sender.send("Hello from the other task!");
```
使用 send 方法发送一个字符串。如果接收端已关闭，send 会返回错误。

3. 接收消息：
```rust
match receiver.await { ... }
```
receiver.await 是异步操作，会阻塞当前任务直到收到消息或通道关闭。

4. 错误处理：
* 发送失败：接收端可能已被丢弃。
* 接收失败：发送端可能已被丢弃，或在发送前就被销毁。

## 运行方式：

创建新的 Rust 项目：
```bash
cargo new oneshot_demo --bin && cd oneshot_demo
```

添加依赖到 Cargo.toml：
```toml
[dependencies]
tokio = { version = "1", features = ["full"] }
futures = "0.3"

运行代码：
```bash
cargo run
```

## 输出示例：
```plaintext
等待消息...
消息发送成功
收到消息: Hello from the other task!
程序结束
```

## 注意事项：
* 一次性使用：oneshot 通道只能发送一次消息，发送后通道关闭。
* 所有权转移：send 方法会转移消息的所有权到接收端。
* 异步特性：需要使用 async/await 或 block_on 来处理异步接收。

这个示例展示了 oneshot 通道的基本用法，常用于跨任务或跨线程的一次性通信场景。
