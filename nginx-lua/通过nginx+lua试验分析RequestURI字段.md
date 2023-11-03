# 通过nginx+lua试验分析RequestURI字段

通过nginx+lua试验分析RequestURI字段，包括RequestURI的编码格式等

```
location /url-test {
    default_type text/html;
    content_by_lua_block {
        ngx.say(ngx.var.uri)
    }
}
```

运行结果：
```
# curl 'http://localhost:18080/url-test/abbb?method=1111'
/url-test/abbb
```
