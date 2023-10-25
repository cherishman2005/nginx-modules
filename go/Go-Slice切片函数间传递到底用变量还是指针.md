# Go Slice切片函数间传递到底用变量还是指针？Go语言数据类型，值传递和引用变量类型的思考


## 1. 切片

    Go语言的切片，slice是动态数组，底层实现和java的arrayList一样，也是go 开发中经常用到的数据结构，但是与java语言不同的是，go语言有指针，java没有指针，因此使用上还是有一定的区别。

## 2.切片在函数间传递使用变量的情况

只有在切片的底层数组长度不变的情况下时，可以使用变量，但是这似乎又不符合切片的功能和使用场景，代码如下：
```
package main
 
import (
	"fmt"
)
 
func funUseVaribale(paramSlice []int) {
	paramSlice[2] = 9
	for index := 0; index < 6; index++ {
		paramSlice = append(paramSlice, index)
	}
	fmt.Printf("modifyE paramSlice =%v\n", paramSlice)
}
 
func main() {
    slice := make([]int, 5)
	slice[0] = 2
	slice[1] = 3
	fmt.Println(slice)
	funUseVaribale(slice)
	fmt.Println(slice)
 
}
```

第三行的打印证明了funUseVaribale函数对slice的改变确实影响到了slice，也就是说slice传递到函数中的时候使用的是该变量的地址，所以slice是值传递的引用形式(不知道这么说贴切不？)

## 3.切片在函数间传递使用指针的情况

只有在切片的底层数组长度会发生变化且这个变化在外部参数也要感知到的情况下时，代码如下：

```
package main
 
import (
	"fmt"
)
 
func funUsePointer(paramSlice *[]int) {
	paramSliceV := *paramSlice
	paramSliceV[2] = 9
	for index := 0; index < 6; index++ {
		*paramSlice = append(*paramSlice, index)
	}
	fmt.Printf("funUsePointer paramSlice =%v\n", paramSlice)
}
 
func main() {
	
	slice := make([]int, 5)
	slice[0] = 2
	slice[1] = 3
	fmt.Println(slice)
	funUsePointer(&slice)
	fmt.Println(slice)
}
```
 

可以看到funUsePointer paramSlice打印的结果和第三行打印的结果一致，也就是说函数 funUsePointer对变量paramSlice的改变也影响到了slice(其实呢，他们指向了同一个内存地址，这个地址就是切片扩容后的底层数组地址)

## 4.值传递的go和引用类型

go语言的数据类型有三种：内置类型，引用类型，结构类型：

内置类型是由语言提供的一组类型，分别是数值类型、字符串类型和布尔类型。这些类型本质上是原始的类型。因此，当对这些值进行增加或者删除的时候，会创建一个新值。基于这个结论，当把这些类型的值传递给方法或者函数时，应该传递一个对应值的副本(见<<go 语言实战>>)

引用类型：切片，映射，通道，接口和函数都是引用类型，当声明一个引用类型的变量时，创建的变量被称作标头值(header)，标头值是包含一个指向底层数据结构的指针，每个引用类型还包含一组独特的字段，用于管理底层数据结构。因为标头值是为复制而设计的，所以永远不需要共享一个引用类型的值。标头值里包含一个指针，因此通过复制来传递一个引用类型的值的副本，本质上就是在共享底层数据结构。

结构类型可以用来描述一组数据值，这组值的本质即可以是原始的，也可以是非原始的。如果决定在某些东西需要删除或者添加某个结构类型的值时该结构类型的值不应该被更改，那么需要遵守之前提到的内置类型和引用类型的规范。大多数情况下，结构类型的本质并不是原始的，而是非原始的。这种情况下，对这个类型的值做增加或者删除的操作应该更改值本身。当需要修改值本身时，在程序中其他地方，需要使用指针来共享这个值。

## 5.切片函数间传递的方式的分析

基于以上知识，我们可以得出，切片在函数间传递的时候，如果参数是非指针的时候，其实传递的是标头值，也就是切片底层的数组指针，如果传递的函数对这个切片的操作不会引起切片的扩容，那么就是底层数组的地址不会改变，因此该函数对切片的操作对这个底层数组上层的切片是可以感知到的，但是假如在这个函数的处理中发生了切片的扩容，则切片的底层数组是新开辟了一块连续内存作为新的数组，这个变量会指向这个新的数组地址。

## 6.基于5的分析，引用类型的切片外的类型以及结构类型的非原始情况，我们都要使用指针而非变量本身，也就是引用类型的值传递。


# 参考链接

- [Go Slice 切片函数间传递到底用变量还是指针？Go语言数据类型，值传递和引用变量类型的思考](https://blog.csdn.net/tony_java_2017/article/details/121593003)
