package main

import "fmt"

type IfA interface {
   Echo(s string)
}

type A struct {
}
func (* A) Echo(s string) {
  fmt.Println("echo:", s)
}

func main() {
	var a IfA
	a = &A{}
	a.Echo("Hello")
}
