

```
#include <pthread.h>

int pthread_join(pthread_t thread, void **value_ptr);
int pthread_detach(pthread_t thread);
int pthread_attr_setdetachstate(pthread_attr_t *attr, int detachstate);
int pthread_attr_setschedpolicy(pthread_attr_t *attr, int policy);
```

说明：
1. value_ptr 是指向线程返回值的指针，就是pthread_exit 中的value_ptr 参数，或者是return语句中的返回值, 不能为线程函数的局部变量，否者该指针指向的栈内存不可知， 会有内存错误。

2. 在pthread_join manul 的DESCRIPTION 中，It is unspecified whether a thread that has exited but remains unjoined countsagainst {PTHREAD_THREADS_MAX}. 意思是当一个线程被创建时是可汇合的joinable（默认的attribution）， 其他线程或父线程有没有调用pthread_join去做相关资源释放（pthread id等）， 该线程运行结束后资源就得不到释放，所在进程的pthread id数目就可能会累积到达最大数目PTHREAD_THREADS_MAX，此时系统就不能再创建线程了，因为pthread id等资源被用光了，这是在多线程编程中很常见的bug之一。

3. 如果子线程的状态设置为detached,脱离线程却象守护进程：该线程运行结束后会自动释放所有资源，我们不能等待它们终止。显然当我们想让线程自己自动释放所有资源时就用pthread_attr_setdetachstate PTHREAD_CREATE_DETACHED属性

4. FIFO，RR 都是实时调度队列的，实时进程调度队列，是从优先级最高的进程运行，如果当前运行的是FIFO进程，如果他不主动让出cpu，其他进程都不能运行 ，如果是RR（时间片轮转）的，则不会一直独占cpu， 运行一段时间会被切换出来。


# 参考链接

- [线程创建之重要属性PTHREAD_CREATE_DETACHED](https://blog.csdn.net/xuyunzhang/article/details/8804298)
