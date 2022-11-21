# go语言string之Buffer与Builder

操作字符串离不开字符串的拼接，但是Go中string是只读类型，大量字符串的拼接会造成性能问题。

* 字符串拼接的方式与性能对比？
* bytes.Buffer 与 strings.Builder？
* Buffer 和 Builder底层原理实现？
* 字符串拼接的方式与性能对比

拼接字符串，无外乎四种方式，采用“+”，“fmt.Sprintf()”,"bytes.Buffer","strings.Builder"
```
import(
    "bytes"
    "fmt"
    "strings"
    "time"
)

func AppendWithAdd(n int) {
    var s string
    for i:=0; i < n; i++{
        s = s + "string"
    }
}

func AppendWithSprintf(n int) {
    var s string
    for i:=0; i < n; i++{
        s = fmt.Sprintf("%s%s", s, "string")
    }
}

func AppendWithBytesBuffer(n int) {
    var byt bytes.Buffer
    for i:=0; i < n; i++{
        byt.WriteString("string")
    }
    byt.String()
}

func AppendWithStringBuilder(n int) {
    var sbuilder strings.Builder
    for i:=0; i < n; i++{
        sbuilder.WriteString("string")
    }
    sbuilder.String()
}
```

```
go test -test.bench=.* -count=5
goos: windows
goarch: amd64
pkg: /studyGo/first
Benchmark_AppendWithAdd-8                     42          27607255 ns/op
Benchmark_AppendWithAdd-8                     36          27935317 ns/op
Benchmark_AppendWithAdd-8                     42          27947221 ns/op
Benchmark_AppendWithAdd-8                     42          27915798 ns/op
Benchmark_AppendWithAdd-8                     37          27857249 ns/op
Benchmark_AppendWithSprintf-8                 30          35487310 ns/op
Benchmark_AppendWithSprintf-8                 32          35940897 ns/op
Benchmark_AppendWithSprintf-8                 32          36376653 ns/op
Benchmark_AppendWithSprintf-8                 32          35950091 ns/op
Benchmark_AppendWithSprintf-8                 32          36089072 ns/op
Benchmark_AppendWithBytesBuffer-8          17287             69116 ns/op
Benchmark_AppendWithBytesBuffer-8          17212             69301 ns/op
Benchmark_AppendWithBytesBuffer-8          17262             69235 ns/op
Benchmark_AppendWithBytesBuffer-8          14858            136065 ns/op
Benchmark_AppendWithBytesBuffer-8          17072            102340 ns/op
Benchmark_AppendWithStringBuilder-8        22573             57630 ns/op
Benchmark_AppendWithStringBuilder-8        21070             87849 ns/op
Benchmark_AppendWithStringBuilder-8        26326             53106 ns/op
Benchmark_AppendWithStringBuilder-8        20924             89193 ns/op
Benchmark_AppendWithStringBuilder-8        26348             52523 ns/op
```
上面我们创建10万字符串拼接的测试，可以发现"bytes.Buffer","strings.Builder"的性能最好，约是“+”的1000倍级别。

这是由于string是不可修改的，所以在使用“+”进行拼接字符串，每次都会产生申请空间，拼接，复制等操作，数据量大的情况下非常消耗资源和性能。而采用Buffer等方式，都是预先计算拼接字符串数组的总长度（如果可以知道长度），申请空间，底层是slice数组，可以以append的形式向后进行追加。最后在转换为字符串。这申请了不断申请空间的操作，也减少了空间的使用和拷贝的次数，自然性能也高不少。

bytes.Buffer 与 strings.Builder
bytes.buffer是一个缓冲byte类型的缓冲器存放着都是byte
 是一个变长的 buffer，具有 Read 和Write 方法。 Buffer 的 零值 是一个 空的 buffer，但是可以使用，底层就是一个 []byte， 字节切片。

Buffer的使用
```
var b bytes.Buffer  //直接定义一个 Buffer 变量，而不用初始化
b.Writer([]byte("Hello ")) // 可以直接使用

b1 := new(bytes.Buffer)   //直接使用 new 初始化，可以直接使用
func NewBuffer(buf []byte) *Buffer
func NewBufferString(s string) *Buffer
```

