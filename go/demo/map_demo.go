package main

import (
    "fmt"
    "errors"
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

func getAuthRouteConfig(authConfigs map[RouteKey]*AuthConfig, appid, auth uint32) (*AuthConfig, error) {
	key := RouteKey{
		Appid: appid,
		BusinessAuth: auth,
	}

	if v, ok := authConfigs[key]; ok {
		return v, nil
	}
	
	return nil, errors.New("no route config")
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
    
    
    
    key1 := RouteKey{
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
    
    
    key2 := RouteKey{
        Appid: 15013,
        BusinessAuth: 16777218,
    }
    if v, ok := configs[key2]; ok {
        fmt.Println(v)
    } else {
        fmt.Println("not found")
    }
    
    if conf, err := getAuthRouteConfig(configs, 15013, 16777218); err == nil{
        fmt.Println(conf)
    } else {
        fmt.Println("not found")
    }
}

/*
    运行结果：
    configs: map[{15013 16777218}:0xc00000c080 {22014 16777218}:0xc00000c060]
    &{22014 16777218 auth_audio_route 0}
    &{15013 16777218 auth_audio_route 0}
    &{15013 16777218 auth_audio_route 0}
*/