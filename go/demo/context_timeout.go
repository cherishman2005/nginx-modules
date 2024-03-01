package main

import (
	"fmt"
	"time"
	"context"
)


func main() {
   ch := make(chan int, 100)
   go func() {
       time.Sleep(3*time.Second)
	   ch <- 1
   }()
   
   ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
   defer cancel()
   select {
       case <-time.After(2*time.Second):
	        fmt.Println("time.After timeout")
	   case <-ctx.Done():
	      	fmt.Println("timeout:", ctx.Err())
	   case i := <-ch:
	        fmt.Println("i=", i)
   }

   fmt.Println("select end")
   time.Sleep(10*time.Second)
   fmt.Println("end")
}
