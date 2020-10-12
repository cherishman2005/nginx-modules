

# nginx-njs

2015年9月，nginx宣布支持类JavaScript语言。

1. nginScript是以nginx插件的方式存在。 插件名叫：njs 。和其他nginx插件一样，需要重新编译nginx来完成安装。
1. nginScript是JavaScript/ECMAscript的子集 。它实现了大部分的JavaScript语言的能力，没有完全遵从ECMAScript标准，同时抛弃了JavaScript比较难懂的部分。
1. nginScript不是通过V8引擎实现的 。而是通过一个更小、能耗更低、更符合nginx应用场景的小虚拟机（VM）来实现。可以理解为nginx为其实现了一套自己的词法解析。
1. nginScript是跑在nginx的配置文件里。 比如：nginx.conf文件里。所以nginScript可以完成传统配置文件所能处理的所有事情，同时可以让配置管理动态化。这也是nginScript出现的最重要的原因。

## nginx-njs安装配置

1. 下载njs插件
    ```
    https://github.com/nginx/njs
    ```

1. 安装

    * configure配置
    ```
    --add-module="path/njs/nginx"
    ```

    编译错误：
    ```
    /bin/sh: 3: /usr/local/openresty/site/lualib/?/init.ljbc: not found
    /bin/sh: 3: /usr/local/openresty/lualib/?.ljbc: not found
    /bin/sh: 3: /usr/local/openresty/lualib/?/init.ljbc: not found
    /bin/sh: 3: /usr/local/openresty/site/lualib/?.lua: not found
    /bin/sh: 3: /usr/local/openresty/site/lualib/?/init.lua: not found
    /bin/sh: 3: /usr/local/openresty/lualib/?.lua: not found
    /bin/sh: 3: /usr/local/openresty/lualib/?/init.lua' -DLUA_DEFAULT_CPATH='/usr/local/openresty/site/lualib/?.so: not found
    /bin/sh: 3: /usr/local/openresty/lualib/?.so' -DNDK_SET_VAR: not found
    objs/Makefile:3065: recipe for target '/home/nginx/openresty-1.13.6.1/add_modules/njs-master/nginx/../build/libnjs.a' failed
    make[2]: *** [/home/nginx/openresty-1.13.6.1/add_modules/njs-master/nginx/../build/libnjs.a] Error 127
    ```

    【解决方法】
    （1） 在目录njs下编译：
    ```
    ./configure

    make
    ```
    生成build/libnjs.a

1. 配置与运行

    http.js
    ```
    function hello(r) {
        var out = {name: "nginx-upstream"};
        r.return(200, JSON.stringify(out));
    }

    export default {hello};
    ```

    nginx.conf配置
    ```
    http {
        js_import http.js;
        
        server {
            listen       80;
            server_name  www.test.com;
            
            location /njs-test {
                js_content http.hello;
            }
        }
    }    
    ```

    运行结果：
    ```
    curl http://127.0.0.1/njs-test
    
    {"name":"nginx-upstream"}
    ```

# 小结

## lua与njs比较

* lua插件由淘宝tengine（或openresty）开发，已经非常成熟商用；并开发了大量的lua插件；

* njs插件不够成熟；

* js代码编写比lua更加方便；

* lua和njs可用作API控制接入层，最好不要用来开发复杂的业务逻辑。

* 可以考虑用nginx + nodejs架构；后端用nodejs开发业务逻辑。—— 效率极高。

# FAQ

不支持E6保留字let
```
nginx: [emerg] SyntaxError: Token "let" not supported in this version in http.js:2
```

# 参考链接

- [http://nginx.org/en/docs/http/ngx_http_js_module.html](http://nginx.org/en/docs/http/ngx_http_js_module.html)




