
```
package main
 
import (
    "encoding/json"
    "fmt"
)
 
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    Addr string `json:"addr,omitempty"`
}
 
func main() {
    p1 := Person{
        Name: "taoge",
        Age:  30,
    }
 
    data, err := json.Marshal(p1)
    if err != nil {
        panic(err)
    }
 
    fmt.Printf("%s\n", data)
    fmt.Println(p1.Name, p1.Age, p1.Addr)
 
    p2 := Person{
        Name: "Cang Laoshi",
        Age:  18,
        Addr: "Japan",
    }
 
    data2, err := json.Marshal(p2)
    if err != nil {
        panic(err)
    }
 
    fmt.Printf("%s\n", data2)
    
    fmt.Println(p2.Name, p2.Age, p2.Addr)
}
```

结果：

```
{"name":"taoge","age":30}
taoge 30 
{"name":"Cang Laoshi","age":18,"addr":"Japan"}
Cang Laoshi 18 Japan
```

可以看到，有了omitempty后，如果addr为空， 则生成的json中没有addr字段。

可以去掉omitempty, 再试试。