# docker模式

docker network下 host模式与bridge不模式一样

bridge模式如果运行容器不指定自定义网桥，默认为docker0

host模式下如果运行容器，和宿主机共享网段。

host模式：
```
docker run -itd --net=host --name nginx5 nginx
```
