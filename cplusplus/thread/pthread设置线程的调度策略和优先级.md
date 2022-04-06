# pthread设置线程的调度策略和优先级

线程的调度有三种策略：SCHED_OTHER、SCHED_RR和SCHED_FIFO。Policy用于指明使用哪种策略。下面我们简单的说明一下这三种调度策略。

## SCHED_OTHER（是Linux默认的分时调度策略）

它是默认的线程分时调度策略，所有的线程的优先级别都是0，线程的调度是通过分时来完成的。简单地说，如果系统使用这种调度策略，程序将无法设置线程的优先级。请注意，这种调度策略也是抢占式的，当高优先级的线程准备运行的时候，当前线程将被抢占并进入等待队列。这种调度策略仅仅决定线程在可运行线程队列中的具有相同优先级的线程的运行次序。

## SCHED_FIFO

它是一种实时的先进先出调用策略，且只能在超级用户下运行。这种调用策略仅仅被使用于优先级大于0的线程。它意味着，使用SCHED_FIFO的可运行线程将一直抢占使用SCHED_OTHER的运行线程J。此外SCHED_FIFO是一个非分时的简单调度策略，当一个线程变成可运行状态，它将被追加到对应优先级队列的尾部((POSIX 1003.1)。当所有高优先级的线程终止或者阻塞时，它将被运行。对于相同优先级别的线程，按照简单的先进先运行的规则运行。我们考虑一种很坏的情况，如果有若干相同优先级的线程等待执行，然而最早执行的线程无终止或者阻塞动作，那么其他线程是无法执行的，除非当前线程调用如pthread_yield之类的函数，所以在使用SCHED_FIFO的时候要小心处理相同级别线程的动作。

## SCHED_RR

鉴于SCHED_FIFO调度策略的一些缺点，SCHED_RR对SCHED_FIFO做出了一些增强功能。从实质上看，它还是SCHED_FIFO调用策略。它使用最大运行时间来限制当前进程的运行，当运行时间大于等于最大运行时间的时候，当前线程将被切换并放置于相同优先级队列的最后。这样做的好处是其他具有相同级别的线程能在“自私“线程下执行。

1.获得线程可以设置的最高和最低优先级，policy: 可以取三个值（SCHED_FIFO、SCHED_RR、SCHED_OTHER）

```
　int sched_get_priority_max(int policy);
　int sched_get_priority_min(int policy);
```
注意：对于 SCHED_OTHER 策略，sched_priority 只能为 0。对于 SCHED_FIFO，SCHED_RR 策略，sched_priority 从 1 到 99。

2.设置和获取优先级通过以下两个函数

```
int pthread_attr_setschedparam(pthread_attr_t *attr, const struct sched_param *param);
int pthread_attr_getschedparam(const pthread_attr_t *attr, struct sched_param *param);
param.sched_priority = 51; //设置优先级
了解sched_param结构体
struct sched_param
{
    int __sched_priority; // 所要设定的线程优先级
};
//param是struct sched_param类型的指针，它仅仅包含一个成员变sched_priority，指明所要设置的静态线程优先级。
```

3.改变策略(静态改变策略和设置优先级)

```
int pthread_attr_setschedpolicy(pthread_attr_t *attr, int policy);
int pthread_attr_getschedpolicy(pthread_attr_t *attr, int policy);
```

```
#include <unistd.h>
#include <pthread.h>
#include <sched.h>
#include "errors.h"

void *thread_routine(void *arg)
{
    int my_policy;
    struct sched_param my_param;

#if defined(_POSIX_THREAD_PRIORITY_SCHEDULING) && !defined(sun)
    pthread_getschedparam(pthread_self(), &my_policy, &my_param);
    printf("thread_routine running at %s/%d\n",
        (my_policy == SCHED_FIFO ? "FIFO"
            : (my_policy == SCHED_RR ? "RR"
            : (my_policy == SCHED_OTHER ? "OTHER"
            : "unknown"))),
        my_param.sched_priority);
#else
    printf("thread_routine running\n");
#endif
    return NULL;
}

int main(int argc, char *argv[])
{
    pthread_t thread_id;
    pthread_attr_t thread_attr;
    int thread_policy;
    struct sched_param thread_param;
    int status, rr_min_priority, rr_max_priority;

    pthread_attr_init(&thread_attr);

#if defined(_POSIX_THREAD_PRIORITY_SCHEDULING) && !defined(sun)
    pthread_attr_getschedpolicy(&thread_attr, &thread_policy);
   
    pthread_attr_getschedparam(&thread_attr, &thread_param);//获取优先级
    printf("Default policy is %s, priority is %d\n",
        (thread_policy == SCHED_FIFO ? "FIFO"
         : (thread_policy == SCHED_RR ? "RR"
            : (thread_policy == SCHED_OTHER ? "OTHER"
               : "unknown"))),
        thread_param.sched_priority);

    status = pthread_attr_setschedpolicy(&thread_attr, SCHED_RR);//改变策略
    if(status != 0)
        printf("Unable to set SCHED_RR policy.\n");
    else
    {
        rr_min_priority = sched_get_priority_min(SCHED_RR);//看当前策略下能设置优先级的范围
        if(rr_min_priority == -1)
            errno_abort("Get SCHED_RR min priority");
        rr_max_priority = sched_get_priority_max(SCHED_RR);
        if(rr_max_priority == -1)
            errno_abort("Get SCHED_RR max priority");
        
　　　　 thread_param.sched_priority = (rr_min_priority + rr_max_priority)/2;//静态改变优先级（线程运行之前）
        printf("SCHED_RR priority range is %d to %d: using %d\n",
            rr_min_priority, rr_max_priority, thread_param.sched_priority);
        pthread_attr_setschedparam(&thread_attr, &thread_param);
        printf("Creating thread at RR/%d\n", thread_param.sched_priority);
        pthread_attr_setinheritsched(&thread_attr, PTHREAD_EXPLICIT_SCHED); //无论何时，当你需要控制一个线程的调度策略或优先级时，必须将inheritsched属性设置为PTHREAD_EXPLICIT_SCHED。
    }
#else
    printf("Priority scheduling not supported\n");
#endif
    pthread_create(&thread_id, &thread_attr, thread_routine, NULL);
    pthread_join(thread_id, NULL);
    printf("Main exiting\n");
    return 0;
}
//运行结果
./sched_attr
Default policy is OTHER, priority is 0
SCHED_RR priority range is 1 to 99: using 50
Creating thread at RR/50
Main exiting
```

4.继承调度属性
我手动设置了调度策略或优先级时，必须显示的设置线程调度策略的inheritsched属性，因为pthread没有为inheritsched设置默认值。所以在改变了调度策略或优先级时必须总是设置该属性。

```
int pthread_attr_setinheritsched(pthread_attr_t *attr, int inheritsched);
int pthread_attr_getinheritsched(pthread_attr_t *attr, int *inheritsched);
```

第一个函数中inheritsched的取值为：PTHREAD_INHERIT_SCHED 或者 PTHREAD_EXPLICIT_SCHED。
前者为继承创建线程的调度策略和优先级，后者指定不继承调度策略和优先级，而是使用自己设置的调度策略和优先级。
无论何时，当你需要控制一个线程的调度策略或优先级时，必须将inheritsched属性设置为PTHREAD_EXPLICIT_SCHED。
总结：
1）调度策略和优先级是分开来描述的。前者使用预定义的SCHED_RR、SCHED_FIFO、SCHED_OTHER，后者是通过结果体struct sched_param给出的。
2）这些设置调度策略和优先级的函数操作的对象是线程的属性pthread_attr_t，而不是直接来操作线程的调度策略和优先级的。函数的第一个参数都是pthread_attr_t。
5.直接设置正在运行的线程的调度策略和优先级(动态设置线程的调度策略和优先级)
前面的那些函数只能通过线程的属性对象 pthread_attr_t 来设置线程的调度策略和优先级，不能够直接设置正在运行的线程的调度策略和优先级。下面的函数可以直接设置：
int pthread_setschedparam(pthread_t thread, int policy, const struct sched_param *param);
int pthread_getschedparam(pthread_t thread, int *policy, struct sched_param *param);
// 在成功完成之后返回零。其他任何返回值都表示出现了错误。如果出现以下任一情况，pthread_setschedparam() 函数将失败并返回相应的值--EINVAL所设置属性的值无效。ENOTSUP--尝试将该属性设置为不受支持的值。
失败条件：

int pthread_setschedparam：thread参数所指向的线程不存在
int pthread_getschedparam：1.参数policy或同参数policy关联的调度参数之一无效；2.数policy或调度参数之一的值不被支持；
3.调用线程没有适当的权限来设置指定线程的调度参数或策略；4.参数thread指向的线程不存在；5.实现不允许应用程序将参数改动为特定的值
注意：当pthread_setschedparam函数的参数 policy == SCHED_RR 或者 SCHED_FIFO 时，程序必须要在超级用户下运行

pthread_setschedparam 函数改变在运行线程的调度策略和优先级肯定就不用调用函数来设置inheritsched属性了：pthread_attr_setinheritsched(&thread_attr, PTHREAD_EXPLICIT_SCHED); 因为该函数设置的对象是pthread_attr_t 

```
/*
 * sched_thread.c
 * Demonstrate dynamic scheduling policy use.
 */
#include <unistd.h>
#include <pthread.h>
#include <sched.h>
#include "errors.h"

#define THREADS 5

typedef struct thread_tag {
    int index;
    pthread_t id;
} thread_t;

thread_t threads[THREADS];
int rr_min_priority;

void *thread_routine(void *arg)
{
    thread_t *self = (thread_t *)arg;
    struct sched_param my_param;
    int my_policy;
    int status;

    my_param.sched_priority = rr_min_priority + self->index;
    if(pthread_setschedparam(self->id, SCHED_RR, &my_param) != 0)//线程运行时设置策略和优先值
        printf("pthread_setschedparam failed\n");
    pthread_getschedparam(self->id, &my_policy, &my_param);
    printf("thread_routine %d running at %s/%d\n",
        self->index,
        (my_policy == SCHED_FIFO ? "FIFO"
            : (my_policy == SCHED_RR ? "RR"
            : (my_policy == SCHED_OTHER ? "OTHER"
            : "unknown"))),
        my_param.sched_priority);

    return NULL;
}

int main(int argc, char *argv[])
{
    int count;

    rr_min_priority = sched_get_priority_min(SCHED_RR);
    if(rr_min_priority == -1){
        errno_abort("Get SCHED_RR min priority");
    }
    for(count = 0; count < THREADS; count++){
        threads[count].index = count;
        pthread_create(&threads[count].id, NULL,
                 thread_routine, (void *)&threads[count]);
    }
    for(count = 0; count < THREADS; count++){
        pthread_join(threads[count].id, NULL);
    }
    printf("Main exiting\n");

    return 0;
}
```

```
//运行结果
gcc -Wall -lpthread -o sched_thread sched_thread.c
./sched_thread
pthread_setschedparam failed
pthread_setschedparam failed
pthread_setschedparam failed
thread_routine 1 running at OTHER/0
thread_routine 3 running at OTHER/0
pthread_setschedparam failed
pthread_setschedparam failed
thread_routine 2 running at OTHER/0
thread_routine 4 running at OTHER/0
thread_routine 0 running at OTHER/0
Main exiting
//以上失败，一下运行用：sudo ./sched_thread (输入密码)时，函数 pthread_setschedparam(self->id, SCHED_RR, &my_param) 调用成功。将SCHED_RR换成SCHED_FIFO，结果也是一样的。
$ sudo ./sched_thread
[sudo] password for digdeep: 
thread_routine 1 running at RR/2
thread_routine 3 running at RR/4
thread_routine 2 running at RR/3
thread_routine 0 running at RR/1
thread_routine 4 running at RR/5
Main exiting
```
它有很多改进的地方，比如可以使用虚优先级（在程序中加入虚实影射表）等
