# go development

- [go-redis-cluster读写分离](./go-redis-cluster读写分离.md)

- [go-mysql](./go-mysql.md)

- [go-sync.Once](./go-sync.Once.md)

- [Golang结构体类型的深浅拷贝](./Golang结构体类型的深浅拷贝.md)

- [Go中map全局变量](./Go中map全局变量.md)

- [golang-linux安装](/go/golang-linux安装.md)

- [go-testing](/go/go-testing/go-testing.md)

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

# 参考链接

- [Golang error 的突围](https://www.cnblogs.com/qcrao-2018/p/11538387.html)

- [理解go的function types](https://www.jianshu.com/p/fc4902159cf5)

- [go benchmark实践与原理](http://cbsheng.github.io/posts/go_benchmark%E5%AE%9E%E8%B7%B5%E4%B8%8E%E5%8E%9F%E7%90%86/)
