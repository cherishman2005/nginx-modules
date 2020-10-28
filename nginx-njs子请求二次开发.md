# nginx njs子请求

nginx subrequest子请求，在nginx做API网关，控制接入层时常用到。

（1）nignx c module实现subrequest的子请求，一般使用在音视频实时转封装场景；—— c module性能好。

（2）nginx-lua(openresty) ngx.location.capture常用于lua脚本实现一些接入控制。

（3）nginx-njs用js实现实现一些接入控制，权限管理（如token鉴权逻辑）。

本文主要采用nginx njs进行subrequest的开发介绍。

## nginx njs的单个子请求

![njs的单个子请求](/img/njs-subrequest.png)

njs_http.js
```
function get_token(r) {
   r.subrequest('/auth')
   .then(reply => {
       r.return(reply.status, reply.responseBody);
   })
  .catch(_ => r.return(500));
}

function authorize(r) {
    var reply = {uid: "8888", username: "zhangbiwu", token: "ABC"};
    r.return(200, JSON.stringify(reply));
}

export default {get_token, authorize};
```

nginx_upstream.conf
```
js_import njs_http.js;

server {
    listen       8004;
    server_name  localhost;
    charset utf-8;

    location /get-token {
        js_content njs_http.get_token;
    }

    location /auth {
        proxy_pass http://127.0.0.1:8005;
    }

}

server {
    listen       8005;
    server_name  localhost;
    charset utf-8;

    location /auth {
        js_content njs_http.authorize;
    }

}
```

## nginx njs并行子请求

![njs的并行子请求](/img/njs-pipeline-subrequest.png)

请求是异步

njs_http.js
```
function join(r) {
        join_subrequests(r, ['/foo', '/bar']);
}

function join_subrequests(r, subs) {
    var parts = [];

    function done(reply) {
        parts.push({ uri:  reply.uri,
                     code: reply.status,
                     body: reply.responseBody });

        if (parts.length == subs.length) {
            r.return(200, JSON.stringify(parts));
        }
    }

    for (var i in subs) {
        r.subrequest(subs[i], done);
    }
}

export default {join}
```

运行示例：
```
curl http://127.0.0.1:8004/join

[{"uri":"/foo","code":200,"body":"foo"},{"uri":"/bar","code":200,"body":"bar"}]
```

在并发子请求的开发时，并发的请求不要太多。因为要等到最后一个请求返回才给客户端回响应。


## nginx njs串行子请求

![njs的串行子请求](/img/njs-serial-subrequest.png)

适用于有些场景：需要2个串行请求才能获取到最终结果，同时第2个请求依赖第1个请求的场景。

njs_http.js
```
function process(r) {
    r.subrequest('/auth')
        .then(reply => JSON.parse(reply.responseBody))
        .then(response => {
            r.log("/auth response=" + JSON.stringify(response));
            if (!response['token']) {
                throw new Error("token is not available");
            }
            return response['token'];
        })
    .then(token => {
        r.subrequest('/foo', `token=${token}`)
            .then(reply => r.return(reply.status, reply.responseBody));
    })
    .catch(e => r.return(500, e));
}
```

在串行请求的开发时，请求不要太多。因为请求相互依赖，并且等到最后一个请求返回才给客户端回响应。

# 小结

* nginx-njs在做API网关，token鉴权开发效率高，也方便后期维护。

# Author

zhangbiwu

欢迎交流nginx二次开发技术。
