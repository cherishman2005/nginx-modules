# notebook

记录golang开发中遇到的一些问题和解决方法。

## replace

go mod
```
module agent

go 1.18

replace github.com/apache/thrift => ./pkg/thrift-0.13.0

require (
    github.com/apache/thrift v0.13.0
)
```

make编译出错：
```
go: github.com/apache/thrift@v0.18.1 (replaced by ./pkg/thrift-0.13.0): reading pkg/thrift-0.13.0/go.mod: open /home/nginx/proxy/thrift_server/proxy/pkg/thrift-0.13.0/go.mod: no such file or directory
```

**解决方法**
需要在pkg/thrift-0.13.0目录下创建go.mod
```
go mod init thrift
```

然后继续在工程目录执行：
```
go mod tidy
make
```
编译成功。

