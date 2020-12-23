package main
 
import (
    "encoding/json"
    "fmt"
)
 
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    Addr string `json:"addr,omitempty"`
    //Addr string `json:"addr"`
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