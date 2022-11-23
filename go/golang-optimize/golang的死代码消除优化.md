# 什么是死代码消除？

死码消除(dead code elimination, DCE)是一种编译器优化技术，用处是在编译阶段去掉对程序运行结果没有任何影响的代码。

# Go中应用分析
使用常量提升性能，先看一段代码：
main.go
```
func max(a, b int) int {
  if a > b {
      return a
  }
  return b
}
var a, b = 10, 20

func main() {
  if max(a, b) == a {
    fmt.Println(a)
  }
}
拷贝一个mainConst.go文件，将a、b改为const：
func max(a, b int) int {
  if a > b {
      return a
  }
  return b
}
const a, b = 10, 20

func main() {
  if max(a, b) == a {
    fmt.Println(a)
  }
}
```
分别编译main.go和mainConst.go，生成的二进制文件大小如下：
```
% go build -o main main.go
% go build -o mainConst mainConst.go
% ls -l main mainconst
-rwxr-xr-x  1 yalla  staff  2027920 Apr 16 17:37 main
-rwxr-xr-x  1 yalla  staff  1825824 Apr 16 17:37 mainconst
```
我们可以看到 mainConst 比 main体积小了约 10%。
为什么会产生这种情况呢？
首先max函数进行了内联优化（具体参考我上篇文章），优化后如下：
```
const a, b = 10, 20
func main() {
    var result int
    if a > b {
        result = a
    } else {
        result = b
    }
    if result == a {
        fmt.Println(a)
    }
}
```
因为a和b是常量，编译器可以在编译时证明该分支永远不会为true。10总是小于20。因此编译器可以进一步优化main为：
```
const a, b = 10, 20
func main() {
    var result int
    if false {
        result = a
    } else {
        result = b
    }
    if result == a {
        fmt.Println(a)
    }
}
```

既然分支的结果是已知的，那么result的内容也是已知的。这就是调用分支消除：

```
const a, b = 10, 20
func main() {
  const result = b
  if result == a {
      fmt.Println(a)
  }
}
```

现在这个分支被消除了，我们知道结果总是等于b，因为b是一个常数，所以我们知道结果是一个常数。编译器将这个证明应用到第二个分支：

```
func main() {
  const result = b
  if false {
    fmt.Println(a)
  }
}
```

再次使用分支消除法将main的最终形式简化为：
```
func main() {
  const result = b
}
```

最后结果就是：
```
func main() {
}
```

结论分析：当全局变量 a、b 不为常时量，即main中声明的一样，编译器并不知道运行过程中a、b 会不会发生改变，因此不能够进行死代码消除，这部分代码被编译到最终的二进制程序中。因此 main 比 mainConst 二
进制体积大了约 10%。

# 服务调试应用

我们可以在源代码中，定义全局常量debugModel，值设置为false，在需要增加调试代码的地方，使用条件语句

if debugModel进行判断，如下：
```
const debugModel = false
func main() {
   if debugModel {
      log.Println("debug is enabled")
   }
}
```

如果是正常编译，常量debugModel始终等于false，调试语句在编译过程中会被消除，不会影响最终的二进
制大小。如果我们想编译出debug版本的二进制，将debugModel修改为true之后编译。这对于开发者日常调试非常有帮助，主要用于日常开发过程中进行单元测试或者是简单的集成测试。最终提交代码时，将debugModel
修改为false即可。
