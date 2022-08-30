package main

import (
        "unsafe"
        "fmt"
)

func main() {
        var a int
        var b int16
        var c rune
        var d bool
        fmt.Printf("%v\t%v\t%v\n", &a,unsafe.Sizeof(a),unsafe.Alignof(a))
        fmt.Printf("%v\t%v\t%v\n", &b,unsafe.Sizeof(b),unsafe.Alignof(b))
        fmt.Printf("%v\t%v\t%v\n", &c,unsafe.Sizeof(c),unsafe.Alignof(c))
        fmt.Printf("%v\t%v\t%v\n", &d,unsafe.Sizeof(d),unsafe.Alignof(d))

        var m map[string]int  = make(map[string]int)
        fmt.Printf("%v\t%v\t%v\n", &m,unsafe.Sizeof(m),unsafe.Alignof(m))
        
        var mb map[string]bool  = make(map[string]bool)
        fmt.Printf("%v\t%v\t%v\n", &mb,unsafe.Sizeof(mb),unsafe.Alignof(mb))
}
