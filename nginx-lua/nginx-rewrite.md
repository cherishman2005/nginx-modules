# nginx-rewrite

## ngx.exec

语法：ngx.exec(uri, args?)
主要实现的是内部的重定向，等价于下面的rewrite指令:
```
rewrite regrex replacement last;
```
