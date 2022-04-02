# thread

## pthread_kill

```
int pthread_kill(pthread_t thread, int sig);
```

所以，如果int sig的参数不是0，那一定要清楚到底要干什么，而且一定要实现线程的信号处理函数，否则，就会影响整个进程。
OK，如果int sig是0呢，这是一个保留信号，一个作用是用来判断线程是不是还活着。

```
int kill_rc = pthread_kill(thread_id,0);

if(kill_rc == ESRCH)
printf("the specified thread did not exists or already quit\n");
else if(kill_rc == EINVAL)
printf("signal is invalid\n");
else
printf("the specified thread is alive\n");
```
上述的代码就可以判断线程是不是还活着了。

# 参考链接

- [对于pthread_kill的一些误解](https://blog.csdn.net/qu1993/article/details/105580924?spm=1001.2101.3001.6650.3&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-3.pc_relevant_antiscanv2&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-3.pc_relevant_antiscanv2&utm_relevant_index=6)

- [pthread_kill()和pthread_cancel()的用法](https://blog.csdn.net/jongden/article/details/25707477?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_title~default-0.pc_relevant_paycolumn_v3&spm=1001.2101.3001.4242.1&utm_relevant_index=3)
