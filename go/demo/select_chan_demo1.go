package main

import (
	"fmt"
	"time"
)

type persistConn struct {
	cacheKey string     
}

var persistConnClosedCh = make(chan *persistConn)

func init() {
	close(persistConnClosedCh)
}

func getIdleConnCh() chan *persistConn {
    //ch := make(chan *persistConn)
	//return ch
	
	//return persistConnClosedCh
	return nil
}

func main() {
	//ch := getIdleConnCh()
	//close(ch)
		
	ch1 :=  make(chan bool)
	time.Sleep(1*time.Second)

	//ch <- true
	
	idleConnCh := make(map[string]chan *persistConn)
	key := "123"
	
	waitingDialer := idleConnCh[key]
		
    go func() {
			select {
			case p := <-waitingDialer:
				fmt.Println("ok, p=", p)
			case p1 := <-ch1:
				fmt.Println("ok, p1=", p1)	
			default:
				fmt.Println("default")
			}
			fmt.Println("select end")
			
	}()
	
	pc := &persistConn{cacheKey:"123"}
    waitingDialer <- pc
	
    time.Sleep(10*time.Second)
	
	fmt.Println("end")
}
