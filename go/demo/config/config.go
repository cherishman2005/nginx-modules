package main

import (
	"fmt"
	"log"

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

func main() {
	var config ApiConfig

	err := yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	fmt.Println(config.Server.HttpPort)
}
