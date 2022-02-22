package main

import (
    "fmt"
)

type RouteKey struct {
	Appid uint32
	BusinessAuth uint32
}

type AuthConfig struct {
	Appid uint32
	BusinessAuth uint32
	Srvname string
	Pass uint32
}




func main() {
    configs := make(map[RouteKey]*AuthConfig)
    
    key := RouteKey{
        Appid: 22014,
        BusinessAuth: 16777218,
    }

    config := &AuthConfig{
        Appid: 22014,
        BusinessAuth: 16777218,
        Srvname: "auth_audio_route",
        Pass: 0,
    }
    
    configs[key] = config
    
    if v, ok := configs[key]; ok {
        fmt.Println(v)
    } else {
        fmt.Println("not found")
    }
    
    
    key1 := RouteKey{
        Appid: 15013,
        BusinessAuth: 16777218,
    }
    if v, ok := configs[key1]; ok {
        fmt.Println(v)
    } else {
        fmt.Println("not found")
    }
}