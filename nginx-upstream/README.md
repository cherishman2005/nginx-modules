# 负载均衡小结

现在 跑了2个拉流解码进程，其他同事负责的分配调度进程。

28个进程，居然一个分配了3个进程，另外一个分配了25个进程。   
nginx平滑加权轮询的 就不借鉴下。

分配调度进程用go写的，网上多的是平滑加权负载均衡示例。成熟稳定。
