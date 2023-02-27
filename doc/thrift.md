# thrift

thrift 最初是 facebook 开发使用的 rpc 通信框架，后来贡献给了 apache 基金会，出来得比较早，几乎支持所有的后端语言，使用非常广泛，是不可不知的一个网络框架。



```
git clone https://github.com/apache/thrift
cd thrift
./bootstrap.sh
./configure --without-qt4 --wihout-qt5
make
sudo make install
```

# FAQ

问题
```
make distclean... ok
Couldn't find libtoolize!
```

解决方法
```
sudo apt install libtool automake build-essential flex bison libboost-dev libssl-dev -y
```

# 参考链接

- [https://studygolang.com/articles/13988](https://studygolang.com/articles/13988)
