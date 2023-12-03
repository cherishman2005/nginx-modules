# go语言http响应头设置

Go 语言的 net/http 包为我们提供了一个 HTTP 客户端和服务器的实现，通过它我们可以快速的搭建一个 HTTP 服务器，本文记录一下在编写 HTTP 服务器时关于设置 HTTP 响应头遇到的一个小问题。

问题描述
问题的表象是通过 w.Header().Set("Content-Type", "application/json") （w 为 http.ResponseWriter 对象）无法设置响应头的 Content-Type 为 application/json。下面是是一个简单的示例：
```
package main

import (
        "encoding/json"
        "net/http"
)

type Status struct {
        Code    int
        Message string
}

func hello(w http.ResponseWriter, r *http.Request) {
        res := Status{Code: 200, Message: "hello world!"}

        w.WriteHeader(200)
        w.Header().Set("Content-Type", "application/json")

        json.NewEncoder(w).Encode(&res)
}

func main() {
        http.HandleFunc("/hello", hello)

        http.ListenAndServe(":8080", nil)
}
```
编译并运行上面的程序，随后我们去访问可以看到如下所示的响应。
```
$ curl -i http://localhost:8080/hello
HTTP/1.1 200 OK
Date: Mon, 14 Sep 2020 09:40:11 GMT
Content-Length: 38
Content-Type: text/plain; charset=utf-8

{"Code":200,"Message":"hello world!"}
```
从响应的结果可以看到，Content-Type 并没有被设置为我想要的 application/json。

解决方案
通过测试我发现，经过如下修改，就可以正确的设置 HTTP 响应头。
```
diff --git a/main.go b/main.go
index 8bec8e3..459fb77 100644
--- a/main.go
+++ b/main.go
@@ -13,8 +13,8 @@ type Status struct {
 func hello(w http.ResponseWriter, r *http.Request) {
        res := Status{Code: 200, Message: "hello world!"}

-       w.WriteHeader(200)
        w.Header().Set("Content-Type", "application/json")
+       w.WriteHeader(200)

        json.NewEncoder(w).Encode(&res)
 }
``` 
以下是测试输出。
```
$ curl -i http://localhost:8080/hello
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 14 Sep 2020 09:48:52 GMT
Content-Length: 38

{"Code":200,"Message":"hello world!"}
```
这是为什么呢？通过查看 Go 语言的源码，发现了如下内容：
```
type ResponseWriter interface {
// Header returns the header map that will be sent by
// WriteHeader. The Header map also is the mechanism with which
// Handlers can set HTTP trailers.
//
// Changing the header map after a call to WriteHeader (or
// Write) has no effect unless the modified headers are
// trailers.
//
// There are two ways to set Trailers. The preferred way is to
// predeclare in the headers which trailers you will later
// send by setting the “Trailer” header to the names of the
// trailer keys which will come later. In this case, those
// keys of the Header map are treated as if they were
// trailers. See the example. The second way, for trailer
// keys not known to the Handler until after the first Write,
// is to prefix the Header map keys with the TrailerPrefix
// constant value. See TrailerPrefix.
//
// To suppress automatic response headers (such as “Date”), set
// their value to nil.
Header() Header
…
}
```
从注释我们可以看到，在调用了 WriteHeader() 或者是 Write() 函数之后去改变响应头（trailers 响应头除外）是没有效果的。
其实，在 Go 的文档里面也出现了上述注释，只是在读文档的时候不够仔细，导致忽略了这一点。

关于 HTTP Trailers 的相关说明可以参考这里。

我们从源码角度来做简要分析。首先，我们看看 ResponseWriter 是如何获取 Header 结构的，如下所示。
```
func (w *response) Header() Header {
        if w.cw.header == nil && w.wroteHeader && !w.cw.wroteHeader {
                // Accessing the header between logically writing it
                // and physically writing it means we need to allocate
                // a clone to snapshot the logically written state.
                w.cw.header = w.handlerHeader.Clone()
        }
        w.calledHeader = true
        return w.handlerHeader
}
```
从上面的代码可以看出，我们在调用 w.Header() 函数获取的其实是 w.handlerHeader，因此，我们添加的 HTTP 响应头也是存放在该结构中的。接着我们在来看看 w.WriteHeader() 函数是如何实现的。
```
func (w *response) WriteHeader(code int) {
        if w.conn.hijacked() {
                caller := relevantCaller()
                w.conn.server.logf("http: response.WriteHeader on hijacked connection from %s (%s:%d)", caller.Function, path.Base(caller.File), caller.Line)
                return
        }
        if w.wroteHeader {
                caller := relevantCaller()
                w.conn.server.logf("http: superfluous response.WriteHeader call from %s (%s:%d)", caller.Function, path.Base(caller.File), caller.Line)
                return
        }
        checkWriteHeaderCode(code)
        w.wroteHeader = true
        w.status = code

        if w.calledHeader && w.cw.header == nil {
                w.cw.header = w.handlerHeader.Clone()
        }

        if cl := w.handlerHeader.get("Content-Length"); cl != "" {
                v, err := strconv.ParseInt(cl, 10, 64)
                if err == nil && v >= 0 {
                        w.contentLength = v
                } else {
                        w.conn.server.logf("http: invalid Content-Length of %q", cl)
                        w.handlerHeader.Del("Content-Length")
                }
        }
}
```
我们注意到，w.WriteHeader() 将 w.handlerHeader 中的 HTTP 响应头克隆到了 w.cw.header 中（其中 cw 为 *chunkWriter 类型）。最终，net/http 是通过 func (cw *chunkWriter) writeHeader(p []byte) 函数来完成响应头的处理的，而这其中关于 HTTP 响应头，除了 Trailer 类型的响应头外，都是从 cw.header 中获取的，因此，在调用 w.WriteHeader() 之后，我们设置的非 Trailer 响应头是无效的。

# 参考

[1] https://github.com/golang/go/blob/master/src/net/http/server.go
[2] https://golang.org/pkg/net/http/
[3] https://www.geeksforgeeks.org/http-headers-trailer/
