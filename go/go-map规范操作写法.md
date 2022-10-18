# go-map规范操作写法

```
g_IfaceMap         map[string]*IfaceInfo
```

写法1：
```
	//if len(this.g_IfaceMap) == 0 || this.g_IfaceMap[ifName] == nil {
```

写法2：
```
	if _, ok := this.g_IfaceMap[ifName]; !ok {
```


写法2 是不是更规范？ 哪个性能好？

## 点评

写法2 规范，性能差别不大

## 示例代码

```
package main 
  
import ( 
    "fmt"
) 
  
func main() { 
    var keyvalue map[string]string = nil
    k := "key"
    if _, ok := keyvalue[k]; ok {
        fmt.Println("ok")
    } else {
        fmt.Println("not found")
    }
    
    if len(keyvalue) != 0 || len(keyvalue[k]) != 0 {
        fmt.Println("ok")
    } else {
        fmt.Println("not found")
    }
}
```

**map为nil也可以查询Find**
