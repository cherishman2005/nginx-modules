package main

import (
	"fmt"
	"time"
)

func main() {
	//loc, _ := time.LoadLocation("UTC")

	//t := time.Now().UTC()
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	t := time.Now().In(loc).Hour()
	d := time.Now().Hour()
	fmt.Println(t, d)



	loc, err = time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	t = time.Now().In(loc).Hour()
	d = time.Now().Hour()

	fmt.Println(t, d)
}
