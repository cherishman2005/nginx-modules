# FAQ

## http Content-Length

nginx http请求头的Content-Length = -1时判定客户端不合法，nginx返回400错误码。

## 多机器集群

nodejs-cluster 适用于在单台机器上，如果应用的流量巨大，多机器是必然的。这时，反向代理就派上用场了，可以用 node 来写反向代理的服务（比如用 http-proxy），好处是可以保持工程师技术栈的统一，不过生产环境，更多的还是nginx。

## warning: ISO C++ forbids converting a string constant to 'char*' [-Wwrite-strings]
在C++11中有明确规定
```C
char* p = "abc"; // valid in C, invalid in C++
```

如果你进行了这样的赋值，那么编译器就会跳出诸如标题的警告。但是如果你改成下面这样就会通过warning
```C
char* p = (char*)"abc"; //OK
```

或者这样：
```C
char const *p="abc";//OK
```

这到底是怎么一回事呢？事实上，我们在学习c或者c++的时候都知道，如果在赋值操作的时候，等号两边的变量类型不一样，那么编译器会进行一种叫做 implicit conversion 的操作来使得变量可以被赋值。
 
在我们上面的表达式中就存在这样的一个问题，等号右边的"abc"是一个不变常量，在c++中叫做string literal，type是const char *，而p则是一个char指针。如果强行赋值会发生什么呢？没错，就是将右边的常量强制类型转换成一个指针，结果就是我们在修改一个const常量。编译运行的结果会因编译器和操作系统共同决定，有的编译器会通过，有的会抛异常，就算过了也可能因为操作系统的敏感性而被杀掉。
 
像这种直接将string literal 赋值给指针的操作被开发者们认为是deprecated，只不过由于以前很多代码都有这种习惯，为了兼容，就保留下来了。

