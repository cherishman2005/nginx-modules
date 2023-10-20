# golang判断tcp断开

Go 语言中判断 TCP 连接是否断开可以使用如下方法：

使用 net.Error 接口：通过检查 net.Error 接口的 Temporary 方法和 Timeout 方法可以确定是否有错误发生，从而判断连接是否断开。
```
conn, err := net.Dial("tcp", "google.com:80")
if err != nil {
    // handle error
}
defer conn.Close()

_, err = conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
    // handle timeout
} else if err != nil {
    // handle other error
}
```

使用 Read 方法：通过调用 Read 方法并检查返回的错误码可以判断连接是否断开。如果读取到的错误码是 io.EOF，则说明连接已断开。
```
conn, err := net.Dial("tcp", "google.com:80")
if err != nil {
    // handle error
}
defer conn.Close()

_, err = conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
if err != nil {
    // handle error
}

buf := make([]byte, 1024)
n, err := conn.Read(buf)
if err == io.EOF {
    // connection closed
} else if err != nil {
    // handle error
}
```

以上是在 Go 语言中判断 TCP 连接是否断开的两种方法。
