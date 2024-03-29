
```
error: templates may not be ‘virtual’
```
# 模板函数不能是虚函数
 
原因如下：

首先呢，模板函数并不是函数，他需要特定的类型去实例化成为函数。你定义一个函数模板，是不生成任何函数的，只有当你用代码去调用它时，才会根据你的类型去实例化成为特定的函数。

而virtual函数是要写入虚函数表的，是必须要存在的。你可能会想到纯虚函数，纯虚函数只是表明这个函数还未实现，但是已经在父类的虚表里存在了。

因此，模板函数是不能声明为virtual的。
 
# 静态函数不能是虚函数

设计方面的原因：
虚函数是为了实现运行期函数和对象（类的实例）的动态绑定，通过对象的指针或引用访问被指向的对象，只要有继承关系，被访问的对象的实际类型可以和指针或引用指向的类型不同。

如果没有对象，那么这种多态就没有意义，因为根本不存在需要在运行期确定对象类型的必要。

所以只从属于类而不和具体对象相关的静态成员函数作为虚函数是没有意义的，因此语言禁止这么做。

# 参考链接

- [https://stackoverflow.com/questions/4961909/templates-may-not-be-virtual](https://stackoverflow.com/questions/4961909/templates-may-not-be-virtual)
