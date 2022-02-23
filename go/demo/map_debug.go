package main

import (
    "fmt"
)



func main() {
    sids := map[uint32]int8{
	    123: 1,
		456: 0,
	}
	
	if v, ok := sids[123]; ok {
	    fmt.Println("value=", v)
	} else {
	    fmt.Println("not found")
	}

	if v, ok := sids[333]; ok {
	    fmt.Println("value=", v)
	} else {
	    fmt.Println("not found")
	}
	
	if sids[333] == 1 {
	    fmt.Println("ok")
	} else {
	    fmt.Println("not found")
	}
}

/*
    运行结果：

*/
