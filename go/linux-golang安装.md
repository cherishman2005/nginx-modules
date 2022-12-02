# linux-golang安装

## 下载golang安装包

https://go.dev/dl/

```
wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz --no-check-certificate
```

## 安装

执行tar解压到/usr/local目录下（官方推荐），得到go文件夹等
```
tar -C /usr/local -zxvf  go1.18.3.linux-amd64.tar.gz
```


## 添加/usr/local/go/bin目录到PATH变量中。添加到/etc/profile 或$HOME/.profile都可以


vim /etc/profile
```
# 在最后一行添加
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

```
source /etc/profile
```
