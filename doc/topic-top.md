# topic - top in container，容器版本的top

推荐一个容器中查看系统信息的工具topic。容器通过cgroups和namespace实现了资源的轻量级隔离和限制，但容器中的/proc文件实际上是宿主机的，因此在执行top命令查看容器运行信息时，部分指标显示不正确，例如启动时间、用户数、平均负载、cpu使用率、内存使用率。
推荐一个容器中查看系统信息的工具topic。

容器通过cgroups和namespace实现了资源的轻量级隔离和限制，但容器中的/proc文件实际上是宿主机的，因此在执行top命令查看容器运行信息时，部分指标显示不正确，例如启动时间、用户数、平均负载、cpu使用率、内存使用率。



目前比较通用的解决方案是通过lxcfs，将容器中相应的文件通过fuse劫持read调用，在打开时显示为容器信息，从而统一解决各种系统状态诊断工具的问题。

考虑到部署lxcfs有一定的成本，topic(top in container)的思路则是改造top命令，去适配容器，读取容器中反映真实运行情况的系统文件，从而展示正确的容器运行信息，对于用户而言成本更低。

如下，在一个1c 1Gi的容器中运行stress --cpu 2，通过topic和top查看容器的运行状态：

topic:

![image](https://user-images.githubusercontent.com/17688273/181487036-d8891cbb-1035-46fd-9ba3-62af3ca9abcf.png)


top:

![image](https://user-images.githubusercontent.com/17688273/181487076-dd3c1c66-f260-44f5-ae63-ca0d21c8a540.png)


可以看到，topic比较好的解决了容器运行信息的问题：

- topic查看的load average是2.03，而top查看到的是1.31(实为宿主机的load average)
- topic查看到的CPU使用率，其us为99.8%，而top查看到的是13.2%(实为宿主机的us信息)
- topic查看到的Mem是1Gi，而top查看到的是16Gi(实为宿主机的内存信息)
- topic查看到的user数是11，而top查看到的user数是1(实为宿主机的当前登录用户数)
- topic查看到的容器运行时间为2days 10:35，而top查看到的是20days 1:57(实为宿主机的运行时间)
- topic和top的进程相关信息显示基本一致。
如果您需要试用，可以下载topic到容器中运行(记得加上执行权限)，好用可以给个Star ^_^

项目地址 https://github.com/silenceshell/topic
