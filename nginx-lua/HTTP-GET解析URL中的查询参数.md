# HTTP-GET解析URL中的查询参数


Query()返回的Values实际上是一个map对象，这个map的key是字符串，value是字符串数组，正好是参数名到参数值数组的map关系，因为同一个参数名可以有多个值，所有参数值是一个字符串数组。例如：
```
map["A"] = ["A1"]
map["B"] = ["B1", "B2"]
```

## nginx-lua运行示例
```
        location /url-test {
            default_type text/html;
            content_by_lua_block {
                ngx.say("uri:", ngx.var.uri)
                ngx.say("request_uri:", ngx.var.request_uri)
                ngx.log(ngx.DEBUG, "uri:", ngx.var.uri, " request_uri:", ngx.var.request_uri)

                local encode = require "cjson" .encode
                local args = ngx.req.get_uri_args()
                ngx.log(ngx.DEBUG, encode(args.method))
                ngx.say("method:", encode(args.method))
            }
        }
```

运行结果
```
# curl 'http://localhost:18080/url-test/ab?method=abc&method=111' 
uri:/url-test/ab
request_uri:/url-test/ab?method=abc&method=111
method:["abc","111"]
```

# 参考链接

- [go语言HTTP GET解析URL中的查询参数](https://www.jianshu.com/p/1075211f0556)
