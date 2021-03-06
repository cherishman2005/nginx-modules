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

# Go语言fmt.Sprintf（格式化输出）

|  动  词      |   功  能  |
| :---------- | :------ |
| %v      | 按值的本来值输出 |
| %+v      | 在 %v 基础上，对结构体字段名和值进行展开 |
| %#v      | 输出 Go 语言语法格式的值 |
| %T       | 输出 Go 语言语法格式的类型和值 |
| %%       | 输出 % 本体 |
| %b       | 整型以二进制方式显示 |
| %o       | 整型以八进制方式显示 |
| %d       | 整型以十进制方式显示 |
| %x       | 整型以十六进制方式显示 |
| %X       | 整型以十六进制、字母大写方式显示 |
| %U       | Unicode 字符 |
| %f       | 浮点数 |
| %p       | 指针，十六进制方式显示 |



# 参考链接

- [https://studygolang.com/articles/9467](https://studygolang.com/articles/9467)

- [https://stackoverflow.com/questions/38362631/go-error-non-constant-array-bound](https://stackoverflow.com/questions/38362631/go-error-non-constant-array-bound)

- [https://www.runoob.com/go/go-arrays.html](https://www.runoob.com/go/go-arrays.html)

- [https://learnku.com/go/t/23460/bit-operation-of-go](https://learnku.com/go/t/23460/bit-operation-of-go)

- [Go net/http包](https://studygolang.com/articles/9467)

- [go goroutine 学习笔记](https://juejin.cn/post/6844904141911752718)

- [https://zhuanlan.zhihu.com/p/34211611](https://zhuanlan.zhihu.com/p/34211611)

- [golang操作mysql使用总结](https://my.oschina.net/u/3553591/blog/1630617)

- [Golang slice 增删改查](https://www.codeleading.com/article/1035931752/)

- [Golang slice 增删改查](https://cloud.tencent.com/developer/article/1428643)

- [Go并发调度器解析之实现一个高性能协程池](https://zhuanlan.zhihu.com/p/37754274)

- [Golang 入门 : goroutine(协程)](https://www.cnblogs.com/sparkdev/p/10930168.html)

- [使用chan的时候选择对象还是指针](https://www.cnblogs.com/yjf512/p/10417698.html)

- [Go语言fmt.Sprintf（格式化输出）](http://c.biancheng.net/view/41.html)

## go-mysql

- [https://github.com/s1s1ty/go-mysql-crud](https://github.com/s1s1ty/go-mysql-crud)