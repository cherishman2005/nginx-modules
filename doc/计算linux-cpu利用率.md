# 计算linux cpu利用率

## docker计算cpu使用率

docker自身的cpu使用率的计算：

```
tstart=$(date +%s%N)#获取到当前的时间（纳秒）
cstart=$(cat /sys/fs/cgroup/cpu/mygroup/cpuacct.usage)#当前时刻mygroup这个cgroup已使用的cpu纳秒数
sleep 5
tstop=$(date +%s%N)#获取当前的时间（纳秒）
cstop=$(cat /sys/fs/cgroup/cpu/mygroup/cpuacct.usage)#当前时刻mygroup这个cgroup已使用的cpu纳秒数
bc -l <<EOF
($cstop - $cstart) / ($tstop - $tstart) * 100 #cpu占用的毫秒数/总的时间数即为cpu利用率
```


# FAQ

## 监控cpu计算

容器的cpu使用就必须要在容器内进行，从宿主机是无法计算的，是这样吗？如果是的，那么，这个监控cpu的代码就需要侵入程序代码(容器中跑的业务代码)，这会不会有些无奈啊？另外，这个侵入的代码，所在的线程如果不能被实时调度，则瞬时速度就算的不准确了吧？
 容器的cpu使用就必须要在容器内进行，从宿主机是无法计算的，是这样吗?

* 不是这样的，从宿主机也可以得到容器对应的CPU Cgroup里的值。 

# 参考链接

- [获取docker的内存，cpu使用率](https://www.cnblogs.com/my_life/articles/14945409.html)
