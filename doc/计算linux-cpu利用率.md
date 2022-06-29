# 计算linux cpu利用率

## 系统cpu计算

Usage：cgroup 的 CPU 占用率，占了物理机的多少CPU 【不是docker中自身cpu的使用率：占被分配的cpu的比例】
上面我们拿到了能用多少个核，自然知道 cgroup 的占用上限，只要知道 cgroup 用了物理机的多少 CPU 就可以知道饱和程度了。要获取这个首先需要知道一段时间内物理机用了多少CPU时间，然后获得 cgroup 用了多少CPU时间，最后相除。

* 物理机使用的CPU时间从 /proc/stat 里获取，将里面 cpu 那一行的数字相加并且乘以物理机CPU核数即可得到从开机到现在用的CPU总时间。可以设置一秒的时间间隔，求差即可得到这一秒内用的CPU时间。
* cgroup 使用的CPU时间可以从 cpuacct/cpuacct.usage 中获得，也是求一段时间的差即可。

### docker计算系统cpu使用率

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
- [https://time.geekbang.org/column/article/313255](https://time.geekbang.org/column/article/313255)
- [https://blog.lichao.xin/back-end/docker/docker-05/](https://blog.lichao.xin/back-end/docker/docker-05/)
- [CPU使用率原理及计算方式](https://www.cnblogs.com/gatsby123/p/11127158.html)
