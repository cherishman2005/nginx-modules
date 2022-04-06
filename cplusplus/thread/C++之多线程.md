# C++之多线程（C++11 thread.h文件实现多线程）

http://www.cnblogs.com/haippy/p/3235560.html

http://www.cnblogs.com/lidabo/p/3908705.html

与 C++11 多线程相关的头文件
C++11 新标准中引入了四个头文件来支持多线程编程，他们分别是<atomic> ,<thread>,<mutex>,<condition_variable>和<future>。
<atomic>：该头文主要声明了两个类, std::atomic 和 std::atomic_flag，另外还声明了一套 C 风格的原子类型和与 C 兼容的原子操作的函数。
<thread>：该头文件主要声明了 std::thread 类，另外 std::this_thread 命名空间也在该头文件中。
<mutex>：该头文件主要声明了与互斥量(mutex)相关的类，包括 std::mutex 系列类，std::lock_guard, std::unique_lock, 以及其他的类型和函数。
<condition_variable>：该头文件主要声明了与条件变量相关的类，包括 std::condition_variable 和 std::condition_variable_any。
<future>：该头文件主要声明了 std::promise, std::package_task 两个 Provider 类，以及 std::future 和 std::shared_future 两个 Future 类，另外还有一些与之相关的类型和函数，std::async() 函数就声明在此头文件中。
demo1：thread "Hello world"
demo2：一次启动多个线程
我们通常希望一次启动多个线程，来并行工作。为此，我们可以创建线程组，而不是在先前的举例中那样创建一条线程。下面的例子中，主函数创建十条为一组的线程，并且等待这些线程完成他们的任务

记住，主函数也是一条线程，通常叫做主线程，所以上面的代码实际上有11条线程在运行。在启动这些线程组之后，线程组和主函数进行协同（join）之前，允许我们在主线程中做些其他的事情。
demo3：在线程中使用带有形参的函数
运行结果：
能看到上面的结果中，程序一旦创建一条线程，其运行存在先后秩序不确定的现象。程序员的任务就是要确保这组线程在访问公共数据时不要出现阻塞。最后几行，所显示的错乱输出，表明8号线程启动的时候，4号线程还没有完成在stdout上的写操作。事实上假定在你自己的机器上运行上面的代码，将会获得全然不同的结果，甚至是会输出些混乱的字符。原因在于，程序内的11条线程都在竞争性地使用stdout这个公共资源（案：Race Conditions）。
要避免上面的问题，可以在代码中使用拦截器（barriers），如std:mutex，以同步（synchronize）的方式来使得一群线程访问公共资源，或者，如果可行的话，为线程们预留下私用的数据结构，避免使用公共资源。我们在以后的教学中，还会讲到线程同步问题，包括使用原子操作类型（atomic types）和互斥体（mutex）。
更多内容请参考：

C++11 并发指南一(C++11 多线程初探)
http://www.cnblogs.com/haippy/p/3235560.html
C++11 并发指南二(std::thread 详解)
http://www.cnblogs.com/haippy/p/3236136.html
C++11 并发指南三(std::mutex 详解)
http://www.cnblogs.com/haippy/p/3237213.html
C++11 并发指南三(Lock 详解)
http://www.cnblogs.com/haippy/p/3346477.html
C++11 并发指南四(<future> 详解一 std::promise 介绍)
http://www.cnblogs.com/haippy/p/3239248.html
C++11 并发指南四(<future> 详解二 std::packaged_task 介绍)
http://www.cnblogs.com/haippy/p/3279565.html
C++11 并发指南四(<future> 详解三 std::future & std::shared_future)
http://www.cnblogs.com/haippy/p/3280643.html
C++11 并发指南五(std::condition_variable 详解)
http://www.cnblogs.com/haippy/p/3252041.html
C++11 并发指南六(atomic 类型详解一 atomic_flag 介绍)
http://www.cnblogs.com/haippy/p/3252056.html
C++11 并发指南六( <atomic> 类型详解二 std::atomic )
http://www.cnblogs.com/haippy/p/3301408.html
C++11 并发指南六(atomic 类型详解三 std::atomic (续))
http://www.cnblogs.com/haippy/p/3304556.html
C++11 并发指南六(atomic 类型详解四 C 风格原子操作介绍)
http://www.cnblogs.com/haippy/p/3306625.html
C++11 并发指南七(C++11 内存模型一：介绍)
http://www.cnblogs.com/haippy/p/3412858.html
C++11 并发指南九(综合运用: C++11 多线程下生产者消费者模型详解)
http://www.cnblogs.com/haippy/p/3252092.html
