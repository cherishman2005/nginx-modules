# Linux环境下进程的CPU占用率

1、Linux 环境下查看 CPU 信息
1.1、查看 CPU 详细信息
通过 cat /proc/cpuinfo 命令，可以查看 CPU 相关的信息：
```
[root@rh ~]$ cat /proc/cpuinfo
processor : 0
vendor_id : GenuineIntel
cpu family : 6
model : 44
model name : Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz
stepping : 2
cpu MHz : 1596.000
cache size : 12288 KB
physical id : 0
siblings : 8
core id : 0
cpu cores : 4
apicid : 0
initial apicid : 0
fpu : yes
fpu_exception : yes
cpuid level : 11
wp : yes
flags : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc arch_perfmon pebs bts rep_good xtopology nonstop_tsc aperfmperf pni pclmulqdq dtes64 monitor ds_cpl vmx smx est tm2 ssse3 cx16 xtpr pdcm pcid dca sse4_1 sse4_2 popcnt aes lahf_lm arat epb dts tpr_shadow vnmi flexpriority ept vpid
bogomips : 4800.15
clflush size : 64
cache_alignment : 64
address sizes : 40 bits physical, 48 bits virtual
power management:
......
```
在查看到的相关信息中，通常有些信息比较让人迷惑，这里列出一些解释：

physical id: 指的是物理封装的处理器的 id。
cpu cores: 位于相同物理封装的处理器中的内核数量。
core id: 每个内核的 id。
siblings: 位于相同物理封装的处理器中的逻辑处理器的数量。
processor: 逻辑处理器的 id。
我们通常可以用下面这些命令获得这些参数的信息：
```
[root@rh ~]$ cat /proc/cpuinfo | grep "physical id" | sort|uniq
physical id     : 0
physical id     : 1
[root@rh ~]$ cat /proc/cpuinfo | grep "cpu cores" | sort|uniq
cpu cores     : 4
[root@rh ~]# cat /proc/cpuinfo | grep "core id" | sort|uniq
core id          : 0
core id          : 1
core id          : 10
core id          : 9
[root@rh ~]$ cat /proc/cpuinfo | grep "siblings" | sort|uniq
siblings     : 8
[root@rh ~]$ cat /proc/cpuinfo | grep "processor" | sort|uniq
processor     : 0
processor     : 1
processor     : 10
processor     : 11
processor     : 12
processor     : 13
processor     : 14
processor     : 15
processor     : 2
processor     : 3
processor     : 4
processor     : 5
processor     : 6
processor     : 7
processor     : 8
processor     : 9
```
通过上面的结果，可以看出这台机器：

1）有 2 个物理封装的处理器（physical id 有 2 个）；
2）每个物理封装的处理器有 4 个内核（cpu cores 为 4）；
3）每个物理封装的处理器有 8 个逻辑处理器（siblings 为 8），可见台机器的处理器开启了超线程技术，每个内核（core）被划分为了 2 个逻辑处理器（processor）；
4）总共有 16 个逻辑处理器（processor 有 16 个）；
超线程技术：超线程技术就是利用特殊的硬件指令，把两个逻辑内核模拟成两个物理芯片，让单个处理器都能使用线程级并行计算，进而兼容多线程操作系统和软件，减少了CPU的闲置时间，提高的CPU的运行效率。

1.2、查看多核 CPU 信息
可以使用 mpstat 命令或 sar 命令来查看。 具体使用可以通过 man mpstat/sar 来查看。

2、在 Linux 环境下计算进程的 CPU 占用
2.1、通过 /proc/stat 文件查看所有的 CPU 活动信息
下面实例数据是内核 2.6.24-24 版本以上的：
```
[root@rh ~]$ cat /proc/stat
cpu  223447 240 4504182 410802165 59753 412 586209 0 0
cpu0 17625 11 193414 25755165 34590 72 16780 0 0
cpu1 12412 9 139234 25860912 1139 27 15697 0 0
cpu2 14953 0 558618 25310182 230 73 87851 0 0
cpu3 4088 0 138057 25873479 3404 19 15862 0 0
cpu4 13417 17 161123 25829756 10429 22 17155 0 0
cpu5 23935 5 331837 25605479 369 20 26095 0 0
cpu6 19320 0 319132 25619728 2060 18 30690 0 0
cpu7 18430 0 324300 25616142 4170 20 33081 0 0
cpu8 31940 0 661718 25231911 1474 15 67661 0 0
cpu9 8079 87 152035 25829339 237 19 16308 0 0
cpu10 4247 24 140052 25869331 168 18 15059 0 0
cpu11 16305 4 401473 25567904 207 14 34674 0 0
cpu12 8227 0 157815 25853353 319 10 14808 0 0
cpu13 18894 46 479709 25354279 653 26 157885 0 0
cpu14 7000 30 204760 25769743 159 16 21059 0 0
cpu15 4567 1 140899 25855452 138 16 15536 0 0
intr 1146734384 997 0 0 1 1 0 0 0 1 0 0 0 0 0 1302607 0 0 262087 0 712704 0 0 0 <...省略若干数据...>
ctxt 89793364
btime 1366591448
processes 27283
procs_running 1
procs_blocked 0
softirq 1262462448 0 63122856 50789329 1074176388 225020 0 461213 9535581 76130 64075931
```
第一行的数据表示的是 CPU 总的使用情况。我们来解释一下这行数据各数值的含义：

