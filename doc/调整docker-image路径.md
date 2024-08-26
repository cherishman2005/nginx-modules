# 调整docker-image路径

docker images的存储过大，现有存储位置已经不足以支撑。需要修改镜像存储的位置。

默认情况下docker pull images存储在/var/lib/docker/ 目录下。如需修改存储位置在需要手动创建/etc/docker/daemon.json 文件，可以指定images存储的位置。

1. 查询docker路径

docker info
```
 Docker Root Dir: /var/lib/docker
```

2. 修改docker image路径

vi  /etc/docker/daemon.json
```
{
  "data-root": "/data2/docker"
}
```

3. 重启docker

```
service docker restart
```

4. 查看路径是否生效

* `docker info `
```
 Docker Root Dir: /data2/docker
```
