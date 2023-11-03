# proxy_pass中使用$request_uri代替$uri

传递原生请求信息:

将未更改的URI传递到上游服务的最佳规则是使用proxy_pass http://<backend>，不带任何参数（uri等部分）。

如果你需要对proxy_pass的上游服务uri做处理，请使用$request_uri参数而代替$uri参数

例如，在proxy_pass指令中不小心使用$uri会导致http头注入漏洞， 因为URL编码字符会被解码(这有时很重要，并不等同于$request_uri)。

更重要的是，$uri的值可能会在请求处理过程中改变，例如在进行内部重定向时，或者在使用索引文件时。

* 不建议的配置方式:
```nginx configuration
location /foo {
    proxy_pass http://django_app_server$uri;
}
```

* 建议的配置方式:

```nginx configuration
location /foo {
    proxy_pass http://django_app_server$request_uri;
}
```

最佳配置（不做任何处理）:
```nginx configuration
location /foo {
    proxy_pass http://django_app_server;
}
```
