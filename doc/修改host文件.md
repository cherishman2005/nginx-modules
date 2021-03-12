修改host文件

在实际的开发中，有时我们会修改windows的hosts文件，达到指定域名映射到指定ip上的功能。修改方式如下：

1. windows 中hosts文件位置(win10)：
```
C:\Windows\System32\drivers\etc\hosts
```
 

2. 修改方式
```
#将www.aaa.com域名映射到127.0.0.1 IP地址上
127.0.0.1   www.aaa.com
```

3. 应用

在不用重启系统情况下，应用修改后的hosts文件，打开cmd，输入如下命令：

3.1  查看DNS缓存内容
```
ipconfig /displaydns
```

3.2 删除DNS缓存内容，从而达到更新DNS目的
```
ipconfig /flushdns
```