1）这些数值的单位都是 jiffies，jiffies 是内核中的一个全局变量，用来记录系统启动以来产生的节拍数，在 Linux 中，一个节拍大致可以理解为操作系统进程调度的最小时间片，不同的 Linux 系统内核这个值可能不同，通常在 1ms 到 10ms 之间。
2）cpu 223447 240 4504182 410802165 59753 412 586209 0 0
user(223447) 从系统启动开始累积到当前时刻，处于用户态的运行时间，不包含 nice 值为负的进程。
nice(240) 从系统启动开始累积到当前时刻，nice 值为负的进程所占用的 CPU 时间。
system(4504182) 从系统启动开始累积到当前时刻，处于核心态的运行时间。
idle(410802165) 从系统启动开始累积到当前时刻，除 IO 等待时间以外的其他等待时间。
iowait(59753) 从系统启动开始累积到当前时刻，IO 等待时间。(since 2.5.41)
irq(412) 从系统启动开始累积到当前时刻，硬中断时间。(since 2.6.0-test4)
softirq(586209) 从系统启动开始累积到当前时刻，软中断时间。(since 2.6.0-test4)
stealstolen(0) Which is the time spent in other operating systems when running in a virtualized environment.(since 2.6.11)
guest(0) Which is the time spent running a virtual CPU for guest operating systems under the control of the Linux kernel.(since 2.6.24)
从以上信息我们可以得到总的 CPU 活动时间为：
```
totalCPUTime = user + nice + system + idle + iowait + irq + softirq + stealstolen + guest
```

2.2、通过 /proc/[PID]/stat 文件查看某一进程的 CPU 活动信息
2.2.1、存储进程信息的文件目录

Linux 系统贯彻“一切都是文件”的思想，所有的进程的运行状态也都可以通过读取文件来获取。 /proc 文件系统是一个伪文件系统，它只存在内存当中，而不占用外存空间。它以文件系统的方式为内核与进程提供通信的接口。用户和应用程序可以通过 /proc 得到系统的信息，并可以改变内核的某些参数。

在 /proc/[PID]/ 目录下的各个文件记录着 这个进程的各项运行指标。 是进程号。

2.2.2、查看进程运行的详细信息

通过查看 /proc/[PID]/stat 文件，可以进程运行的详细信息，其中就包括 CPU 占用信息。 比如：
```
[root@rh ~]$ cat /proc/1/stat
1 (init) S 0 1 1 0 -1 4202752 3026 2635222 9 483 5 165 102346 3188016 20 0 1 0 1 19820544 384 18446744073709551615 1 1 0 0 0 0 0 4096 536962595 18446744073709551615 0 0 0 4 0 0 34 0 0
```

/proc/[PID]/stat 文件信息解释

看到上面这些信息，肯定会很迷惑，不知道每个字段都是什么意思。

1）我们可以通过 man 5 proc 命令查看文档，找到 /proc/[pid]/stat 节点，就可以看到各字段的意思了。如：
```
/proc/[pid]/stat
Status information about the process.  This is used by ps(1).  It is defined in /usr/src/linux/fs/proc/array.c.
The fields, in order, with their proper scanf(3) format specifiers, are:
pid %d      The process ID.
comm %s     The  filename  of the executable, in parentheses.  This is visible whether or not the executable is swapped out.
state %c    One character from the string "RSDZTW" where R is running, S is sleeping in an interruptible  wait,  D  is waiting in uninterruptible disk sleep, Z is zombie, T is traced or stopped (on a signal), and W is paging.
ppid %d     The PID of the parent.
......
```

