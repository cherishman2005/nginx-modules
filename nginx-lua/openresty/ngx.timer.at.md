# ngx.timer.at

ngx.timer.at 回调的上下文并不允许使用子请求，但很多其他 API是可以使用的，包括非阻塞的 cosocket API:

https://github.com/openresty/lua-nginx-module#ngxsockettcp

你可以直接使用 cosocket API 访问远方服务，或者使用第三方现成的某个 lua-resty-http 库，例如

https://github.com/pintsized/lua-resty-http#readme

比如我的 lua-resty-upstream-healthcheck 库就使用了这种方式：

https://github.com/openresty/lua-resty-upstream-healthcheck#readme

另外，lua-resty-logger-socket 库的实现策略也很值得参考：

https://github.com/cloudflare/lua-resty-logger-socket#readme
