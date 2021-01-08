# go-note

## array增删改查

```
func search(slice *[]uint64, value uint64) int {
	for i, v := range *slice {
		if v == value {
			return i
		}
	}
	return -1
}

func delete(slice *[]uint64, index int) {
	if index > len(*slice) - 1 {
	   return  
	}
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}

func clear(slice *[]uint64) {
	*slice = []uint64{}
}
```

## golang map 判断key是否存在

判断方法示例代码
```
if _, ok := map[key]; ok {
    // 存在
}
 
if _, ok := map[key]; !ok {
    // 不存在
}
```

判断方式为value,ok := map[key], ok为true则存在


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

- [Go net/http包](https://studygolang.com/articles/9467)

