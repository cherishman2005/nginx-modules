# Go语言中[]byte和string类型相互转换


我们在使用Go语言时，经常涉及到[]byte和string两种类型间的转换。本篇文章将讨论转换时的开销，Go编译器在一些特定场景下对转换做的优化，以及在高性能场景下，我们自己如何做相应的优化。

[]byte其实就是byte类型的切片，对应的底层结构体定义如下（在runtime/slice.go文件中）
```
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```
string对应的底层结构体定义如下（在runtime/string.go文件中）
```
type stringStruct struct {
	str unsafe.Pointer
	len int
}
```
可以看到它们内部都有一个指针类型(array或str)，指向真实数据。另外还有一个len字段，标识数据的长度。

slice多了一个cap字段，表示容量大小。当要往slice尾部追加数据而空余容量又不够时，会重新分配更大的内存块，将当前内存块的内容拷贝至新内存块，再在新内存块做追加。

slice变量间做赋值操作时，只是修改指针指向，不会拷贝真实数据。string变量间赋值也是同样的道理。


`[]byte和string相互转换，就需要重新申请内存并拷贝内存。因为Go语义中，slice的内容是可变的（mutable），而string是不可变的（immutable）。如果他们底部指向同一块数据，那么由于slice可对数据做修改，string就做不到immutable了。`

# 参考链接

- [Go语言中[]byte和string类型相互转换时的性能分析和优化](https://www.pengrl.com/p/31544/)
