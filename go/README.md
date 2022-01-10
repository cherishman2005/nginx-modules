# go development


- [go-redis-cluster读写分离](./go-redis-cluster读写分离.md)


# FAQ

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
