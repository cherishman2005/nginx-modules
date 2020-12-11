# nginx-lua


# FAQ

## ngx.timer.at

ngx.timer.at()。 ngx.timer.at 会创建一个 Nginx timer。在事件循环中，Nginx 会找出到期的 timer，并在一个独立的协程中执行对应的 Lua 回调函数。 有了这种机制，ngx_lua 的功能得到了非常大的扩展，我们有机会做一些更有想象力的功能出来。比如 批量提交和 cron 任务。随便一提，官方的 resty-cli 工具，也是基于 ngx.timer.at 来运行指定的代码块。

ngx.timer.at 的 delay 参数，指定的是以秒为单位的延迟触发时间。跟 OpenResty 的其他函数一样，指定的时间最多精确到毫秒。如果你想要的是一个当前阶段结束后立刻执行的回调，可以直接设置 delay 为 0。 handler 回调第一个参数 premature，则是用于标识触发该回调的原因是否由于 timer 的到期。Nginx worker 的退出，也会触发当前所有有效的 timer。这时候 premature 会被设置为 true。回调函数需要正确处理这一参数（通常直接返回即可）。

需要特别注意的是：有一些 ngx_lua 的 API 不能在这里调用，比如子请求、ngx.req.*和向下游输出的 API(ngx.print、ngx.flush 之类)，原因是这些调用需要依赖具体的请求。但是 ngx.timer.at 自身的运行，与当前的请求并没有关系的。

## openresty(nginx-lua)脚本调试

```
/usr/local/openresty/bin/resty  h.lua
```


# 参考文献

- [https://moonbingbing.gitbooks.io/openresty-best-practices/content/ngx_lua/timer.html](https://moonbingbing.gitbooks.io/openresty-best-practices/content/ngx_lua/timer.html)

- [https://xiangxianzui.github.io/2017/11/nginx-lua%E5%8A%A8%E6%80%81%E6%94%B9%E5%8F%98upstream/](https://xiangxianzui.github.io/2017/11/nginx-lua%E5%8A%A8%E6%80%81%E6%94%B9%E5%8F%98upstream/)

- [Dynamic NGINX Upstreams from Consul via lua-nginx-module](https://medium.com/@sigil66/dynamic-nginx-upstreams-from-consul-via-lua-nginx-module-2bebc935989b)

- [Nginx 负载均衡策略之“快者优先”的 Lua 实现](https://toutiao.io/posts/ocepp3/preview)

- [https://github.com/openresty/lua-resty-balancer](https://github.com/openresty/lua-resty-balancer)

- [https://github.com/jaderhs/lua-consistent-hash](https://github.com/jaderhs/lua-consistent-hash)
