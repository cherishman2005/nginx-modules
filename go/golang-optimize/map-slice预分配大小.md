# map-slice预分配大小

```
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    // map
    m := make(map[int]int, 5)
    fmt.Println("sizeof(m)=", unsafe.Sizeof(m))
    fmt.Println("before, size=", len(m))
    for i := 0; i < 10; i++ {
        m[i] = i+1
    }
    fmt.Println("after, size=", len(m))
    fmt.Println(m)
    
    fmt.Println()
    
    // slice
    s := make([]string, 0, 5)
    fmt.Println("sizeof(s)=", unsafe.Sizeof(s))
    fmt.Println("before, size=", len(s))
    for i := 0; i < 10; i++ {
        s = append(s, "aa")
    }
    fmt.Println("after, size=", len(s))
    fmt.Println(s)
}
```

运行结果
```
sizeof(m)= 8
before, size= 0
after, size= 10
map[0:1 1:2 2:3 3:4 4:5 5:6 6:7 7:8 8:9 9:10]

sizeof(s)= 24
before, size= 0
after, size= 10
[aa aa aa aa aa aa aa aa aa aa]
```

* map/slice可以预先分配大小，不是只能使用预先分配的大小
