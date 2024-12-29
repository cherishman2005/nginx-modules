# ingress-nginx

* 节点变更不用nginx reload重启；ingress-controller会触发nginx-lua upstream模块更新后端节点信息。
* 只有路由规则变化时，才需要nginx-reload。


## Resource caches

![image](https://github.com/user-attachments/assets/8968c78a-3e04-46fa-9d01-6862104fdcf4)


# 参考链接

- [https://docs.nginx.com/nginx-ingress-controller/overview/design/](https://docs.nginx.com/nginx-ingress-controller/overview/design/)
