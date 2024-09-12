package main

import (
	"fmt"
)



func main() {
	flag := -1
	level := 0
	switch level {
	case 1:
		flag = 10
	case 2:
		flag = 100
	}
	fmt.Println("flag=", flag)
}
