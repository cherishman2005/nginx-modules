# ab压测

## ab安装

```
apt install apache2-utils
```

## 压测示例

```
 ab -n 20 -c 10 'http://localhost:28080/test-sinfo?action=getSidAuthState&tid=36569435&sid=2790311106'
```

其中－n表示请求数，－c表示并发数


# 参考链接

- [https://www.jianshu.com/p/43d04d8baaf7](https://www.jianshu.com/p/43d04d8baaf7)
