# go语言中五种字符串的拼接方式

## +拼接方式

这种方式是我在写golang经常用的方式，go语言用+拼接，php使用.拼接，不过由于golang中的字符串是不可变的类型，因此用 + 连接会产生一个新的字符串对效率有影响。

```
func main() {
    s1 := "hello"
    s2 := "word"
    s3 := s1 + s2
    fmt.Print(s3) //s3 = "helloword"
}
```

## sprintf函数

```
s1 := "hello"
s2 := "word"
s3 := fmt.Sprintf("%s%s", s1, s2) //s3 = "helloword"
```

这种方式也是开发过程中经常使用到的，这样写的好处就是不会直接产生临时字符串，但是效率好像也是不是特别高。

## Join函数

使用Join函数我们需要先引入strings包才能调用Join函数。Join函数会先根据字符串数组的内容，计算出一个拼接之后的长度，然后申请对应大小的内存，一个一个字符串填入，在已有一个数组的情况下，这种效率会很高，如果没有的话效率也不高。我一般用来切片转字符串使用。

```
s1 := "hello"
s2 := "word"
var str []string = []string{s1, s2}
s3 := strings.Join(str, "")
fmt.Print(s3)
```

## buffer.Builderbuffer.WriteString函数

```
s1 := "hello"
s2 := "word"
var bt bytes.Buffer
bt.WriteString(s1)
bt.WriteString(s2)
s3 := bt.String()
fmt.Println(s3)
```
效率比上面的高不少。

## buffer.Builder函数

```
s1 := "hello"
s2 := "word"
var build strings.Builder
build.WriteString(s1)
build.WriteString(s2)
s3 := build.String()
fmt.Println(s3)
```

官方建议使用的的拼接方式，和上面的使用方法差不多，一般情况下用+拼接，如果拼接的字符串比较长的话最好使用最后一种方式。
