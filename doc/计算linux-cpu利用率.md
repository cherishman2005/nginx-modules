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


# 参考链接

- [获取docker的内存，cpu使用率](https://www.cnblogs.com/my_life/articles/14945409.html)
