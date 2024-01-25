// Golang program to illustrate the usage of
// io.ReadFull() function

// Including main package
package main

// Importing fmt, io, and strings
import (
	"fmt"
	"io"
	"strings"
)

// Calling main
func main() {

	// Defining reader using NewReader method
	reader := strings.NewReader("Geeks")

	// Defining buffer of specified length
	// using make keyword
	buffer := make([]byte, 4)

	// Calling ReadFull method with its parameters
	n, err := io.ReadFull(reader, buffer)

	// If error is not nil then panics
	if err != nil {
		panic(err)
	}

	// Prints output
	fmt.Printf("Number of bytes in the buffer: %d\n", n)
	fmt.Printf("Content in buffer: %s\n", buffer)
	//此时只剩s没有读取，即大小为1，当设置大于1时报错：panic: unexpected EOF
	buffer2 := make([]byte, 1)
	n, err = io.ReadFull(reader, buffer2)
	// If error is not nil then panics
	if err != nil {
		panic(err)
	}

	// Prints output
	fmt.Printf("Number of bytes in the buffer: %d\n", n)
	fmt.Printf("Content in buffer: %s\n", buffer2)

}
