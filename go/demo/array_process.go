package main

import (
    "fmt"
    "strings"
)


func main() {
    RName := "./demo   111"
    rName := ""
    i := strings.Index(RName, " ")
    if i > 0 {
        rName = RName[:i]
    }
    
    fmt.Println("i:", i, "RName:", RName)
    fmt.Println("rName:", rName, "len:", len(rName))
}