2）具体解释，一个示例：
```
pid=6873 进程(包括轻量级进程，即线程)号
comm=a.out 应用程序或命令的名字。
task_state=R 任务的状态，R:runnign, S:sleeping (TASK_INTERRUPTIBLE), D:disk sleep (TASK_UNINTERRUPTIBLE), T: stopped, T:tracing stop, Z:zombie, X:dead。
ppid=6723 父进程ID。
pgid=6873 线程组号。
sid=6723 该任务所在的会话组 ID。
tty_nr=34819(pts/3) 该任务的 tty 终端的设备号，INT（34817/256）= 主设备号，（34817-主设备号）= 次设备号。
tty_pgrp=6873 终端的进程组号，当前运行在该任务所在终端的前台任务(包括 shell 应用程序)的 PID。
task->flags=8388608 进程标志位，查看该任务的特性。
min_flt=77 该任务不需要从硬盘拷数据而发生的缺页（次缺页）的次数。
cmin_flt=0 累计的该任务的所有的 waited-for 进程曾经发生的次缺页的次数目。
maj_flt=0 该任务需要从硬盘拷数据而发生的缺页（主缺页）的次数。
cmaj_flt=0 累计的该任务的所有的 waited-for 进程曾经发生的主缺页的次数目。
utime=1587 该任务在用户态运行的时间，单位为 jiffies。
stime=1 该任务在核心态运行的时间，单位为 jiffies。
cutime=0 累计的该任务的所有的 waited-for 进程曾经在用户态运行的时间，单位为 jiffies。
cstime=0 累计的该任务的所有的 waited-for 进程曾经在核心态运行的时间，单位为 jiffies。
priority=25 任务的动态优先级。
nice=0 任务的静态优先级。
num_threads=3 该任务所在的线程组里线程的个数。
it_real_value=0 由于计时间隔导致的下一个 SIGALRM 发送进程的时延，以 jiffy 为单位。
start_time=5882654 该任务启动的时间，单位为 jiffies。
vsize=1409024（page） 该任务的虚拟地址空间大小。
rss=56(page) 该任务当前驻留物理地址空间的大小；Number of pages the process has in real memory,minu 3 for administrative purpose. 这些页可能用于代码，数据和栈。
rlim=4294967295（bytes） 该任务能驻留物理地址空间的最大值。
start_code=134512640 该任务在虚拟地址空间的代码段的起始地址。
end_code=134513720 该任务在虚拟地址空间的代码段的结束地址。
start_stack=3215579040 该任务在虚拟地址空间的栈的结束地址。
kstkesp=0 esp(32 位堆栈指针) 的当前值, 与在进程的内核堆栈页得到的一致。
kstkeip=2097798 指向将要执行的指令的指针, EIP(32 位指令指针)的当前值。
pendingsig=0 待处理信号的位图，记录发送给进程的普通信号。
block_sig=0 阻塞信号的位图。
sigign=0 忽略的信号的位图。
sigcatch=082985 被俘获的信号的位图。
wchan=0 如果该进程是睡眠状态，该值给出调度的调用点。
nswap 被 swapped 的页数，当前没用。
cnswap 所有子进程被 swapped 的页数的和，当前没用。
exit_signal=17 该进程结束时，向父进程所发送的信号。
task_cpu(task)=0 运行在哪个 CPU 上。
task_rt_priority=0 实时进程的相对优先级别。
task_policy=0 进程的调度策略，0=非实时进程，1=FIFO实时进程；2=RR实时进程
```

2.2.3、关于进程占用 CPU 的相关信息

在上述的时间中，这些信息会在计算 CPU 占用率时用到：

pid 进程号。
utime 该任务在用户态运行的时间，单位为 jiffies。
stime 该任务在核心态运行的时间，单位为 jiffies。
cutime 累计的该任务的所有的 waited-for 进程曾经在用户态运行的时间，单位为 jiffies。
cstime 累计的该任务的所有的 waited-for 进程曾经在核心态运行的时间，单位为 jiffies。
该进程的 CPU 占用时间（该值包括其所有线程的 CPU 时间）：
```
processCPUTime = utime + stime + cutime + cstime
```

2.3、通过 /proc/[PID]/task/[TID]/stat 文件查看某一进程下的某一线程的活动信息
该文件包含了某一轻量级进程（lwp，即通常所说的线程）所有的活动信息，该文件中的所有值都是从系统启动开始累计到当前时刻。该文件的内容格式以及各字段的含义与 /proc/[PID]/stat 文件类似。该文件中的 tid 字段表示的是轻量级线程号。

