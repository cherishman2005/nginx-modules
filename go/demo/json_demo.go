package main

import (
    "fmt"
    //"strings"
    "encoding/json"
)

type BaseJson struct {
    Auth string  `json:"auth,omitempty"`
    Uid  *uint64 `json:"uid,omitempty"`
}

func main() {
    //uid := uint64(2345678)
    req := BaseJson{
        Auth: "abc123",
        //Uid: &uid,
    }

    jsonData, err := json.Marshal(req)
    if err != nil {
        fmt.Println("encode json failed")
        return
    }
    
    fmt.Println("jsonData=", jsonData)
    
    res := &BaseJson{}
    err = json.Unmarshal([]byte(jsonData), &res)
    if err != nil {
        fmt.Println("decode json failed")
        return
    }
    if res.Uid != nil {
        fmt.Println("res=", res, " uid=", *res.Uid)
    }
}