# C++强制类型转换运算符

C++强制类型转换运算符（static_cast、reinterpret_cast、const_cast和dynamic_cast）

```
强制类型转换运算符 <要转换到的类型> (待转换的表达式)
```

## static_cast

```
double d = static_cast <double> (3*5);  //将 3*5 的值转换成实数
```

## reinterpret_cast

reinterpret_cast 用于进行各种不同类型的指针之间、不同类型的引用之间以及指针和能容纳指针的整数类型之间的转换。转换时，执行的是逐个比特复制的操作。

# 参考

- [C++强制类型转换运算符（static_cast、reinterpret_cast、const_cast和dynamic_cast）](http://c.biancheng.net/view/410.html)