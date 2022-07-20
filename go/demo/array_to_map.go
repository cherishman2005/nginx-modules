package main

import (
    "fmt"
    "strings"
)



func transferStrArrayToMap(arr string) (map[string]bool) {
    a := strings.Split(arr, ",")
    m := make(map[string]bool)
    
    for _, v := range a {
        e := strings.TrimSpace(v)
        if len(e) == 0 {
            continue
        }
        
        m[e] = true
    }

    return m
}

func main() {
    m := transferStrArrayToMap("zhangbiwu, zhangbiwu, zhang    ")
    fmt.Println(m)
}