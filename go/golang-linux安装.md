# golang-linux安装

## 下载

```
 wget https://go.dev/dl/go1.14.linux-amd64.tar.gz
```

## 解压

```
tar -C /usr/local -xzf go1.14.linux-amd64.tar.gz
```

## go环境变量设置


sudo vim 打开/etc/profile文件，追加导出命令
```
export PATH=$PATH:/usr/local/go/bin
```

# 参考链接

- [https://go.dev/dl/](https://go.dev/dl/)