该线程的 CPU 占用时间：
```
threadCPUTime = utime + stime
```

2.4、单核情况下 CPU 使用率的计算
2.4.1、基本思想

首先，通过读取 /proc/stat 文件获取总的 CPU 时间，读取 /proc/[PID]/stat 获取进程 CPU 时间，读取 /proc/[PID]/task/[TID]/stat 获取线程 CPU 时间。然后，采样两个足够短的时间间隔的 CPU 快照与进程或线程快照来计算其 CPU 使用率。

2.4.2、计算总的 CPU 使用率 totalCPUUse

1）采样两个足够短的时间间隔的 CPU 快照，即读取 /proc/stat 文件，获取两个时间点的下列数据：
```
CPUT1 (user1, nice1, system1, idle1, iowait1, irq1, softirq1, stealstolen1, guest1);
CPUT2 (user2, nice2, system2, idle2, iowait2, irq2, softirq2, stealstolen2, guest2);
```
2）计算总的 CPU 时间 totalCPUTime：
```
CPUTime1 = user1 + nice1 + system1 + idle1 + iowait1 + irq1 + softirq1 + stealstolen1 + guest1;
CPUTime2 = user2 + nice2 + system2 + idle2 + iowait2 + irq2 + softirq2 + stealstolen2 + guest2;
totalCPUTime = CPUTime2 – CPUTime1;
```
3）计算 CPU 空闲时间 idleCPUTime：

idleCPUTime = idle2 – idle1;

4）计算总的 CPU 使用率 totalCPUUse：

totalCPUUse = (totalCPUTime – idleCPUTime) / totalCPUTime;

2.4.3、计算某一进程的 CPU 使用率 processCPUUse

1）采样两个足够短的时间间隔的 CPU 快照和对应的进程快照，即读取 /proc/stat 文件，获取两个时间点的下列数据：

CPUT1 (user1, nice1, system1, idle1, iowait1, irq1, softirq1, stealstolen1, guest1);
CPUT2 (user2, nice2, system2, idle2, iowait2, irq2, softirq2, stealstolen2, guest2);
读取 /proc/[PID]/stat 文件，获取两个时间点的下列数据：

ProcessT1 (utime1, stime1, cutime1, cstime1);
ProcessT2 (utime2, stime2, cutime2, cstime2);
2）计算总的 CPU 时间 totalCPUTime 和进程时间 processTime：

CPUTime1 = user1 + nice1 + system1 + idle1 + iowait1 + irq1 + softirq1 + stealstolen1 + guest1;
CPUTime2 = user2 + nice2 + system2 + idle2 + iowait2 + irq2 + softirq2 + stealstolen2 + guest2;
totalCPUTime = CPUTime2 – CPUTime1;

processTime1 = utime1 + stime1 + cutime1 + cstime1;
processTime2 = utime2 + stime2 + cutime1 + cstime2;
processTime = processTime2 – processTime1;

3）计算该进程的 CPU 使用率 processCPUUse:

processCPUUse = processTime / totalCPUTime;

2.4.4、计算某一线程的 CPU 使用率 threadCPUUse

1）采样两个足够短的时间间隔的 CPU 快照和对应的线程快照，即读取 /proc/stat 文件，获取两个时间点的下列数据：

CPUT1 (user1, nice1, system1, idle1, iowait1, irq1, softirq1, stealstolen1, guest1);
CPUT2 (user2, nice2, system2, idle2, iowait2, irq2, softirq2, stealstolen2, guest2);
读取 /proc/[PID]/task/[TID]/stat 文件，获取两个时间点的下列数据：

threadT1 (utime1, stime1);
threadT2 (utime2, stime2);
2）计算总的 CPU 时间 totalCPUTime 和线程时间 threadTime：

CPUTime1 = user1 + nice1 + system1 + idle1 + iowait1 + irq1 + softirq1 + stealstolen1 + guest1;
CPUTime2 = user2 + nice2 + system2 + idle2 + iowait2 + irq2 + softirq2 + stealstolen2 + guest2;
totalCPUTime = CPUTime2 – CPUTime1;

threadTime1 = utime1 + stime1;
threadTime2 = utime2 + stime2;
threadTime = threadTime2 – threadTime1;

