package main

import (
	"fmt"
	"time"
)


func main() {
		ch := make(chan bool)
		ch1 := make(chan bool)
		//close(ch)
		//ch <- true
		
    go func() {
			select {
			case p := <-ch:
				fmt.Println("ok, p=", p)
			case p1 := <-ch1:
				fmt.Println("ok, p1=", p1)	
//			default:
//				fmt.Println("default")
			}
			fmt.Println("select end")
			
	}()

    ch1 <- true
    time.Sleep(10*time.Second)
	
	fmt.Println("end")
}
