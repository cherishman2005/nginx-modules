# apt get update

/etc/resolv.conf 

/etc/apt/sources.list

/etc/hosts

## docker ubunt12.04更新失败

https://askubuntu.com/questions/1385440/ubuntu-sudo-apt-get-update-404-not-found-problem

Most of the repositories and PPAs in your sources.list are no longer available and are throwing errors. I'd recommend restoring the default repositories.

First, restore the default focal repositories using these commands:

mkdir ~/solution
cd ~/solution/

cat << EOF > ~/solution/sources.list
deb http://archive.ubuntu.com/ubuntu/ focal main restricted universe multiverse
deb-src http://archive.ubuntu.com/ubuntu/ focal main restricted universe multiverse

deb http://archive.ubuntu.com/ubuntu/ focal-updates main restricted universe multiverse
deb-src http://archive.ubuntu.com/ubuntu/ focal-updates main restricted universe multiverse

deb http://archive.ubuntu.com/ubuntu/ focal-security main restricted universe multiverse
deb-src http://archive.ubuntu.com/ubuntu/ focal-security main restricted universe multiverse

deb http://archive.ubuntu.com/ubuntu/ focal-backports main restricted universe multiverse
deb-src http://archive.ubuntu.com/ubuntu/ focal-backports main restricted universe multiverse

deb http://archive.canonical.com/ubuntu focal partner
deb-src http://archive.canonical.com/ubuntu focal partner
EOF

sudo rm /etc/apt/sources.list
sudo cp ~/solution/sources.list /etc/apt/sources.list
Remove all the PPAs in your system:

sudo mv /etc/apt/sources.list.d/* ~/solution
Update the repositories:

sudo apt update
Now there should be no errors.


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
