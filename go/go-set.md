# Golang中使用set

## 使用map实现

在Golang中通常使用map来实现set，map中的key为唯一值，这与set的特性一致。
简单实现，如下：
```
set := make(map[string]bool) // New empty set
set["Foo"] = true            // Add
for k := range set {         // Loop
    fmt.Println(k)
}
delete(set, "Foo")    // Delete
size := len(set)      // Size
exists := set["Foo"]  // Membership
```
map的value值是布尔型，这会导致set多占用内存空间，解决这个问题，则可以将其替换为空结构。在Go中，空结构通常不使用任何内存。

```
unsafe.Sizeof(struct{}{}) // 结果为 0
```

优化后，如下：
```
type void struct{}
var member void

set := make(map[string]void) // New empty set
set["Foo"] = member          // Add
for k := range set {         // Loop
    fmt.Println(k)
}
delete(set, "Foo")      // Delete
size := len(set)        // Size
_, exists := set["Foo"] // Membership
```

## golang-set

golang-set-A simple set type for the Go language. Also used by Docker, 1Password, Ethereum.
在github上已经有了一个成熟的包，名为golang-set，包中提供了线程安全和非线程安全的set。提供了五个set函数：
```
// NewSet创建并返回空集的引用，结果集上的操作是线程安全的
func NewSet(s ...interface{}) Set {}
// NewSetFromSlice从现有切片创建并返回集合的引用，结果集上的操作是线程安全的
func NewSetFromSlice(s []interface{}) Set {}
// NewSetWith创建并返回具有给定元素的新集合，结果集上的操作是线程安全的
func NewSetWith(elts ...interface{}) Set {}
// NewThreadUnsafeSet创建并返回对空集的引用，结果集上的操作是非线程安全的
func NewThreadUnsafeSet() Set {}
// NewThreadUnsafeSetFromSlice创建并返回对现有切片中集合的引用，结果集上的操作是非线程安全的。
func NewThreadUnsafeSetFromSlice(s []interface{}) Set {}
```
简单案例，如下：
```
package main

import (
    "fmt"
    "github.com/deckarep/golang-set"
)

func main() {
    // 默认创建的线程安全的，如果无需线程安全
    // 可以使用 NewThreadUnsafeSet 创建，使用方法都是一样的。
    s1 := mapset.NewSet(1, 2, 3, 4)
    fmt.Println("s1 contains 3: ", s1.Contains(3))
    fmt.Println("s1 contains 5: ", s1.Contains(5))

    // interface 参数，可以传递任意类型
    s1.Add("poloxue")
    fmt.Println("s1 contains poloxue: ", s1.Contains("poloxue"))
    s1.Remove(3)
    fmt.Println("s1 contains 3: ", s1.Contains(3))

    s2 := mapset.NewSet(1, 3, 4, 5)

    // 并集
    fmt.Println(s1.Union(s2))
}
```
结果为：
```
s1 contains 3:  true
s1 contains 5:  false
s1 contains poloxue:  true
s1 contains 3:  false
Set{1, 2, 4, poloxue, 3, 5}
```
