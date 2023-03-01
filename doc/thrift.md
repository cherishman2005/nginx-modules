# thrift

thrift 最初是 facebook 开发使用的 rpc 通信框架，后来贡献给了 apache 基金会，出来得比较早，几乎支持所有的后端语言，使用非常广泛，是不可不知的一个网络框架。



```
git clone https://github.com/apache/thrift
cd thrift
./bootstrap.sh
./configure --without-qt4 --without-qt5
make
sudo make install
```

## 编译方式2

thrift-v0.11.0
``` 
./bootstrap.sh
./configure --with-boost=/usr/local
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

## 安装boost

Debian/Ubuntu install
The following command will install tools and libraries required to build and install the Apache Thrift compiler and C++ libraries on a Debian/Ubuntu Linux based system.
```
sudo apt-get install automake bison flex g++ git libboost-all-dev libevent-dev libssl-dev libtool make pkg-config
```
Debian 7/Ubuntu 12 users need to manually install a more recent version of automake and (for C++ library and test support) boost:
```
wget http://ftp.debian.org/debian/pool/main/a/automake-1.15/automake_1.15-3_all.deb
sudo dpkg -i automake_1.15-3_all.deb
```

```
wget http://sourceforge.net/projects/boost/files/boost/1.60.0/boost_1_60_0.tar.gz                                                                      tar xvf boost_1_60_0.tar.gz
cd boost_1_60_0
./bootstrap.sh
sudo ./b2 install
```

# 参考链接

- [https://studygolang.com/articles/13988](https://studygolang.com/articles/13988)

- [https://www.cnblogs.com/52fhy/p/11146047.html](https://www.cnblogs.com/52fhy/p/11146047.html)

- [https://www.cnblogs.com/yhp-smarthome/p/8982758.html](https://www.cnblogs.com/yhp-smarthome/p/8982758.html)
