# C++禁止使用拷贝构造函数和赋值运算符方法

1.将拷贝构造函数和赋值运算符声明为私有，并不予实现

```
class Uncopyable  
{  
private:  
    Uncopyable(const Uncopyable &); // 阻止copying  
    Uncopyable &operator=(const Uncopyable &);  
}; 
```

2.使用delete

```
class Uncopyable  
{  

    Uncopyable(const Uncopyable &) =delete; // 阻止copying  
    Uncopyable &operator=(const Uncopyable &)=delete;  
};  
```

# 禁用拷贝（复制）构造函数

关于C++的拷贝构造函数，很多的建议是直接禁用。为什么大家会这么建议呢？没有拷贝构 造函数会有什么限制呢？如何禁用拷贝构造呢？这篇文章对这些问题做一个简单的总结。

这里讨论的问题以拷贝构造函数为例子，但是通常赋值操作符是通过拷贝构造函数来实现 的（ copy-and-swap 技术，详见《Exceptional C++》一书），所以这里讨论也适用于赋 值操作符，通常来说禁用拷贝构造函数的同时也会禁用赋值操作符。

为什么禁用拷贝构造函数
关于拷贝构造函数的禁用原因，我目前了解的主要是两个原因。第一是浅拷贝问题，第二 个则是基类拷贝问题。

浅拷贝问题
编译器默认生成的构造函数，是memberwise拷贝^1，也就是逐个拷贝成员变量，对于 下面这个类的定义^2：
```
class Widget {
 public:
    Widget(const std::string &name) : name_(name), buf_(new char[10]) {}
    ~Widget() { delete buf_; }

 private:
    std::string name_;
    char *buf_;
};
```
默认生成的拷贝构造函数，会直接拷贝buf_的值，导致两个Widget对象指向同一个缓 冲区，这会导致析构的时候两次删除同一片区域的问题（这个问题又叫双杀问题）。

解决这个问题的方式有很多：

自己编写拷贝构造函数，然后在拷贝构造函数中创建新的buf_，不过拷贝构造函数的 编写需要考虑异常安全的问题，所以编写起来有一定的难度。

使用 shared_ptr 这样的智能指针，让所有的 Widget 对象共享一片 buf_，并 让 shared_ptr 的引用计数机制帮你智能的处理删除问题。

禁用拷贝构造函数和赋值操作符。如果你根本没有打算让Widget支持拷贝，你完全可 以直接禁用这两操作，这样一来，前面提到的这些问题就都不是问题了。

基类拷贝构造问题
如果我们不去自己编写拷贝构造函数，编译器默认生成的版本会自动调用基类的拷贝构造 函数完成基类的拷贝：
```
class Base {
 public:
    Base() { cout << "Base Default Constructor" << endl; }
    Base(const Base &) { cout << "Base Copy Constructor" << endl; }
};

class Drived : public Base {
 public:
    Drived() { cout << "Drived Default Constructor" << endl; }
};

int main(void) {
    Drived d1;
    Drived d2(d1);
}
```
上面这段代码的输出如下：
```
Base Default Constructor
Drived Default Constructor

Base Copy Constructor  // 自动调用了基类的拷贝构造函数
```
但是如果我们出于某种原因编写了，自己编写了拷贝构造函数（比如因为上文中提到的浅 拷贝问题），编译器不会帮我们安插基类的拷贝构造函数，它只会在必要的时候帮我们安 插基类的默认构造函数：
```
class Base {
 public:
    Base() { cout << "Base Default Constructor" << endl; }
    Base(const Base &) { cout << "Base Copy Constructor" << endl; }
};

class Drived : public Base {
 public:
    Drived() { cout << "Drived Default Constructor" << endl; }
    Drived(const Drived& d) {
    	cout << "Drived Copy Constructor" << endl;
    }
};

int main(void) {
    Drived d1;
    Drived d2(d1);
}
```
上面这段代码的输出如下：
```
Base Default Constructor
Drived Default Constructor

Base Default Constructor // 调用了基类的默认构造函数
Drived Copy Constructor
```
这当然不是我们想要看到的结果，为了能够得到正确的结果，我们需要自己手动调用基类 的对应版本拷贝基类对象。
```
Drived(const Drived& d) : Base(d) {
    cout << "Drived Copy Constructor" << endl;
}
```
这本来不是什么问题，只不过有些人编写拷贝构造函数的时候会忘记这一点，所以导致基 类子对象没有正常复制，造成很难察觉的BUG。所以为了一劳永逸的解决这些蛋疼的问题， 干脆就直接禁用拷贝构造和赋值操作符。

没有拷贝构造的限制
在C++11之前对象必须有正常的拷贝语义才能放入容器中，禁用拷贝构造的对象无法直接放 入容器中，当然你可以使用指针来规避这一点，但是你又落入了自己管理指针的困境之中 （或许使用智能指针可以缓解这一问题）。

C++11中存在移动语义，你可以通过移动而不是拷贝把数据放入容器中。

拷贝构造函数的另一个应用在于设计模式中的原型模式，在C++中没有拷贝构造函数，这 个模式实现可能比较困难。

如何禁用拷贝构造
如果你的编译器支持 C++11，直接使用 delete

否则你可以把拷贝构造函数和赋值操作符声明成private同时不提供实现。

你可以通过一个基类来封装第二步，因为默认生成的拷贝构造函数会自动调用基类的拷 贝构造函数，如果基类的拷贝构造函数是 private，那么它无法访问，也就无法正常 生成拷贝构造函数。
```
class NonCopyable {
protected:
    ~NonCopyable() {}  // 关于为什么声明成为 protected，参考
    		       // 《Exceptional C++ Style》
private:
    NonCopyable(const NonCopyable&);
}

class Widget : private NonCopyable { // 关于为什么使用 private 继承
				     // 参考《Effective C++》第三版
}

Widget widget(Widget()); // 错误
```
上不会生成memberwise的拷贝构造函数，详细内容可以参考《深度探索C++对象模型》一 书
