# goroutine闭包的坑

```
package main

import (
    "fmt"
    "time"
    //"unsafe"
)

func main() {
    // map
    m := make(map[string]bool)
    m["ni"] = true
    m["hao"] = true
    m["hello"] = true
    m["world"] = true
    
    for k := range m {
        go func(){
            fmt.Println(k)
        }()
    }

    time.Sleep(time.Second * 2)
}

//func main() {
//    for i := 0; i < 10; i++ {
//        go func(){
//            fmt.Println(i)
//        }()
//    }
//
//    time.Sleep(time.Second * 2)
//}
```

运行结果

```
# ./goroutine_demo 
world
world
world
world
# ./goroutine_demo 
hello
hello
hello
hello
```
非预期结果

## 正确写法1

中间临时变量

```
package main

import (
    "fmt"
    "time"
    //"unsafe"
)

func main() {
    // map
    m := make(map[string]bool)
    m["ni"] = true
    m["hao"] = true
    m["hello"] = true
    m["world"] = true
    
    for k := range m {
        tmp := k
        go func(){
            fmt.Println(tmp)
        }()
    }

    time.Sleep(time.Second * 2)
}

//func main() {
//    for i := 0; i < 10; i++ {
//        go func(){
//            fmt.Println(i)
//        }()
//    }
//
//    time.Sleep(time.Second * 2)
//}
```

## 正确写法2

参数传递

```
package main

import (
    "fmt"
    "time"
    //"unsafe"
)

func main() {
    // map
    m := make(map[string]bool)
    m["ni"] = true
    m["hao"] = true
    m["hello"] = true
    m["world"] = true
    
    for k := range m {
        go func(name string){
            fmt.Println(name)
        }(k)
    }

    time.Sleep(time.Second * 2)
}

//func main() {
//    for i := 0; i < 10; i++ {
//        go func(){
//            fmt.Println(i)
//        }()
//    }
//
//    time.Sleep(time.Second * 2)
//}
```

运行结果：
```
# ./goroutine_demo 
ni
hao
hello
world
```

# 小结

* 中间临时变量 简单

* 参数传递 代码更加紧凑，没有赘余感。

# 参考链接

- [Go 语言使用 goroutine 运行闭包的“坑”](https://www.51cto.com/article/715815.html)
