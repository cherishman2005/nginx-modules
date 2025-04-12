package main

import (
	"fmt"
	"reflect"
)

/*
如何判断 2 个字符串切片（slice) 是相等的？
reflect.DeepEqual() ， 但反射非常影响性能。
 */

type Author struct {
	Name         int      `json:Name`
	Publications []string `json:Publication,omitempty`
}

func main() {
	t := reflect.TypeOf(Author{})
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		s, _ := t.FieldByName(name)
		fmt.Println(name, s.Tag)
	}

	a := "abc"
	b := "abc"

	aa := []byte(a)
	bb := []byte(b)
	fmt.Println("a is equal to b:", reflect.DeepEqual(aa, bb))
}
