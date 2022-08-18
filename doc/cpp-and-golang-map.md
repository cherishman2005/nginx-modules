# C++和golang数据类型区别

## C++ STL

STL中的容器直接赋值是安全的，即是深拷贝。

```
vecotr<string> vec1; vec1.push_back("helloWorld");    vecotr<string> vec2;
```
vec1 = vec2,赋值之后，vec1和vec2之间就没有关联了  。


## golang map

golang map是指针类型，赋值是浅拷贝。
