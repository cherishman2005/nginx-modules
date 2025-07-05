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
