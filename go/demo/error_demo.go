//error_demo.go
package main


import "fmt"
import "strconv"


func foo(x string) (err error) {
    //if true {
        ret, err := strconv.Atoi(x)
        if err != nil {
            return
        }
    //}
    fmt.Println("ret:", ret)
    return nil
}


func main() {
    fmt.Println("err:", foo("a123"))
}
