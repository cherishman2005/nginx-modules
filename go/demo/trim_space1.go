package main  
  
import (
    "fmt"
    "strings"
)

func transferStringToArray(s string) []string {
    var arr []string
    elements := strings.Split(s, ",")

    for _, v := range elements {
        t := strings.TrimSpace(v)
        arr = append(arr, t)
    }
    return arr
}

func main() {
    s := "\t hello ,  world   "
    arr := transferStringToArray(s)

    fmt.Printf("arr:%v", arr)
}
