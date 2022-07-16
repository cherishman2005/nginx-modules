 # rdtsc
 
 ## constant_tsc
 
 ```
 cat /proc/cpuinfo |grep constant_tsc
 ```
 
 ## TSC的坑

【坑1】比如有的CPU会根据机器负载情况动态调节工作频率， 那么单位时间CPU的指令周期数就会发生变化，也就很难将其转换成时间。另外，CPU进入休眠再次重启后，TSC会清零。
 
 # 参考链接
 
 - [细说RDTSC的坑](http://www.wangkaixuan.tech/?p=901)
