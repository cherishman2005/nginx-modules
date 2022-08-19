# apt get update

## docker ubunt12.04更新失败


### 解决方法参考

- [https://askubuntu.com/questions/597323/command-sudo-apt-get-update-runs-with-some-error](https://askubuntu.com/questions/597323/command-sudo-apt-get-update-runs-with-some-error)


##  问题出现原因

最近利用Dockerfile构建镜像，其基础镜像是ubuntu18.04，在apt-get install时，出现了
E: Failed to fetch http://archive.ubuntu.com/ubuntu/dists/trusty-security/main/source/Sources Hash Sum mismatch 的错误。

这是由于GFW的原因，我们获取到的文件和官方源上的文件大小不一致，导致md5值也不一样，所以会出现Hash Sum mismatch。

解决办法
可以使用国内163，阿里或者清华的源替换，代码如下：
```
#更新源
#RUN sed -i s:/archive.ubuntu.com:/mirrors.aliyun.com/ubuntu:g /etc/apt/sources.list
RUN sed -i s:/archive.ubuntu.com:/mirrors.tuna.tsinghua.edu.cn/ubuntu:g /etc/apt/sources.list
RUN cat /etc/apt/sources.list
RUN apt-get clean
RUN apt-get -y update --fix-missing
```
