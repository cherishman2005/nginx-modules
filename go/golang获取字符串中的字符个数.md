# golang获取字符串中的字符个数

在 Go 语言中，是不能用 len 函数来统计字符串中的字符个数的，这是因为在 Go 中，字符串是以 UTF-8 为格式进行存储的，在字符串上调用 len 函数，取得的是字符串包含的 byte 的个数。

例如字符串：Hello, 世界 包含 9 个字符。使用 len("Hello, 世界") 得到的结果为 13，因为这个字符串占用 13 个字节：
```
s := "Hello, 世界"
fmt.Println(len(s)) // 13
fmt.Println([]byte(s)) // [72 101 108 108 111 44 32 228 184 150 231 149 140]
```
那么如何统计 Golang 字符串中的字符个数呢？有下面几种方法：

使用 bytes.Count() 统计
使用 strings.Count() 统计
将字符串转换为 []rune 后调用 len 函数进行统计
使用 utf8.RuneCountInString() 统计
具体代码：
```
package count

import "bytes"
import "strings"
import "unicode/utf8"

func f1(s string) int {
    return bytes.Count([]byte(s), nil) - 1
}

func f2(s string) int {
    return strings.Count(s, "") - 1
}

func f3(s string) int {
    return len([]rune(s))
}

func f4(s string) int {
    return utf8.RuneCountInString(s)
}
```
用上面 4 个函数计算字符串 Hello, 世界 都会得到正确的字符个数：9。那究竟哪个方法更好一些呢？做一下测试看看：

创建文件：count_test.go，内容如下：
```
package count

import "bytes"
import "strings"
import "unicode/utf8"
import "testing"

func f1(s string) int {
    return bytes.Count([]byte(s), nil) - 1
}

func f2(s string) int {
    return strings.Count(s, "") - 1
}

func f3(s string) int {
    return len([]rune(s))
}

func f4(s string) int {
    return utf8.RuneCountInString(s)
}

var s = "Hello, 世界"

func Benchmark1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        f1(s)
    }
}

func Benchmark2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        f2(s)
    }
}

func Benchmark3(b *testing.B) {
    for i := 0; i < b.N; i++ {
        f3(s)
    }
}

func Benchmark4(b *testing.B) {
    for i := 0; i < b.N; i++ {
        f4(s)
    }
}
```
在命令行中运行命令：go test count_test.go -bench ".*"，输出如下内容：
```
testing: warning: no tests to run
PASS
Benchmark1	20000000	       109 ns/op
Benchmark2	20000000	        61.6 ns/op
Benchmark3	10000000	       220 ns/op
Benchmark4	30000000	        58.1 ns/op
ok  	command-line-arguments	7.843s
```
看来，速度最快的是 utf8.RuneCountInString()。
