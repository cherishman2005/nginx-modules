# go语言内存对齐

1、一些基础知识
内存对齐的底层原因是内存IO是以字长(64bit)为单 位进行的。每次内存IO获取数据都是从同行同列的8个chip中各自读取一个字节拼起来。从内存的0地址开始，0-63bit的数据可以一次从IO读取出来，64-127bit的数据也可以一次读取出来。CPU和内存IO的硬件限制导致没有办法一次跨在两个数据宽度中间进行IO。

例如对于C程序来说，如果把一个64bit的int写入内存，从地址0x0001开始，而不是从0x000开始，那么数据并没有存在同一行列地址上。因此CPU必须得让内存工作两次才能取于完整的数据。

在类型的值在内存对齐的情况下，计算机的加载和写入会很高效。例如，一个两字节值（如int16)的地址应该是一个偶数，一个四字节值（如rune)的地址应该是4的整数倍，一个8字节值如（float64、uint64或64位指针）的地址应该是8的整数倍。

看下面的示例
```
package main

import (
        "unsafe"
        "fmt"
)

func main() {
        var a int
        var b int16
        var c rune
        fmt.Printf("%v\t%v\t%v\n", &a,unsafe.Sizeof(a),unsafe.Alignof(a))
        fmt.Printf("%v\t%v\t%v\n", &b,unsafe.Sizeof(b),unsafe.Alignof(b))
        fmt.Printf("%v\t%v\t%v\n", &c,unsafe.Sizeof(c),unsafe.Alignof(c))

        var m map[string]int  = make(map[string]int)
        fmt.Printf("%v\t%v\t%v\n", &m,unsafe.Sizeof(m),unsafe.Alignof(m))
        
        var mb map[string]bool  = make(map[string]bool)
        fmt.Printf("%v\t%v\t%v\n", &mb,unsafe.Sizeof(mb),unsafe.Alignof(mb))
}
```

输出为


```
0xc000016098    8       8
0xc0000160b0    2       2
0xc0000160b4    4       4
&map[]  8       8
&map[]  8       8
```

每一次分配给变量a的地址总是8的整数倍。而分配给b的地址总是2的整数倍。分配给c的地址总是4的整数倍。

2、聚合类型
聚合类型（结构体或数组）的值的长度至少是它的成员或元素之和，并且由于内存对齐的原因，或许比这个更大。内存空位是由编译器添加的未使用的内存地址，用来确保连续的成员或者元素相对于结构体或数组的起始地址是对齐的。

| 类型  | 大小（sizes are guaranteed） |
| ------------- | ------------- |
| bool   | 1个字  |
| intN、uintN、floatN、complexN  | N/8字节，如float64是8字节  |
| int、uint、uintptr  | 32位机：4字节； 64位机：8字节  |
| *T  | 1个字  |
|  string | 2个字（数据、长度） |
| []T  | 3个字（指向底层数组的指针、长度和容量）长度和容量都 是int，即都是1个字，指向底层数组的指针也是1个字。因此[]T是3个字。  |
| map  | 1个字  |
| func  | 1个字  |
| chan  | 1个字  |
| interface  | 两个字（类型、值） |

如果结构体成员的类型是不同的，那么将相同类型的成员定义在一起，sizes are guaranteed的值大的定义在前可以更节约内存空间。


