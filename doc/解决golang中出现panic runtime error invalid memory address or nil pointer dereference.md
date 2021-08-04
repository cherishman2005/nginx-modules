# 解决golang中出现panic runtime error invalid memory address or nil pointer dereference

当在nil上调用一个属性或者方法的时候 , 会报空指针

尤其是结构体指针 , 非常容易出现这个问题 , 下面是测试代码

```
package main

import "fmt"

func main() {
    type MConn struct {
        Name string
    }
    var conn *MConn
    var conn2 MConn
    conn3 := new(MConn)
    conn4 := &MConn{}
    fmt.Printf("%v,%v,%v,%v", conn, conn2, conn3, conn4)
}
```

分别返回
```
<nil>,{},&{},&{}
```

当声明了一个结构体指针变量var conn *MConn , 但是没有初始化 , 直接调用属性时候 , 就会出现
```
panic: runtime error: invalid memory address or nil pointer dereference
```
因为conn这个时候是nil, 是个空指针

一定要进行判空操作 , if conn != nil {}

 

当然我们有时候不会出现这么明显的错误 , 但是在和map进行配合时 , 无意中可能会出现这个错误
```
var mMap map[string]*MConn
m1 := mMap["name"]
m1.Name = "qqq"
```
这个代码map中 , 当key元素不存在时 , 返回的是value的零值 , 恰好是*MConn 零值是nil , 也会报错

所以map这里也要进行判断
```
var mMap map[string]*MConn
m1, ok := mMap["name"]
if ok {
    m1.Name = "qqq"
}
```
