# golang进阶笔记

## 如何判断 2 个字符串切片（slice) 是相等的？
reflect.DeepEqual() ， 但反射非常影响性能。



## 空 struct{} 的用途

用map模拟一个set，那么就要把值置为struct{}，struct{}本身不占任何空间，可以避免任何多余的内存分配。

* C++和golang的map完全是不相同的数据结构；
  * C++ map/set采用红黑树，数据是有序； unordered_map/unordered_set才与golang的 map对应；