示例
```
package main

import (
        "fmt"
        "unsafe"
)

type T1 struct {
        flag    bool
        pi      float64
        counter int16
}
type T2 struct {
        pi      float64
        counter int16
        flag    bool
}

type T3 struct {
        flag    bool
        counter int16
        pi      float64
}

func main() {
        var t1 T1
        var t2 T2
        var t3 T3
        var t4 struct {
                a bool
                b int16
                c []int
        }
        var t5 struct {
                a bool
                b int16
        }
        var t6 struct {
                a bool
        }
        var t7 struct {
                c []byte
        }
        var t8 struct {
                c []int
        }

        fmt.Printf("%v\t%v\t%v\t%v\t%v\n", unsafe.Sizeof(t1), unsafe.Alignof(t1), unsafe.Alignof(t1.flag), unsafe.Alignof(t1.pi), unsafe.Alignof(t1.counter))
        fmt.Printf("%v\t%v\t%v\t%v\t%v\n", unsafe.Sizeof(t2), unsafe.Alignof(t2), unsafe.Alignof(t2.pi), unsafe.Alignof(t2.counter), unsafe.Alignof(t2.flag))
        fmt.Printf("%v\t%v\t%v\t%v\t%v\n", unsafe.Sizeof(t3), unsafe.Alignof(t3), unsafe.Alignof(t3.flag), unsafe.Alignof(t3.counter), unsafe.Alignof(t3.pi))
        fmt.Printf("%v\t%v\t%v\t%v\t%v\n", unsafe.Sizeof(t4), unsafe.Alignof(t4), unsafe.Alignof(t4.a), unsafe.Alignof(t4.b), unsafe.Alignof(t4.c))
        fmt.Printf("%v\t%v\t%v\t%v\n", unsafe.Sizeof(t5), unsafe.Alignof(t5), unsafe.Alignof(t5.a), unsafe.Alignof(t5.b))
        fmt.Printf("%v\t%v\t%v\n", unsafe.Sizeof(t6), unsafe.Alignof(t6), unsafe.Alignof(t6.a))
        fmt.Printf("%v\t%v\t%v\n", unsafe.Sizeof(t7), unsafe.Alignof(t7), unsafe.Alignof(t7.c))
        fmt.Printf("%v\t%v\t%v\n", unsafe.Sizeof(t8), unsafe.Alignof(t8), unsafe.Alignof(t8.c))
}
```
输出结果


```
$ ./u2 
24      8       1       8       2
16      8       8       2       1
16      8       1       2       8
32      8       1       2       8
4       2       1       2
1       1       1
24      8       8
24      8       
```

结果显示t1占3个字，t2、t3分别占2个字。t4占4个字。

3 对齐规则
内存对齐会影响聚合类型如struct的内存占用大小。对于某一具体的成员而言，有一个unsafe.Alignof的值，同时还有一个值是计算机字长对应的值。如64位计算机的字长为8个字节。两者中取最小值作为该成员的对齐值。

看下面的例子：
```
package main

import (
        "fmt"
        "unsafe"
)

func main() {
        var x struct {
                a bool
                b int16
        }
        fmt.Printf("%v\t%v\t%v\t%v\n", &x.a,&x.b,unsafe.Offsetof(x.a),unsafe.Offsetof(x.b))
        var y struct{
                a bool
                b int16
                c int32
        }
        var z struct {
                a bool
                b int16
                c int32
                d int64
        }
        fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", &y.a, &y.b,&y.c,unsafe.Offsetof(y.a),unsafe.Offsetof(y.b),unsafe.Offsetof(y.c))
        fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", &z.a, &z.b,&z.c,unsafe.Offsetof(z.a),unsafe.Offsetof(z.b),unsafe.Offsetof(z.c),unsafe.Offsetof(z.d))
}
```

输出结果为
```
0x400012e010    0x400012e012    0       2
0x400012e018    0x400012e01a    0x400012e01c    0       2       4
0x400012e030    0x400012e032    0x400012e034    0x400012e038    0       2       4       8
```
对于结构体变量x，有两个成员x.a和x.b，x.a的sizes are guaranteed值为1<8。x.b的sizes are guaranteed为2。考虑x.a和x.b，它们的sizes are guaranteed最大值为2<8。所以x的地址分配偏移量要从2的倍数开始。示例中0x400012e010能被2整除。当在0x400012e010处为x.a分配了地址之后，由于x.a只占用1个字节。而x.b要从2的倍数偏移量开始分。所以x.b只能从0x400012e012处开始分。x.a和x.b之间会空1个字节。

