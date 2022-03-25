# 对于pthread_kill的一些误解

今天在写代码的时候，想通过pthread_kill接口杀死一个线程，因为之前学习《Linux/Unix系统编程手册》的时候，里面对于该接口的功能介绍是如下：

函数pthread_kill()向同一进程下的另一线程发送信号。因为仅在同一进程中可保证线程ID的唯一性，所以无法调用pthread_kill向其他进程中的线程发送信号。

所以就感觉pthread_kill与kill差不多，不过一个是向线程发送，一个是向进程发送，但是在使用中，与设想的完全不一样。上测试代码：

```
#include <stdio.h>
#include <signal.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>
 
void *func(void *argv)
{
    while (1) {
        printf("I am thread, %ld\n", pthread_self());
        sleep(1);
    }
    return NULL;
}
 
int main()
{
    pthread_t t;
    pthread_create(&t, NULL, func, NULL);
    sleep(10);
    printf("send kill to thread\n");
    pthread_kill(t, SIGKILL);
    sleep(10);
    printf("end\n");
}
```

代码很简单，就是创建一个线程，一直打印信息，然后过10s，主线程杀死子线程，然后再过10s退出。理想中的运行状态就是子线程运行10s，然后被主线程杀死，再过10s，主线程退出。但是实际上，最终的结果是，10s后，pthread_kill直接把主线程给杀死了，根本不会运行最后的10s睡眠和打印。后来还尝试了SIGTERM信号，结果是一样的。

对这个结果感到很意外，于是只能去查了man手册，man手册的解释如下：

    Signal dispositions are process-wide: if a signal handler is installed, the handler will be invoked in the thread thread, but if the disposition of the signal is "stop", "continue", or "terminate", this action will affect the whole process.

可以看到，在man手册中清楚的说明了，如果信号是“stop”， “continue”或者“terminate”这种类型的信号，那么将会影响整个进程，所以整个进程都被杀掉了。

其实为了达到之前的目的，可以使用pthread_cancel函数，直接向指定线程发送取消请求就可以了。
