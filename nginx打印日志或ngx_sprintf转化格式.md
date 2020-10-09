---
layout: post
title: nginx打印日志或ngx_sprintf转化格式
categories: nginx
description: nginx打印日志或ngx_sprintf转化格式
keywords: nginx
# topmost: true
---

# nginx打印日志或ngx_sprintf转化格式

格式化参数：如果误用的话，`轻则输出不正确，重则nginx可能core`。

```
/*
 * supported formats:
 *    %[0][width][x][X]O        off_t
 *    %[0][width]T              time_t
 *    %[0][width][u][x|X]z      ssize_t/size_t
 *    %[0][width][u][x|X]d      int/u_int
 *    %[0][width][u][x|X]l      long
 *    %[0][width|m][u][x|X]i    ngx_int_t/ngx_uint_t
 *    %[0][width][u][x|X]D      int32_t/uint32_t
 *    %[0][width][u][x|X]L      int64_t/uint64_t
 *    %[0][width|m][u][x|X]A    ngx_atomic_int_t/ngx_atomic_uint_t
 *    %[0][width][.width]f      double, max valid number fits to %18.15f
 *    %P                        ngx_pid_t
 *    %M                        ngx_msec_t
 *    %r                        rlim_t
 *    %p                        void *
 *    %V                        ngx_str_t *
 *    %v                        ngx_variable_value_t *
 *    %s                        null-terminated string
 *    %*s                       length and string
 *    %Z                        '\0'
 *    %N                        '\n'
 *    %c                        char
 *    %%                        %
 *
 *  reserved:
 *    %t                        ptrdiff_t
 *    %S                        null-terminated wchar string
 *    %C                        wchar
 */
```

![ngx_sprintf](/img/ngx_sprintf.png)
![ngx_sprintf](/img/ngx_sprintf1.png)

# 参考链接

- 陶辉 《深入理解Nginx模块开发与架构解析》（第2版P148）

- [https://tengine.taobao.org/book/chapter_02.html](https://tengine.taobao.org/book/chapter_02.html)

- [https://www.nginx.com/resources/wiki/extending/api/utility/](https://www.nginx.com/resources/wiki/extending/api/utility/)
