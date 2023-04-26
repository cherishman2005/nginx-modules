package main

import (
	"log"
)

type Request struct {
	// User context associated with this request
	Context map[interface{}]interface{}

}

func main() {
	req := Request{
		map[interface{}]interface{}{
			"appkey": 123456,
			"appsecret": "abcd",
			"code": 200,
		},
	}

	log.Print("req:", req)

	if v,ok := req.Context["code"]; ok {
		log.Print("code:", v)
	} else {
		log.Print("code not found")
	}

	if v,ok := req.Context["id"]; ok {
		log.Print("id:", v)
	} else {
		log.Print("id not found")
	}
}