3）计算该线程的 CPU 使用率 threadCPUUse:

threadCPUUse = threadTime / totalCPUTime;

2.5、多核情况下 CPU 使用率的计算
2.5.1、基本思想

首先，通过读取 /proc/stat 文件获取总的 CPU 时间，读取 /proc/[PID]/stat 获取进程 CPU 时间，读取 /proc/[PID]/task/[TID]/stat 获取线程 CPU 时间，读取 /proc/cpuinfo 获取 CPU 个数。

在多核情况下计算进程或线程的 CPU 使用率，用上面的方式得到的通常是相对于 CPU 所有核的总共时间的占用率，而我们通常习惯得到进程或线程对某一个单核的占用率。所以我们可以按上面的方式计算得到 CPU 占用率，然后把结果乘上 CPU 的核数，即可得到进程或线程相对于一个单核的占用率。

2.5.2、计算总的 CPU 使用率

同 2.4.2。

2.5.3、计算某一进程的 CPU 使用率 mProcessCPUUse

1）同 2.4.3 计算某一进程的 CPU 使用率 processCPUUse；

2）读取 /proc/cpuinfo 文件获取逻辑 CPU（processor） 的个数（参见 1.1）： processorNum

3）多核情况下该进程的 CPU 使用率 mProcessCPUUse：

mProcessCPUUse = processCPUUse * processorNum;

2.5.4、计算某一线程的 CPU 使用率 mThreadCPUUse

1）同 2.4.4 计算某一线程的 CPU 使用率 threadCPUUse；

2）读取 /proc/cpuinfo 文件获取逻辑 CPU（processor） 的个数（参见 1.1）： processorNum

3）多核情况下该线程的 CPU 使用率 mThreadCPUUse：

mThreadCPUUse = threadCPUUse * processorNum;

2.6、问题
1）不同内核版本 /proc/stat 文件格式不一致。/proc/stat 文件中第一行是总的 CPU 使用情况。
各个内核版本都有的 4 个字段：user, nice, system, idle；
2.5.41 版本新增字段：iowait；
2.6.0-test4 版本新增字段：irq, softirq；
2.6.11 版本新增字段：stealstolen；
2.6.24 版本新增字段：guest；
2）/proc/[PID]/task 目录是 Linux 2.6.0-test6 之后才有的功能。
3）关于 CPU 使用率为负的情况，解决方案是如果出现负值，连续采样计算 CPU 使用率直到为非负。
4）有些线程生命周期较短，可能在采样期间就已经死掉了。
2.7、一个计算总的 CPU 占用率的小程序
```
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <time.h>
 
typedef struct procstat {
     char processorName[20];
     unsigned int user;
     unsigned int nice;
     unsigned int system;
     unsigned int idle;
     unsigned int iowait;
     unsigned int irq;
     unsigned int softirq;
     unsigned int stealstolen;
     unsigned int guest;
} Procstat;
 
Procstat getCPUStatus() {
     // Get "/proc/stat" info.
     FILE* inputFile = NULL;
      
     chdir("/proc");
     inputFile = fopen("stat", "r");
     if (!inputFile) {
          perror("error: Can not open file.\n");
     }
 
     char buff[1024];
     fgets(buff, sizeof(buff), inputFile); // Read 1 line.
     printf(buff);     
     Procstat ps;
     sscanf(buff, "%s %u %u %u %u %u %u %u %u %u", ps.processorName, &ps.user, &ps.nice, &ps.system, &ps.idle, &ps.iowait, &ps.irq, &ps.softirq, &ps.stealstolen, &ps.guest); // Scan from "buff".     
     printf("user: %u\n", ps.user);
 
     fclose(inputFile);
 
     return ps;
      
}
 
float calculateCPUUse(Procstat ps1, Procstat ps2) {
     unsigned int totalCPUTime = (ps2.user + ps2.nice + ps2.system + ps2.idle + ps2.iowait + ps2.irq + ps2.softirq + ps2.stealstolen + ps2.guest) - (ps1.user + ps1.nice + ps1.system + ps1.idle + ps1.iowait + ps1.irq + ps1.softirq + ps1.stealstolen + ps1.guest);
     unsigned int idleCPUTime = ps2.idle - ps1.idle;
 
     float CPUUse = ((float) totalCPUTime - (float) idleCPUTime) / (float) totalCPUTime;
 
     printf("totalCPUTime: %u\nidleCPUTime: %u\n", totalCPUTime, idleCPUTime);
 
     return CPUUse;
}
 
int main(int argc, char* argv[]) {
     printf("Test CPU\n");
 
     // Get processor num.
     int processorNum = sysconf(_SC_NPROCESSORS_CONF); // "unistd.h" is required.
     printf("Processors: %d\n", processorNum);
      
 
     // Test
     Procstat ps1, ps2;
     int i = 0;
     for (i = 0; i <= 100000; i++) {
           
          srand((unsigned) time(NULL));
          int m = rand() % 100000;
          int n = 1 + rand() % 100000;
          int k = m / n;
           
          if (i == 10) {
               ps1 = getCPUStatus();
          }
 
          if (i == 10000) {
               ps2 = getCPUStatus();     
          }
     }
     float CPUUse = calculateCPUUse(ps1, ps2);
     printf("CPUUse: %f\n", CPUUse);
 
 
     return 0;
}
```

