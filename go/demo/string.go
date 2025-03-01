// lru
package main

import (
	"fmt"
)

func main() {
	str := "hello"
	// 尝试修改字符串中的字符，这会导致编译错误
	//str[0] = 'H' // 编译错误：cannot assign to str[0] (strings are immutable)

	fmt.Println(str)

	strb := []rune(str)
	strb[0] = 'z'
	fmt.Printf("strb=%s\n", string(strb))
}

