# go语言map和slice深拷贝

## map深拷贝

map 其实是不能拷贝的，如果想要拷贝一个 map ，只有一种办法就是循环赋值，就像这样
```
originalMap := make(map[string]int)
originalMap["one"] = 1
originalMap["two"] = 2

// Create the target map
targetMap := make(map[string]int)

// Copy from the original map to the target map
for key, value := range originalMap {
    targetMap[key] = value
}
```
如果 map 中有指针，还要考虑深拷贝的过程
```
originalMap := make(map[string]*int)
var num int = 1
originalMap["one"] = &num

// Create the target map
targetMap := make(map[string]*int)

// Copy from the original map to the target map
for key, value := range originalMap {
var tmpNum int = *value
    targetMap[key] = &tmpNum
}
```
如果想要更新 map 中的value，可以通过赋值来进行操作
```
map["one"] = 1
```
但如果 value 是一个结构体，可以直接替换结构体，但无法更新结构体内部的值
```
originalMap := make(map[string]Person)
originalMap["minibear2333"] = Person{age: 26}
originalMap["minibear2333"].age = 5
```
你可以 试下源码函数[脚注1] 会报这个错误
```
★ Cannot assign to originalMap["minibear2333"].age
”
```
问题链接 issue-3117[脚注2] , 其中 ianlancetaylor[脚注3] 的回答很好的解释了这一点

简单来说就是map不是一个并发安全的结构，所以，并不能修改他在结构体中的值。

这如果目前的形式不能修改的话，就面临两种选择，

1.修改原来的设计;
2.想办法让map中的成员变量可以修改，
因为懒得该这个结构体，就选择了方法2

要么创建个临时变量，做拷贝，像这样
```
tmp := m["foo"]
tmp.x = 4
m["foo"] = tmp
```
要么直接用指针，比较方便
```
originalPointMap := make(map[string]*Person)
originalPointMap["minibear2333"] = &Person{age: 26}
originalPointMap["minibear2333"].age = 5
```
## slice复制陷阱

切片有一种方式复制方式，比较快速
```
 slice3 :=  slice2[:]
```
但是有一种致命的缺点，这是浅拷贝，slice3和slice2是同一个切片，无论改动哪个，另一个都会产生变化。

可能这么说你还是不能加深理解。在源码bytes.buffer[脚注4]中出现了这一段
```
func (b *Buffer) Bytes() []byte {
    return b.buf[b.off:] 
}
```
我们在读入读出输入流的时候，极易出现这样的问题

下面的例子，使用abc模拟读入内容，修改返回值内容
```
 buffer := bytes.NewBuffer(make([]byte, 0, 100))
 buffer.Write([]byte("abc"))
 resBytes := buffer.Bytes()
 fmt.Printf("%s \n", resBytes)
 resBytes[0] = 'd'
 fmt.Printf("%s \n", resBytes)
 fmt.Printf("%s \n", buffer.Bytes())
```
输出，可以看出会影响到原切片内容
```
abc
dbc
dbc
```
这种情况在并发使用的时候尤为危险，特别是流式读写的时候容易出现上一次没处理完成，下一次的数据覆盖写入的错乱情况

# 参考链接

* updateMapValue：https://github.com/golang-minibear2333/golang/blob/master/2.func-containers/2.4-map/map1.go#L89

* issue-3117：https://github.com/golang/go/issues/3117

* ianlancetaylor：https://github.com/golang/go/issues/3117#issuecomment-430632750

* 源码 https://github.com/golang/go/bl

