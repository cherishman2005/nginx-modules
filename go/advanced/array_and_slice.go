package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}

	sc := s[1:3]
	s1 := s
	sc[0] = 100
	fmt.Println("main sc=",sc)
	fmt.Println("main s=",s)
	fmt.Println("main s1=",s1)

	str := "abc"
	strb := []byte(str)
	strb[0] = 'k'
	fmt.Println("main str=", str)
	fmt.Println("main strb=", string(strb))
}
