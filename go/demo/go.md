# main示例

main.go
```
package main

import(
    "fmt"
)

func foo (args ...int) {
    fmt.Println(len(args))
}

func main() {
    foo() // 0
    foo(1,2,3) // 3
}
```

编译运行：
```
go build main.go

./main
```