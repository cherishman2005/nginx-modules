c++中关于std::thread的join的思考

std::thread是c++11新引入的线程标准库，通过其可以方便的编写与平台无关的多线程程序，虽然对比针对平台来定制化多线程库会使性能达到最大，但是会丧失了可移植性，这样对比其他的高级语言，可谓是一个不足。终于在c++11承认多线程的标准，可谓可喜可贺！！！

在使用std::thread的时候，对创建的线程有两种操作：等待/分离，也就是join/detach操作。join()操作是在std::thread t(func)后“某个”合适的地方调用，其作用是回收对应创建的线程的资源，避免造成资源的泄露。detach()操作是在std::thread t(func)后马上调用，用于把被创建的线程与做创建动作的线程分离，分离的线程变为后台线程,其后，创建的线程的“死活”就与其做创建动作的线程无关，它的资源会被init进程回收。

在这里主要对join做深入的理解。

由于join是等待被创建线程的结束，并回收它的资源。因此，join的调用位置就比较关键。比如，以下的调用位置都是错误的。

例子一：
```
void test()
{
}

bool do_other_things()
{
}

int main()
{
    std::thread t(test);
    int ret = do_other_things();
    if(ret == ERROR) {
        return -1;
    }

    t.join();
    return 0;
}
```
很明显，如果do_other_things()函数调用返ERROR, 那么就会直接退出main函数，此时join就不会被调用，所以线程t的资源没有被回收，造成了资源泄露。

例子二：
```
void test()
{
}

bool do_other_things()
{
}

int main()
{
    std::thread t(test);

    try {
        do_other_things();
    }
    catch(...) {
        throw;
    }
    t.join();
    return 0;
}
```
这个例子和例子一差不多，如果调用do_other_things()函数抛出异常，那么就会直接终止程序，join也不会被调用，造成了资源没被回收。

那么直接在异常捕捉catch代码块里调用join就ok啦。
例子三：
```
void test()
{
}

bool do_other_things()
{
}

int main()
{
    std::thread t(test);

    try {
        do_other_things();
    }
    catch(...) {
        t.join();
        throw;
    }
    t.join();
    return 0;
}
```
是不是很多人这样操作？这样做不是万无一失的， try/catch块只能够捕捉轻量级的异常错误，在这里如果在调用do_other_things()时发生严重的异常错误，那么catch不会被触发捕捉异常，同时造成程序直接从函数调用栈回溯返回，也不会调用到join，也会造成线程资源没被回收，资源泄露。

所以在这里有一个方法是使用创建局部对象，利用函数调用栈的特性，确保对象被销毁时触发析构函数的方法来确保在主线程结束前调用join()，等待回收创建的线程的资源。
```
class mythread {
private:
    std::thread &m_t;

public:
    explicit mythread(std::thread &t):m_t(t){}
    ~mythread() {
        if(t.joinable()) {
            t.join()
        }
    }

    mythread(mythread const&) = delete;
    mythread& operate=(mythread const&) = delete;
}

void test()
{
}

bool do_other_things()
{
}

int main()
{
    std::thread t(test);
    mythread q(t);

    if(do_other_things()) {
        return -1;
    }

    return 0;
}
```
在上面的例子中，无论在调用do_other_things()是发生错误，造成return main函数，还是产生异常，由于函数调用栈的关系，总会回溯的调用局部对象q的析构函数，同时在q的析构函数里面先判断j.joinable()是因为join操作对于同一个线程只能调用一次，不然会出现错误的。这样，就可以确保线程一定会在主函数结束前被等待回收了。
