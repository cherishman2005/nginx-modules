# golang端口复用


在 Golang 中实现端口复用，可以使用 net 包中的 Listen 函数的 Setsockopt 方法，设置 SO_REUSEADDR 选项。

SO_REUSEADDR 选项表示在一个端口上绑定的套接字关闭后，该端口仍然可以被立即重新使用。这个选项可以解决 Address already in use 的错误。

下面是一个示例代码：
```
package main

import (
    "net"
)

func main() {
    addr := "localhost:8080"
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        panic(err)
    }

    // 设置 SO_REUSEADDR 选项
    if tcp, ok := ln.(*net.TCPListener); ok {
        err := tcp.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
        if err != nil {
            panic(err)
        }
    }

    // 处理连接
    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        }
        // 处理 conn
    }
}
```
在上面的代码中，我们使用 net.Listen 函数创建了一个 TCPListener，并将 SO_REUSEADDR 选项设置为 1。如果设置成功，我们可以继续接受连接。

需要注意的是，SO_REUSEADDR 选项应该在调用 Listen 函数之后，Accept 函数之前设置，否则设置可能会失败。
