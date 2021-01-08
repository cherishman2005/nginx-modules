package main
 
import (
	"fmt"
	"strings"
)
 
func main() {
	demo := "I&love&Go,&and&I&also&love&Python."
	string_slice := strings.Split(demo, "&")
 
	fmt.Println("result:",string_slice)
	fmt.Println("len:",len(string_slice))
	fmt.Println("cap:", cap(string_slice))
}