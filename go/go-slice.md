# go 切片（增删改查、内存消耗）

go语言中的切片，可以看作是可变化长度的数组（动态数组）。有长度(len)和容量(cap)，容量必大于等于长度。

切片的结构体定义如下：
```
type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}
```

可以使用 reflect.SliceHeader来访问切片信息。如果碰见切片为nil，但是长度和容量不为0，说明切片损坏了（例：通过 reflect.SliceHeader来修改切片）。

切片的定义方式
```
var (
    a []int               // nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
    b = []int{}           // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
    c = []int{1, 2, 3}    // 有3个元素的切片, len和cap都为3
    d = c[:2]             // 有2个元素的切片, len为2, cap为3
    e = c[0:2:cap(c)]     // 有2个元素的切片, len为2, cap为3
    f = c[:0]             // 有0个元素的切片, len为0, cap为3
    g = make([]int, 3)    // 有3个元素的切片, len和cap都为3
    h = make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
    i = make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
)
```

遍历
````
	for i := range a {
        fmt.Printf("a[%d]: %d\n", i, a[i])
    }
    for i, v := range b {
        fmt.Printf("b[%d]: %d\n", i, v)
    }
    for i := 0; i < len(c); i++ {
        fmt.Printf("c[%d]: %d\n", i, c[i])
    }
```

添加元素
开头添加（一般都会导致重新分配内存）
```
var a = []int{1,2,3}
a = append([]int{0}, a...)        // 在开头添加1个元素
a = append([]int{-3,-2,-1}, a...) // 在开头添加1个切片
```

末尾添加（注意：容量不足，append会重新分配内存）
```
var a []int
a = append(a, 1)               // 追加1个元素
a = append(a, 1, 2, 3)         // 追加多个元素, 手写解包方式
a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包
```

中间添加（append和copy组合实现）
```
// append实现，第二个append调用会创建一个临时的切片
var a []int
a = append(a[:i], append([]int{x}, a[i:]...)...)     // 在第i个位置插入x
a = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入切片

// append和copy组合实现，避免创建中间的临时切片
a = append(a, 0) // 先扩容
copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置
a[i] = x             // 设置新添加的元素

// append和copy组合，在指定位置插入切片（多个元素）
a = append(a, x...) // 没有专门的函数来扩容，只有使用append
copy(a[i+len(x):], a[i:])
copy(a[i:], x)
```

## 删除元素

### 删除头部

```
a := []int{1, 2, 3}

// 移动数据指针
a = a[1:] // 删除开头1个元素
a = a[N:] // 删除开头N个元素

// 可以用append原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）
a = append(a[:0], a[1:]...) // 删除开头1个元素
a = append(a[:0], a[N:]...) // 删除开头N个元素

// 用copy完成删除开头的元素
a = a[:copy(a, a[1:])] // 删除开头1个元素
a = a[:copy(a, a[N:])] // 删除开头N个元素
```

### 删除尾部（最快）
```
a = a[:len(a)-1]   // 删除尾部1个元素
a = a[:len(a)-N]   // 删除尾部N个元素
```

### 删除中间部分

```
// 使用append
a = append(a[:i], a[i+1:]...) // 删除中间1个元素
a = append(a[:i], a[i+N:]...) // 删除中间N个元素

// 使用copy
a = a[:i+copy(a[i:], a[i+1:])]  // 删除中间1个元素
a = a[:i+copy(a[i:], a[i+N:])]  // 删除中间N个元素
```

# 总结
切片高效操作的要点是要降低内存分配的次数，尽量保证append操作不会超出cap的容量，降低触发内存分配的次数和每次分配内存大小。

# 参考

- [go语言高级编程](https://books.studygolang.com/advanced-go-programming-book/ch1-basic/ch1-03-array-string-and-slice.html)
