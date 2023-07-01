# Go编程艺术-异步读写网络连接

在Go中，读写channel与读写Conn是的不一样的。一个协程可以通过select监听并读写多个channel, 但却不能类似这样监听并读写多个Conn（除非自行封装epoll）。如果想在一个协程中监听并读写多个channel和Conn, 那更不好办了。一般的方式是使用专门的协程来处理Conn的读写，使其转换为channel，以供主协程多路监听读写。

由于Conn读写都会阻塞协程，一般情况下，需要两个协程。一个用来读，一个用来写。此时一定要处理好读写的错误，要及时地关闭掉Conn,并退出两个协程，释放掉资源。

这里示例一下处理方法：

![image](https://github.com/cherishman2005/nginx-modules/assets/17688273/5daf8a81-9c66-4dd1-ad3b-cb03331f917a)

转化读写Conn为chan的收发

```
func process(conn net.Conn, reqch chan<- Request, rspch <-chan Response) {
    var wg sync.WaitGroup
    ctx, cancel := context.WithCancel(context.Background())
    
    // 写
    wg.Add(1)
    go func() {
        defer wg.Done()
    
        for {
            select {
            case req := <-reqch:
                _, err := conn.Write(req.Marshal())
                if err != nil {
                    conn.Close()
                    return
                }
              
              	// 处理数据
              	// ...
            }
            
            case <-ctx.Done():
          			// 读协程已关闭conn, 写协程需要立即退出。
                return
        }
    }()
    
    // 读
    go func() {
        defer wg.Done()
        
        rbuf := make([]byte, 1024)
        
        for {
            rlen, err := conn.Read(rbuf)
            if err != nil{
                if err != net.ErrClosed() {
                    conn.Close()
                  
                  	// 及时通知写协程
                    cancel()
                }
                return
            }
            
            rsp := Unmarshal(rbuf[:rlen])
            
            // 处理数据
            // ...
            
            rspch <- rsp
        }
    }()
    
    wg.Wait()
}
```

需要注意的是context，这个能保证两个协程的协调退出。

写协程写的时候出错，关闭掉Conn，退出；读协程会立即发现Conn已被关掉，退出。
读协程读的时候出错，关闭掉Conn，cancel(), 退出。写协程立即发现Conn已被关闭，立即退出。
关键的点在于，如果不用context, 写协程只有在下次Write的时候才能发现Conn已被关闭，而这个“下次”，不知道是什么时候。

这样就把对Conn的读写转化为了两个chan，主协程只需要两个chan通信。

```
func processProxy() {
  var conn net.Conn
  var reqch chan Request
  var rspch chan Response
	// ...
  go process(conn, reqch, rspch)
  
	
  select {
    case reqch <- req:
    	// send
    
    case rsp := <-rspch:
    	// recv
    
    default:
    	// ok
  }
  
}
```

当然，示例程序也不完美。比如，往reqch发送的时候，可以用一个函数封装一下，及时发现chan已满，发送报错。
