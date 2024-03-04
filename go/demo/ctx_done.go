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
	   //ch <- 1
   }()

   ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
   defer cancel()
   
   ch <- 1
   cancel()
   go func() {
	   var closeTicker <-chan time.Time
	   for {
		   select {
		   case <-ctx.Done():
			   closeTicker = time.After(1 * time.Millisecond)
			   fmt.Println("timeout:", ctx.Err())
			   default:
           }
		   select {
			   case i := <-ch:
			         closeTicker = time.After(2 * time.Millisecond)
					 fmt.Println("i=", i)
				case <- closeTicker:
				     fmt.Println("closeTicker:", ctx.Err())

				     return
		   }
	   }
   }()
   
   fmt.Println("select waiting")
   
   //select {
	//   case <-ctx.Done():
	//      	fmt.Println("timeout222:", ctx.Err())
	//   //case i := <-ch:
	//   //     fmt.Println("i=", i)
   //}
   //
   //fmt.Println("select end")
   time.Sleep(10*time.Second)
   fmt.Println("end")
}
