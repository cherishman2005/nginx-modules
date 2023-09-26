# install nginx in docker

* docker拉起nginx，nginx不应该在后台运行；使用`sleep infinity`

```
cd bin/ && cp ../conf/nginx.conf /data/services/nginx/openresty/nginx/conf/ && /data/services/nginx/openresty/nginx/sbin/nginx; sleep infinity
```
