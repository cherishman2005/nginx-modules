# nodejs-cluster


![nodejs-cluster fork](/img/nodejs-cluster.png)


![nodejs-cluster statistics](/img/nodejs-cluster1.png)


	（1）各个worker进程统计自己的请求数。并通过IPC进程间通信发送给master进程；
	（2）master进程收集worker的统计信息，累加统计；

Nginx做多机器集群上的负载均衡，然后用Node.js Cluster来实现单机多进程上的负载均衡。

