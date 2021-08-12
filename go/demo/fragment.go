package main

import (
    //"os"
    "encoding/json"
    "fmt"
)

type Fragment struct {
    NormalLength   int   `json:"normalLength"`
    HdLength       int   `json:"hdLength"`
}

func main() {
    fragment := Fragment{
        NormalLength: 5000,
        //HdLength: 1000,
    }
    b, err := json.Marshal(fragment)
    if err != nil {
        fmt.Println("error:", err)
    }
    //os.Stdout.Write(b)
    fmt.Printf("%v", string(b))
}