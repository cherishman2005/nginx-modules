# golang运行先后顺序

```
package main                                                                                                                     

import (
   "fmt"              
)

var T int64 = a()

func init() {
   fmt.Println("init in main.go ")
}

func a() int64 {
   fmt.Println("calling a()")
   return 2
}
func main() {                  
   fmt.Println("calling main")     
}
```

## 运行结果

```
calling a()
init in main.go
calling main
```
* 初始化顺序：变量初始化->init()->main()
