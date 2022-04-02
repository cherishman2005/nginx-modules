# c++中临时变量不能作为非const的引用参数

示例代码：
```
#include <iostream>
using namespace std;

void f(int &a)

 {
 cout << "f(" << a  << ") is being called" << endl;
}

void g(const int &a)

{
 cout << "g(" << a << ") is being called" << endl;
}

 int main()

 {
 int a = 3, b = 4;
 f(a + b);  //编译错误，把临时变量作为非const的引用参数传递了
 g(a + b);  //OK，把临时变量作为const&传递是允许的
}
```

上面的两个调用之前，a+b的值会存在一个临时变量中，当把这个临时变量传给f时，由于f的声明中，参数是int&，不是常量引用，所以产生以下编译错误：
```
const_ref.cpp: In function `int main()':
const_ref.cpp:14: error: invalid initialization of non-const reference of type '
   int&' from a temporary of type 'int'
const_ref.cpp:4: error: in passing argument 1 of `void f(int&)' 
```
而在g(a+b)中，由于g定义的参数是const int&，编译通过。   问题是为什么临时变量作为引用参数传递时，必须是常量引用呢？很多人对此的解释是临时变量是常量，不允许赋值，改动，所以当作为非常量引用传递时，编译器就会报错。这个解释在关于理解临时变量不能作为非const引用参数这个问题上是可以的，但不够准确。事实上，临时变量是可以被作为左值(LValue)并被赋值的，请看下面的代码：

```
#include   <iostream> 
using namespace std;

class CComplex {   
friend CComplex operator+(const CComplex &cp1, const CComplex &cp2);
friend ostream& operator<<(ostream &os, const CComplex &cp);
private: 
 int x; 
public: 
 CComplex(){}
  
  CComplex(int x1) { 
  x = x1; 
 }
};
 
CComplex operator+(const CComplex &cp1, const CComplex &cp2)

{ 
 CComplex cp3; 
 cp3.x = cp1.x + cp2.x; 
 return cp3; 
} ostream& operator<<(ostream &os, const CComplex &cp)

{
 os << cp.x;
 return os;
}

int main()

{ 
 CComplex a(2), b(3), c(4); 
 cout << (a + b) << endl;
 cout << ((a + b) = c) << endl;   //临时对象作为左值
 return 0; 
}
```
上面的程序编译通过，而且运行结果是：
```
5

4
```
临时变量确实被赋值，而且成功了。
所以，临时变量不能作为非const引用参数，不是因为他是常量，而是因为c++编译器的一个关于语义的限制。如果一个参数是以非const引用传入，c++编译器就有理由认为程序员会在函数中修改这个值，并且这个被修改的引用在函数返回后要发挥作用。但如果你把一个临时变量当作非const引用参数传进来，由于临时变量的特殊性，程序员并不能操作临时变量，而且临时变量随时可能被释放掉，所以，一般说来，修改一个临时变量是毫无意义的，据此，c++编译器加入了临时变量不能作为非const引用的这个语义限制，意在限制这个非常规用法的潜在错误。
还不明白？OK，我们说直白一点，如果你把临时变量作为非const引用参数传递，一方面，在函数申明中，使用非常量型的引用告诉编译器你需要得到函数对某个对象的修改结果，可是你自己又不给变量起名字，直接丢弃了函数的修改结果，编译器只能说：“大哥，你这是干啥呢，告诉我把结果给你，等我把结果给你了，你又直接给扔了，你这不是在玩我呢吗？”所以编译器一怒之下就不让过了。这下大家明白了吧？
