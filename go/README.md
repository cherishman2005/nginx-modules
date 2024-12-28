# go development

**最优秀的golang开源就是golang源码**

- [golang-channel](/go/golang-channel.md)

- [go-set](/go/go-set.md)

- [Go如何优雅的解决方法重载的问题](/go/Go如何优雅的解决方法重载的问题.md)

- [go-redis-cluster读写分离](./go-redis-cluster读写分离.md)

- [go-mysql](./go-mysql.md)

- [go-sync.Once](./go-sync.Once.md)

- [Golang结构体类型的深浅拷贝](./Golang结构体类型的深浅拷贝.md)

- [Go中map全局变量](./Go中map全局变量.md)

- [golang-linux安装](/go/golang-linux安装.md)

- [go-testing](/go/go-testing/go-testing.md)

- [超时](/go/timeout.md)

- [golang获取字符串中的字符个数](/go/golang获取字符串中的字符个数.md)

- [Golang优雅保持main函数不退出的办法](/go/Golang优雅保持main函数不退出的办法.md)

- [Gin框架介绍及使用](./Gin框架介绍及使用.md)


## Go语言的基本类型

Go语言的基本类型有：
```
bool
string
int、int8、int16、int32、int64
uint、uint8、uint16、uint32、uint64、uintptr
byte // uint8 的别名
rune // int32 的别名 代表一个 Unicode 码
float32、float64
complex64、complex128
```
当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil 等。所有的内存在 Go 中都是经过初始化的。

## golang map数据结构

![image](https://github.com/user-attachments/assets/def454b8-0bbb-4123-8487-fce56f5eee06)


# FAQ

## 引用外部模块函数的首字母要大写

golang public属性：
```
# command-line-arguments
./main.go:31:2: cannot refer to unexported name handler.startHttpService
./main.go:31:2: undefined: handler.startHttpService
```
**修改方法**

startHttpService改为StartHttpService

## 怎样取消代理GOPROXY

取消代理
```
go env -u GOPROXY
```

查看GO的配置
```
go env

//以JSON格式输出
go env -json
```

## go mod tidy 

问题描述
```
go: gopkg.in/yaml.v2@v2.4.0: missing go.sum entry; to add it:
        go mod download gopkg.in/yaml.v2
```

解决方法：
```
go mod tidy 
```

## go clean -modcache

清除依赖，重新构建
```
go clean -modcache
```

## golang开发小结

- [golang开发小结](/go/notebook.md)

## http URL.Path

*http URL.Path不带参数

```
package main
  
import (
        "fmt"
        "net/http"
        //"strings"
)

func helloHandlers(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("URL.Path: %s\n", r.URL.Path)
        //remPartOfURL := r.URL.Path[len("/hello/"):] // get everything after the /hello/ part of the URL
        fmt.Fprintf(w, "Hello %s", r.URL.Path)
}

func main() {
        http.HandleFunc("/", helloHandlers)
        http.ListenAndServe("localhost:9999", nil)
}
```

```
curl 'http://localhost:9999/hello/111?a=3'  -v 
```

运行结果URL.Path=`/hello/111`

### url-scheme

```
// Maybe rawurl is of the form scheme:path.
// (Scheme must be [a-zA-Z][a-zA-Z0-9+-.]*)
// If so, return scheme, path; else return "", rawurl.
func getscheme(rawurl string) (scheme, path string, err error) {
	for i := 0; i < len(rawurl); i++ {
		c := rawurl[i]
		switch {
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z':
		// do nothing
		case '0' <= c && c <= '9' || c == '+' || c == '-' || c == '.':
			if i == 0 {
				return "", rawurl, nil
			}
		case c == ':':
			if i == 0 {
				return "", "", errors.New("missing protocol scheme")
			}
			return rawurl[:i], rawurl[i+1:], nil
		default:
			// we have encountered an invalid character,
			// so there is no valid scheme
			return "", rawurl, nil
		}
	}
	return "", rawurl, nil
}
```

## golang发展前景

1，开发团队，这点go语言创始团队Google背景，c语言之父撑腰，论财力，论背景，论后续维护程度都当之无愧为top5。

2，社区影响力，go的社区影响在分布式领域当之无愧为top3，docker，k8s周边技术栈都是go技术栈！

3，语言易用性，go目前为止有仅次于Python的易用性和超过Python十倍的性能。


# 参考链接

- [https://jogendra.dev/import-cycles-in-golang-and-how-to-deal-with-them](https://jogendra.dev/import-cycles-in-golang-and-how-to-deal-with-them)

- [Golang error 的突围](https://www.cnblogs.com/qcrao-2018/p/11538387.html)

- [理解go的function types](https://www.jianshu.com/p/fc4902159cf5)

- [go benchmark实践与原理](http://cbsheng.github.io/posts/go_benchmark%E5%AE%9E%E8%B7%B5%E4%B8%8E%E5%8E%9F%E7%90%86/)

- [关于go：将一个元素添加到nil slice可以将容量增加两个](https://www.codenong.com/38543825/)

- [https://github.com/struCoder/pidusage](https://github.com/struCoder/pidusage)

- [golang获取用户真实的ip地址](https://blog.thinkeridea.com/201903/go/get_client_ip.html)

- [https://stackoverflow.com/questions/28891531/piping-http-response-to-http-responsewriter](https://stackoverflow.com/questions/28891531/piping-http-response-to-http-responsewriter)

- [golang开源日志组件logrus](https://darjun.github.io/2020/02/07/godailylib/logrus/)

- [golang-ticker](https://wangbjun.site/2020/coding/golang/ticker.html)
- [golang日志库对比](https://blog.csdn.net/General_zy/article/details/124914349)
- [golang日志库对比](https://xiazemin.github.io/MyBlog/golang/2020/05/27/log.html)
- [Go 日志库 Zap 使用](https://chenhe.me/post/go-zap/)
- [golang日志库比较](https://codeantenna.com/a/dBCxE1KrgM)
- [https://stackoverflow.com/questions/48001918/recursively-re-spawn-file-on-fsnotify-remove-rename-golang](https://stackoverflow.com/questions/48001918/recursively-re-spawn-file-on-fsnotify-remove-rename-golang)
- [https://github.com/facebook/watchman/tree/main/watchman](https://github.com/facebook/watchman/tree/main/watchman)
- [go vet](https://zhuanlan.zhihu.com/p/357406395)
- [golang websocket](https://tonybai.com/2019/09/28/how-to-build-websockets-in-go/)
- [https://github.com/phyuany/gin-demo](https://github.com/phyuany/gin-demo)
- [https://github.com/gin-gonic/examples](https://github.com/gin-gonic/examples)
- [Golang框架选型比较: goframe, beego, iris和gin](https://goframe.org/pages/viewpage.action?pageId=3673375)