向Buffer中写数据，可以看出Buffer中有个Grow函数用于对切片进行扩容。
```
// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p; err is always nil. If the
// buffer becomes too large, Write will panic with ErrTooLarge.
func (b *Buffer) Write(p []byte) (n int, err error) {
    b.lastRead = opInvalid
    m := b.grow(len(p))
    return copy(b.buf[m:], p), nil
}
```

从Buffer中读取数据
```
func (b *Buffer) Read(p []byte) (n int, err error) {}

 //声明一个空的slice,容量为8
    l := make([]byte, 8)
    //把bufs的内容读入到l内,因为l容量为8,所以只读了8个过来
    bufs.Read(l)
    fmt.Println("::bufs缓冲器内容::")
    fmt.Println(bufs.String())

···
//ReadString需要一个byte作为分隔符，读的时候从缓冲器里找第一个出现的分隔符
func (b *Buffer) ReadString(delim byte) (line string, err error) {}

//返回缓冲器头部的第一个byte，缓冲器头部第一个byte被拿掉
func (b *Buffer) ReadByte() (c byte, err error) {}
```

Builder的使用
```
var sbuilder strings.Builder
sbuilder.WriteString("string")

// 使用new也可以创建
var sb = new(strings.Builder)
sb.Write([]byte("hello"))
fmt.Printf("%s",sb.String())
```

strings.Builder的方法和bytes.Buffer的方法的命名几乎一致。
```
func (b *Builder) WriteByte(c byte) error {
    b.copyCheck()
    b.buf = append(b.buf, c)
    return nil
}
```
但实现并不一致，Builder的Write方法直接将字符拼接slice数组后。

其没有提供read方法，但提供了strings.Reader方式

Reader 结构:
```
type Reader struct {
    s        string
    i        int64 // current reading index
    prevRune int   // index of previous rune; or < 0
}
```

Buffer 和 Builder底层原理实现
Buffer:
```
// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
type Buffer struct {
    buf      []byte // contents are the bytes buf[off : len(buf)]
    off      int    // read at &buf[off], write at &buf[len(buf)]
    lastRead readOp // last read operation, so that Unread* can work correctly.
}
```

Builder:
```
// A Builder is used to efficiently build a string using Write methods.
// It minimizes memory copying. The zero value is ready to use.
// Do not copy a non-zero Builder.
type Builder struct {
    addr *Builder // of receiver, to detect copies by value
    buf  []byte
}
```
可以看出Buffer和Builder底层都是采用[]byte数组进行装载数据。

先来说说Buffer:
```
// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p; err is always nil. If the
// buffer becomes too large, Write will panic with ErrTooLarge.
func (b *Buffer) Write(p []byte) (n int, err error) {
    b.lastRead = opInvalid
    m, ok := b.tryGrowByReslice(len(p))
    if !ok {
        m = b.grow(len(p))
    }
    return copy(b.buf[m:], p), nil
}
```
创建好Buffer是一个empty的，off 用于指向读写的尾部。
 在写的时候，先判断当前写入字符串长度是否大于Buffer的容量，如果大于就调用grow进行扩容，扩容申请的长度为当前写入字符串的长度。如果当前写入字符串长度小于最小字节长度64，直接创建64长度的[]byte数组。如果申请的长度小于二分之一总容量减去当前字符总长度，说明存在很大一部分被使用但已读，可以将未读的数据滑动到数组头。如果容量不足，扩展2*c + n 。
```
// To build strings more efficiently, see the strings.Builder type.
func (b *Buffer) String() string {
    if b == nil {
        // Special case, useful in debugging.
        return "<nil>"
    }
    return string(b.buf[b.off:])
}
```
其String()方法就是将字节数组强转为string

