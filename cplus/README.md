# cpp

## C++四种类型转换运算符

| 关键字      |   说明                 |
| --------- | --------------------------- |
| static_cast | 用于良性转换，一般不会导致意外发生，风险很低。      | 
| const_cast | 用于 const 与非 const、volatile 与非 volatile 之间的转换。      | 
| reinterpret_cast | 高度危险的转换，这种转换仅仅是对二进制位的重新解释，不会借助已有的转换规则对数据进行调整，但是可以实现最灵活的 C++ 类型转换。    | 
| dynamic_cast | 借助 RTTI，用于类型安全的向下转型（Downcasting）。      | 

这四个关键字的语法格式都是一样的，具体为：
```
xxx_cast<newType>(data)
```

# 参考链接

- [https://github.com/wuye9036/CppTemplateTutorial](https://github.com/wuye9036/CppTemplateTutorial)
- [关于lower_bound( )和upper_bound( )的常见用法](https://blog.csdn.net/qq_40160605/article/details/80150252)
- [C++11 std::unique_lock与std::lock_guard区别及多线程应用实例](https://www.cnblogs.com/fnlingnzb-learner/p/9542183.html)
- [C++四种类型转换运算符：static_cast、dynamic_cast、const_cast和reinterpret_cast](http://c.biancheng.net/cpp/biancheng/view/3297.html)
