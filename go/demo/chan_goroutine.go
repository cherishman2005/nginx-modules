package main

import (
	"fmt"
	"time"
)


func main() {
   start := make(chan bool, 1)
   ch := make(chan int, 100)
   go func() {
       cnt := 0
       <- start
	   for {
	      select {
			  case i,ok := <-ch:
			      fmt.Println("i=", i, "ok=", ok)
			  default:
			      cnt++
			      fmt.Println("default cnt=", cnt)
		  }
	   }
   }()
   
   for i := 0; i < 10; i++ {
       ch <- i
	   if i == 5 {
	       close(ch)
		   ch = nil
		   start <- true
		   break
	   }
   }
   
   time.Sleep(10*time.Second)
   fmt.Println("end")

}
