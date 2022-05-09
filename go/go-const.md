# go-const

Go学习：常量报错 const initializer is not a constant

Go的常量const是属于编译时期的常量，即在编译时期就可以完全确定取值的常量。只支持数字，字符串和布尔，及上述类型的表达式。而切片，数组，正则表达式等等需要在运行时分配空间和执行若干运算才能赋值的变量则不能用作常量。这一点和Java,Nodejs(javascript)不同。Java的final和Nodejs的const代表的是一次性赋值的变量，本质上还是变量，只是不允许后续再做修改，任意类型都可以，可以在运行时赋值。

可以这样类比：Go的常量对应于C#的const,而Java，Nodejs的常量对应于C#的readonly。
```
package main
 
import(
	"regexp"
)
 
//正确
const R1 = 1
const R2 = 1.02
const R3 = 1 * 24 * 1.03
const R4 = "hello" + " world"
const R5 = true
 
//错误： const initializer ... is not a constant
const R6 = [5]int{1,2,3,4,5}
const R7 = make([]int,5)
const R8 = regexp.MustCompile(`^[a-zA-Z0-9_]*$`)
 
func main(){
 
}
```
编译报错：
```
./checkconst.go:15:7: const initializer [5]int literal is not a constant
./checkconst.go:16:7: const initializer make([]int, 5) is not a constant
./checkconst.go:17:7: const initializer regexp.MustCompile("^[a-zA-Z0-9_]*$") is not a constant
```
