package main

import (
    "fmt"
)

type AuthKey struct {
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
    configs := make(map[*AuthKey]*AuthConfig)
    
    key := &AuthKey{
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
    
    
    
    key1 := &AuthKey{
        Appid: 15013,
        BusinessAuth: 16777218,
    }
    config1 := &AuthConfig{
        Appid: 15013,
        BusinessAuth: 16777218,
        Srvname: "auth_audio_route",
        Pass: 0,
    }
    configs[key1] = config1
    
    fmt.Println("configs:", configs)
    
    if v, ok := configs[key]; ok {
        fmt.Println(v)
    } else {
        fmt.Println("not found")
    }
    
    
    key2 := &AuthKey{
        Appid: 15013,
        BusinessAuth: 16777218,
    }
    if v, ok := configs[key2]; ok {
        fmt.Println(v)
    } else {
        fmt.Println("not found")
    }
}

/*
    运行结果：
    configs: map[0xc0000140b8:0xc00000c060 0xc0000140d0:0xc00000c080]
    &{22014 16777218 auth_audio_route 0}
    not found
*/
