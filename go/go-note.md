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

## go数组内存分配示例

调整数组顺序使奇数位于偶数前面
```
func exchange(nums []int) []int {
    i := 0
    j := len(nums)-1
    res := make([]int, len(nums))
    for _, value := range nums {
        if value&1 == 1  {
            res[i] = value
            i++
        } else {
            res[j] = value
            j--
        }
    }

    return res
}
```

# 参考链接

- [https://studygolang.com/articles/9467](https://studygolang.com/articles/9467)

- [https://stackoverflow.com/questions/38362631/go-error-non-constant-array-bound](https://stackoverflow.com/questions/38362631/go-error-non-constant-array-bound)

- [https://www.runoob.com/go/go-arrays.html](https://www.runoob.com/go/go-arrays.html)

- [https://learnku.com/go/t/23460/bit-operation-of-go](https://learnku.com/go/t/23460/bit-operation-of-go)