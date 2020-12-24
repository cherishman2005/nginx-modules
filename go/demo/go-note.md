# go-note



## main示例

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

# 参考链接

- [https://studygolang.com/articles/9467](https://studygolang.com/articles/9467)