3、Linux 环境查看进程运行相关信息
3.1、使用 ps 命令查看进程信息
几个常用参数：

a: 与终端无关的所有进程。
u: 有效用户（effective user）的相关进程。
x: 通常与 a 参数因使用，可列出较完整的信息。
-e: 选择所有进程。
-L: 显示线程，一般是 LWP 或 NLWP 列。
-o: 用户自定义显示选项。
示例1）列出所有当前所有正在内存中的进程

[root@rh ~]$ ps aux
USER PID %CPU %MEM VSZ RSS TTY STAT START TIME COMMAND
root 1 0.0 0.0 19356 1536 ? Ss Apr22 0:01 /sbin/init
root 2 0.0 0.0 0 0 ? S Apr22 0:00 [kthreadd]
root 3 0.0 0.0 0 0 ? S Apr22 0:00 [migration/0]
root 4 0.0 0.0 0 0 ? S Apr22 0:00 [ksoftirqd/0]
root 5 0.0 0.0 0 0 ? S Apr22 0:00 [migration/0]
root 6 0.0 0.0 0 0 ? S Apr22 0:00 [watchdog/0]
root 7 0.0 0.0 0 0 ? S Apr22 0:00 [migration/1]
root 8 0.0 0.0 0 0 ? S Apr22 0:00 [migration/1]
......
示例2）列出进程号为 13560 这个进程的所有线程及 CPU 占用率

[root@rh ~]$ ps -eLo pid,lwp,pcpu | grep 13560
13560 13560 49.5
3.2、使用 top 命令查看进程信息
几个常用的参数：

-d: 后面接秒数，就是整个进程画面更新的频率。默认是 5 秒。
-b: 以批处理的方式执行 top，还有更多的参数可用。通常会搭配数据流重导向，将批处理的结果输出为文件。
-n: 与 -b 搭配，意义是，需要进行几次 top 的输出结果。
-p: 指定某个 PID 来进行观察监测。
在 top 执行过程中可以使用的按键命令：
?: 显示在 top 中可以输入的按键命令。
P: 按照 CPU 的使用资源排序显示。
M: 按内存（Memory）的使用资源排序显示。
N: 按 PID 来排序。
T: 按该进程使用的 CPU 时间积累（TIME+）排序。
k: 给某个 PID 一个信号（signal）。
r: 给某个 PID 重新确定一个值。
1: 显示所有 CPU 占用信息。
示例1）将 top 命令执行两次，然后将结果输出到 /top_result.data。

[root@rh ~]$ top -b -n 2 &gt; /top_result.data
示例2）监测进程 13620

[root@rh ~]$ top -d 2 -p 13620
top - 16:27:35 up 4 days, 7:43, 2 users, load average: 0.35, 0.47, 0.44
Tasks: 1 total, 1 running, 0 sleeping, 0 stopped, 0 zombie
Cpu(s): 0.1%us, 3.1%sy, 0.0%ni, 96.5%id, 0.0%wa, 0.0%hi, 0.3%si, 0.0%st
Mem: 16320632k total, 1790796k used, 14529836k free, 233168k buffers
Swap: 8232952k total, 0k used, 8232952k free, 941540k cached
 
PID USER PR NI VIRT RES SHR S %CPU %MEM TIME+ COMMAND
13620 test1370 20 0 11060 944 760 R 53.4 0.0 0:04.78 netperf
本文参考
http://www.blogjava.net/fjzag/articles/317773.html

http://www.brokestream.com/procstat.html

http://blog.csdn.net/zg_hover/article/details/4356210

《鸟哥Linux私房菜》
