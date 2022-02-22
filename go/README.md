# go development


- [go-redis-cluster读写分离](./go-redis-cluster读写分离.md)

- [go-mysql](./go-mysql.md)

- [go-sync.Once](./go-sync.Once.md)


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
