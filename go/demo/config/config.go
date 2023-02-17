package main

import (
    "fmt"
    "log"
    "os"
    "gopkg.in/yaml.v3"
)

type ConfigBasic struct {
    HttpPort       int  `yaml:"http_port"`// listen port for http
}

type ApiConfig struct {
    // basic server config
    Server ConfigBasic  `yaml:"server"`
}

var data = `
server:
    # listen port for http request
    http_port: 8080
`

func ReadManifest() (*ApiConfig, error) {
    p := "./api.yml"

    file, err := os.Open(p)
    if err != nil {
        return nil, fmt.Errorf("failed to open the plugin manifest %s: %w", p, err)
    }

    defer func() { _ = file.Close() }()

    m := &ApiConfig{}
    err = yaml.NewDecoder(file).Decode(m)
    if err != nil {
        return nil, fmt.Errorf("failed to decode the plugin manifest %s: %w", p, err)
    }

    return m, nil
}

func main() {
    var config ApiConfig

    err := yaml.Unmarshal([]byte(data), &config)
    if err != nil {
        log.Fatalf("cannot unmarshal data: %v", err)
    }
    fmt.Println(config.Server.HttpPort)
    

    config1, err = ReadManifest()
    if err != nil {
        log.Fatalf("cannot unmarshal data: %v", err)
    }
    fmt.Println(config1.Server.HttpPort)
    
}
