# openssl

## 编译与安装
```
./config --prefix=/usr/local/openssl
 
 make -j32 && make install
```

## luarocks配置

```
sudo luarocks config variables.OPENSSL_INCDIR /usr/local/openssl/include/
```