Builder是如何实现的。
```
// Write appends the contents of p to b's buffer.
// Write always returns len(p), nil.
func (b *Builder) Write(p []byte) (int, error) {
    b.copyCheck()
    b.buf = append(b.buf, p...)
    return len(p), nil
}
```
Builder采用append的方式向字节数组后添加字符串。
```
// The append built-in function appends elements to the end of a slice. If
// it has sufficient capacity, the destination is resliced to accommodate the
// new elements. If it does not, a new underlying array will be allocated.
// Append returns the updated slice. It is therefore necessary to store the
// result of append, often in the variable holding the slice itself:
//  slice = append(slice, elem1, elem2)
//  slice = append(slice, anotherSlice...)
// As a special case, it is legal to append a string to a byte slice, like this:
//  slice = append([]byte("hello "), "world"...)
func append(slice []Type, elems ...Type) []Type
```

append源码
```
expand append(l1, l2...) to
//   init {
//     s := l1
//     n := len(s) + len(l2)
//     // Compare as uint so growslice can panic on overflow.
//     if uint(n) > uint(cap(s)) {
//       s = growslice(s, n)
//     }
//     s = s[:n]
//     memmove(&s[len(l1)], &l2[0], len(l2)*sizeof(T))
//   }
```

```
func main() {
    var sb = new(strings.Builder)
    sb.Write([]byte("12345678"))
    fmt.Printf("%s, %d\n",sb.String(),sb.Cap())
    sb.Write([]byte("9"))
    fmt.Printf("%s, %d\n",sb.String(),sb.Cap())
}
```

```
12345678, 8
123456789, 16
```
从上面可以看出，[]byte的内存大小也是以倍数进行申请的，初始大小为 0，第一次为大于当前申请的最大 2 的指数，不够进行翻倍.

append源码
```
if cap > doublecap {
        newcap = cap
    } else {
        if old.len < 1024 {
            newcap = doublecap
        } else {
            // Check 0 < newcap to detect overflow
            // and prevent an infinite loop.
            for 0 < newcap && newcap < cap {
                newcap += newcap / 4
            }
            // Set newcap to the requested cap when
            // the newcap calculation overflowed.
            if newcap <= 0 {
                newcap = cap
            }
        }
    }
```
可以看出如果旧容量小于1024进行翻倍，否则扩展四分之一。（2048 byte 后，申请策略的调整）。
```
// String returns the accumulated string.
func (b *Builder) String() string {
    return *(*string)(unsafe.Pointer(&b.buf))
}
```
其次String()方法与Buffer的string方法也有明显区别。Buffer的string是一种强转，我们知道在强转的时候是需要进行申请空间，并拷贝的。而Builder只是指针的转换。

这里我们解析一下*(*string)(unsafe.Pointer(&b.buf))这个语句的意思。

先来了解下unsafe.Pointer 的用法。

Pointer类型代表了任意一种类型的指针，类型Pointer有四种专属的操作：

任意类型的指针能够被转换成Pointer值
一个Pointer值能够被转换成任意类型的指针值
一个uintptr值能够被转换从Pointer值
一个Pointer值能够被转换成uintptr值

也就是说，unsafe.Pointer 可以转换为任意类型，那么意味着，通过unsafe.Pointer媒介，程序绕过类型系统，进行地址转换而不是拷贝。

即*A => Pointer => *B
```
func main() {
    var b = []byte{'H', 'E', 'L', 'L', 'O'}

    s := *(*string)(unsafe.Pointer(&b))

    fmt.Println("b =", b)
    fmt.Println("s =", s)

    b[1] = 'B'
    fmt.Println("s =", s)

    s = "WORLD"
    fmt.Println("b =", b)
    fmt.Println("s =", s)
    
    //b = [72 69 76 76 79]
    //s = HELLO
    //s = HBLLO
    //b = [72 66 76 76 79]
    //s = WORLD
}
```
就像上面例子一样，将字节数组转为unsafe.Pointer类型，再转为string类型，s和b中内容一样，修改b,s也变了，说明b和s是同一个地址。但是对s重新赋值后，意味着s的地址指向了“WORLD”,它们所使用的内存空间不同了，所以s改变后，b并不会改变。

所以他们的区别就在于 bytes.Buffer 是重新申请了一块空间，存放生成的string变量， 而strings.Builder直接将底层的[]byte转换成了string类型返回了回来，去掉了申请空间的操作。
