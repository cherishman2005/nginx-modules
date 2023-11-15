# 捕捉异常

* `recover()用来捕捉本协程的panic`

```
package main

import (
	"fmt"
	"errors"
)

func test() (flag bool, err error) {
    defer func() {
	   if r := recover(); r != nil {
	       fmt.Println("recover=", r)
	   }
	   
	   fmt.Println("hh")
	   flag = false
	   err = errors.New("panic")
	}()
	
	err = errors.New("capture panic")
	panic(err)

	fmt.Println("err=", err)
	return true, err
}

func main() {
    flag, err := test()
	
    fmt.Println("flag=", flag, "err=", err)
}
```
