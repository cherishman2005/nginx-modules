#C++ explicit的作用

explicit作用:

在C++中，explicit关键字用来修饰类的构造函数，被修饰的构造函数的类，不能发生相应的隐式类型转换，只能以显示的方式进行类型转换。

explicit使用注意事项:

    *

      explicit 关键字只能用于类内部的构造函数声明上。

    *

      explicit 关键字作用于单个参数的构造函数。

    * 在C++中，explicit关键字用来修饰类的构造函数，被修饰的构造函数的类，不能发生相应的隐式类型转换

例子：

未加explicit时的隐式类型转换

```
class Circle 
{ 
public: 
    Circle(double r) : R(r) {} 
    Circle(int x, int y = 0) : X(x), Y(y) {} 
    Circle(const Circle& c) : R(c.R), X(c.X), Y(c.Y) {} 
private: 
    double R; 
    int    X; 
    int    Y; 
}; 
 
int _tmain(int argc, _TCHAR* argv[]) 
{ 
//发生隐式类型转换 
//编译器会将它变成如下代码 
//tmp = Circle(1.23) 
//Circle A(tmp); 
//tmp.~Circle(); 
    Circle A = 1.23;  
//注意是int型的，调用的是Circle(int x, int y = 0) 
//它虽然有2个参数，但后一个有默认值，任然能发生隐式转换 
    Circle B = 123; 
//这个算隐式调用了拷贝构造函数 
    Circle C = A; 
     
    return 0; 
} 
```

加了explicit关键字后，可防止以上隐式类型转换发生

```
 class Circle 
 { 
 public: 
     explicit Circle(double r) : R(r) {} 
     explicit Circle(int x, int y = 0) : X(x), Y(y) {} 
     explicit Circle(const Circle& c) : R(c.R), X(c.X), Y(c.Y) {} 
 private: 
     double R; 
     int    X; 
     int    Y; 
 }; 
  
 int _tmain(int argc, _TCHAR* argv[]) 
 { 
 //一下3句，都会报错 
     //Circle A = 1.23;  
     //Circle B = 123; 
     //Circle C = A; 
      
 //只能用显示的方式调用了 
 //未给拷贝构造函数加explicit之前可以这样 
          Circle A = Circle(1.23); 
         Circle B = Circle(123); 
         Circle C = A; 
  
 //给拷贝构造函数加了explicit后只能这样了 
          Circle A(1.23); 
         Circle B(123); 
         Circle C(A); 
     return 0; 
 } 
```
