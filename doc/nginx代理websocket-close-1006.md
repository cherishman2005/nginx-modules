# nginx代理websocket，出现websocket: close 1006 (abnormal closure)错误

部署到服务器上的websocket程序，未加心跳，发现隔一阵子就会断开，或十分钟、或半小时的概率，出现1006 (abnormal closure)错误。

```
[Serve] conn closed with err. err:websocket: close 1006 (abnormal closure): unexpected EOF"
```

代理配置如下：

```
location / {
        proxy_set_header Host $host;
        proxy_pass       http://test_view;
        proxy_redirect   off;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    location /websocket/v2 {
        proxy_redirect off;
        proxy_read_timeout 60s;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-Proto   "http";
        proxy_set_header Host                $host;
        proxy_set_header X-NG-PRX-Host $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_pass http://test_view;
    }
```

## 原因

代理参数的 proxy_read_timeout 默认是60s，只要超过这段时间没有通信，就会自动断开tcp连接。

## 解决办法

1. 修改参数 proxy_read_timeout，比如在代理设置添加 proxy_read_timeout 300s;

2. 添加心跳，心跳的间隔要小于默认的读超时proxy_read_timeout 60s，这样程序就会可以在60s的时间窗口内读到数据，不会被nginx断开连接。

