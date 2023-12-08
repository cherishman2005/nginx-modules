# Wireshark：No interfaces found解决方法（Windows 10） 


启动Wireshark时有时会报“No interfaces found”，找不到网卡进行截包。造成这种情况的原因可能有两个，一是npf服务没启动，二是当前用启对网卡没有拦截权限。


## 一、npf服务未启动处理

### 1.1 找到cmd右键以管理员身份运行


### 1.2 用以下命令启动npf

```
net query npf       #检查npf运行情况

net start npf       #启动npf
```

 

## 二、当前用户对网卡没有拦截权限处理

由于Windows 10默认使用的不是Administrator账号，所以这种情况比较常见。

右键Wireshark快键方式，选择以管理员身份运行Wireshark


# 参考链接

- [Wireshark：No interfaces found解决方法（Windows 10）](https://www.cnblogs.com/lsdb/p/7144605.html)
