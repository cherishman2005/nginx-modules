# 开发nginx c module的一些调试技巧

## 观察workder进程是否有core?

* 调试nginx c module的一个技巧： 看进程core?
```
root      77161      1  0 17:48 ?        00:00:00 nginx: master process /usr/local/openresty/nginx/sbin/nginx
nobody    77199  77161  0 17:49 ?        00:00:00 nginx: worker process
```

* 通过error日志，发现进程ID不断跳跃。

  nginx正常过程中，一般master进程ID与worker进程ID是连续 或者相差不大。

## 调整为worker单进程模式

nginx.conf设置关闭daemon模式：
```
daemon off;
```
然后gdb调试；

## http content-length和chunk

`抓包分析`：

* 分析body长度；或者chunk格式是否正确？
  
  有时会出现抓包时最后一个字节时完整的，但是后端日志打印会少一个字节，特别是json解析失败。
  
  —— 所以不仅仅要抓包，还有后端日志确认。

## 添加日志打印

  因为nginx很多异步操作，gdb调试不一定方便；nginx日志较少。通过`添加一些日志打印`，来定位问题。

  【注】nginx编译时要打开debug选项，并在nginx.conf打开debug开关。