对于结构体变量y，有三个成员y.a、y.b和y.c，三个成员的sizes are guaranteed分别是1、2和4，均小于8。所以y的地址分配偏移量要从4的倍数开始。示例中0x400012e018能被4整作。当在0x400012e018处为y.a分配了地址之后，y.a只占用1个字节，而y.b要从2的倍数偏移量开始分。所以y.b只能从0x400012e01a处开始分配。y.a和y.b之间空了1个字节。因为y.c的sizes are guaranteed为4，所以y.c的地址分配偏移量要从4的倍数开始分。所以y.c从0x400012e01c开始分配。y.b和y.c之间没有空字节。

对于结构体变量z，有四个成员z.a、z.b、z.c和z.d，它们的sizes are guaranteed分别是1、2、4和8。所以整个z的地址分配偏移量要从8的倍数开始。示例中0x400012e030能被8整除。z.a从0x400012e030开始，占1个字节，z.b要从2的倍数偏移，z.b只能从0x400012e032开始，z.a和z.b之间空1个字节。z.c要从4的倍数偏移，所以z.c要从0x400012e034开始。z.d刚好从0x400012e038开始。

再看一个例子
```
package main

import (
        "fmt"
        "unsafe"
)

func main() {
        var x struct {
                a bool
                b int16
                c []int
        }
        fmt.Printf("%v\t%v\n",unsafe.Sizeof(x), unsafe.Alignof(x))
        fmt.Printf("%v\t%v\n",unsafe.Alignof(x.a),unsafe.Offsetof(x.a))
        fmt.Printf("%v\t%v\n",unsafe.Alignof(x.b),unsafe.Offsetof(x.b))
        fmt.Printf("%v\t%v\n",unsafe.Alignof(x.c),unsafe.Offsetof(x.c))
}
```
输出为
```
32      8
1       0x4000070020    0
2       0x4000070022    2
8       &[]     8
```
一共占用了32个字节即4个字。x.a x.b x.c的Alignof分别为1、2和8，最大为8。所以x要从8的倍数的偏移个开始分。从 0x4000070020开始分是满足要求的。a点1字节，而b要从2的倍数的偏移开始，所以b接着要从 0x4000070022开始分，x.a和x.b之间会空1个字节。x.c要从8的倍数偏移开始。x.c实际上要从 0x4000070028开始。x.b和x.c之间空4个字节。

结论：

（1）在结构体所有成员中，找出Alignof最大的值（最大值不会超过8），整体结构体的内存分配从最大值的倍数地址偏移开始分配。

（2）逐个成员根据Sizeof的长度分配地址空间，每个成员都要从它的Alignof的倍数地址偏移开始分。相邻两个成员可能会空出若干字节。

（3）在设计结构体的时候，最好的方式是将Alignof的值从高到低排序后，再写入结构体。

4、一个面试的例子
听说有一道腾讯的面试题

下面的程序输出什么？
```
package main

import (
        "fmt"
        "unsafe"
)

type S struct {
        A       uint32
        B       uint64
        C       uint64
        D       uint64
        E       struct{}
}

func main() {
        fmt.Println(unsafe.Offsetof(S{}.E))
        fmt.Println(unsafe.Sizeof(S{}.E))
        fmt.Println(unsafe.Sizeof(S{}))
}
```
输出为
```
32
0
40
```
可以明确S是8字节对齐的。为什么不是 32 0 32而是32 0 40呢？

一个非空结构体包含有尾部size为0的变量(字段)，如果不给它分配内存，那么该变量(字段)的指针地址将指向一个超出该结构体内存范围的内存空间。这可能会导致内存泄漏，或者在内存垃圾回收过程中，程序crash掉。

# 参考链接

- [https://zhuanlan.zhihu.com/p/362433737](https://zhuanlan.zhihu.com/p/362433737)
