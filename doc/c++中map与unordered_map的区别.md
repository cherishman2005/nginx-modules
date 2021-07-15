#c++中map与unordered_map的区别

## 头文件
* map: #include < map >
* unordered_map: #include < unordered_map >

## 内部实现机理
* map： map内部实现了一个红黑树，该结构具有自动排序的功能，因此map内部的所有元素都是有序的，红黑树的每一个节点都代表着map的一个元素，因此，对于map进行的查找，删除，添加等一系列的操作都相当于是对红黑树进行这样的操作，故红黑树的效率决定了map的效率。
* unordered_map: unordered_map内部实现了一个哈希表，因此其元素的排列顺序是杂乱的，无序的

## 优缺点以及适用处
* map
  * 优点：
    * 有序性，这是map结构最大的优点，其元素的有序性在很多应用中都会简化很多的操作
    * 红黑树，内部实现一个红黑书使得map的很多操作在lgn的时间复杂度下就可以实现，因此效率非常的高
  * 缺点：
    * 空间占用率高，因为map内部实现了红黑树，虽然提高了运行效率，但是因为每一个节点都需要额外保存父节点，孩子节点以及红/黑性质，使得每一个节点都占用大量的空间
  * 适用处，对于那些`有顺序要求的问题`，用map会更高效一些

* unordered_map
  * 优点：
    * 因为内部实现了哈希表，因此其查找速度非常的快
  * 缺点：
    * 哈希表的建立比较耗费时间
  * 适用处，对于`查找问题`，unordered_map会更加高效一些，因此遇到查找问题，常会考虑一下用unordered_map

# note:
* 对于unordered_map或者unordered_set容器，其遍历顺序与创建该容器时输入元素的顺序是不一定一致的，遍历是按照哈希表从前往后依次遍历的


## unordered_map hash问题解决


### 问题
G++使用unordered_map时候，编译报错：invalid use of incomplete type ‘struct std::hash<，。。。，放在G++6.5交叉编译环境是OK的，但是放在ubuntu14.04报错。

### 解决&代码

既然G++早期版本不能自动生成枚举类型的hash模板类，那么手动添加template<> struct std::hash<...。

添加如下代码 #if ....#endif区域代码，即可解决问题。
```
#include <unordered_map>
#include <utility>
#include <cstdint>
#include <iostream>
#include <functional>
 
namespace test{
  enum COLOR{ WHITE, BLAC };
}
 
#if 0  // 如果没有这里，G++4.8.4和G++5.4.0会报错
namespace std {
template<>
struct hash<test::COLOR> {
   typedef test::COLOR argument_type;
   typedef size_t result_type;
 
   result_type operator () (const argument_type& x) const {
      using type = typename std::underlying_type<argument_type>::type;
      return std::hash<type>()(static_cast<type>(x));
   }
};
}
#endif
 
namespace test{
class mytest{
 public:
  std::unordered_map<COLOR, int> id_map_;
};
}
 
int main(){
    test::mytest t;
    return 0;
}
```
### 结论

发现是G++的问题，ubuntu14.04默认是g++4.8.4,ubuntu16.04是g++5.4。

参考：https://stackoverflow.com/questions/48294401/error-invalid-use-of-incomplete-type-struct-stdhash

# 参考连接

- [自定义hash键C++](https://www.cnblogs.com/Shinered/p/9193329.html)

- [解决C++ unordered_map“invalid use of incomplete type ‘struct std::hash“ 问题](https://blog.csdn.net/li459461891/article/details/104910925)

- [unordered_set/map自定义哈希函数](https://blog.kedixa.top/2017/cpp-user-defined-hash/)
