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
