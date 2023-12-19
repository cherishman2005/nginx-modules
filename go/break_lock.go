// 错误代码
package main

import (
	"fmt"
	"sync"
)

var mu  sync.RWMutex

func main() {
    fmt.Println("start")
    for i := 0; i < 10; i++ {
	    mu.Lock()
	    fmt.Println("aa i=", i)
		
		if i== 5 {
		    fmt.Println("aaaaaaaaa i=", i)
		    break;
		}
		fmt.Println("bb i=", i)
		
		mu.Unlock()
		
	}
	fmt.Println("end")
}

/*
start
aa i= 0
bb i= 0
aa i= 1
bb i= 1
aa i= 2
bb i= 2
aa i= 3
bb i= 3
aa i= 4
bb i= 4
aa i= 5
aaaaaaaaa i= 5
end
*/
