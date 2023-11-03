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


# FAQ

## nginx $uri 和 $request_uri 的区别

$uri 指的是请求的文件和路径，不包含”?”或者”#”之类的东西

$request_uri 则指的是请求的整个字符串，包含了后面请求的东西

例如：
```
$uri： www.baidu.com/document

$request_uri： www.baidu.com/document?x=1
```
