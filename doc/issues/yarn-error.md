# yarn报错

```
yarn 错误There appears to be trouble with your network connection. Retrying…
```

解决途径：
```
#查看代理
yarn config list
#删除代理
yarn config delete proxy
#更换淘宝镜像
yarn config set registry https://registry.npm.taobao.org
```
