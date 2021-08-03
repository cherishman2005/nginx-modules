

# Go语言首字母大小写问题

Go语言通过首字母的大小写来控制访问权限。

`无论是方法，变量，常量或是自定义的变量类型，如果首字母大写，则可以被外部包访问，反之则不可以。`

而结构体中的字段名，如果首字母小写的话，则该字段无法被外部包访问和解析，比如，json解析。

```
package main

import (
    //"os"
    "encoding/json"
    "fmt"
)

type colorGroup struct {
    ID     int
    name   string
    Colors []string
}
func main() {
    group := colorGroup{
        ID:     1,
        name:   "Redis",
        Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
    }
    b, err := json.Marshal(group)
    if err != nil {
        fmt.Println("error:", err)
    }
    //os.Stdout.Write(b)
    fmt.Printf("%s", b)
}
```

输出

```
{"ID":1,"Colors":["Crimson","Red","Ruby","Maroon"]}
```

如果希望，json化之后的属性名是小写字母的，可以使用struct tag。如下：

```
type colorGroup struct {
    ID     int        `json:"id"`
    Name   string     `json:"name"`
    Colors []string   `json:"colors"`
}
```

输出

```
{"id":1,"name":"Redis","colors":["Crimson","Red","Ruby","Maroon"]}
```

另外，struct tag还可以通过"-"，"omitempty"属性控制json的输出。

`如果没有特别的访问控制的话，建议字段名首字母都使用大写字母，从而避免无法解析的错误。`

# 参考链接

- [Go语言首字母大小写问题](https://www.jianshu.com/p/db16d3bd5908)