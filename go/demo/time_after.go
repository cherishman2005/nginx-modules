package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	<-time.After(5*time.Second)
	fmt.Println(time.Now().Unix())
}
