# How to Install OpenSSL from Source in CentOS and Ubuntu

```
$ wget -c https://www.openssl.org/source/openssl-1.0.2p.tar.gz
$ tar -xzvf openssl-1.0.2p.tar.gz
```

编译
```
$ cd openssl-1.0.2p/
$ ./config
$ make
$ make test
$ sudo make install 
```

查看版本
```
/usr/local/ssl/bin/openssl version
```

## 配置

To use the newly installed OpenSSL version on your system, you need to add the directory /usr/local/ssl/bin/ to your PATH, in the file ~/.bashrc (or the equivalent for your shell).

```
$ vim ~/.bashrc
```

Add this line at the bottom of the file.
```
export PATH="/usr/local/ssl/bin:${PATH}"
```
Save and close the file and reload the configuration using the command below.
```
$ source ~/.bashrc
```
