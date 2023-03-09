# go development

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

# 参考链接

- [Golang error 的突围](https://www.cnblogs.com/qcrao-2018/p/11538387.html)

- [理解go的function types](https://www.jianshu.com/p/fc4902159cf5)

- [go benchmark实践与原理](http://cbsheng.github.io/posts/go_benchmark%E5%AE%9E%E8%B7%B5%E4%B8%8E%E5%8E%9F%E7%90%86/)

- [关于go：将一个元素添加到nil slice可以将容量增加两个](https://www.codenong.com/38543825/)

- [https://github.com/struCoder/pidusage](https://github.com/struCoder/pidusage)

- [golang获取用户真实的ip地址](https://blog.thinkeridea.com/201903/go/get_client_ip.html)

- [https://stackoverflow.com/questions/28891531/piping-http-response-to-http-responsewriter](https://stackoverflow.com/questions/28891531/piping-http-response-to-http-responsewriter)
