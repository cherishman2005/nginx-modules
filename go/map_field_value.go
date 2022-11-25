package main

import "fmt"

type ActiveClient struct {
   IpPort string
}

func main() {
    m := map[int]*ActiveClient {
	     1:&ActiveClient{"127.0.0.1:80"},
		 2:&ActiveClient{"127.0.0.1:88"},
	    }
	if v,ok := m[1]; ok {
	    v.IpPort = "127.0.0.1:99"
	}
	for k, v := range m {
	    fmt.Println(k, v)
	}
	
	mi := map[int]int{1:1, 2:2}
	if _,ok := mi[1]; ok {
	    mi[1] = 3
	}
	fmt.Println(mi)
	
}

/*
1 &{127.0.0.1:99}
2 &{127.0.0.1:88}
map[1:3 2:2]
*/
