package main

import (
	"fmt"
	//"strings"
)

type Req struct {
	u32Params []int
}

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)

		req := &Req{
			u32Params: []int {8, 8, 8},
		}
		s1 = append(s1, req.u32Params...)
		fmt.Println(s1)

		s3 := []int{}
		s3 = append(s3, s2...)
	fmt.Println(s3)
}

/*
运行结果
[1 2 3 4 5 8 8 8]
[4 5]
*/