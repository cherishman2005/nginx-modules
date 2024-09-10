# thrift安装

1. 下载

https://archive.apache.org/dist/thrift/0.13.0/

```
wget https://archive.apache.org/dist/thrift/0.13.0/thrift-0.13.0.tar.gz
```

2. 解压与编译

```
tar -xzvf thrift-0.13.0.tar.gz

cd thrift-0.13.0/
```

```
./configure   #这一步配置可以选择支持的语言，如python,c++，go等。可以./configure --help来查看可配置项。如./configure --without-go
make && make install 
```
