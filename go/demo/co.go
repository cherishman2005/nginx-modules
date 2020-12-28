package main

import (
  "fmt"
  "time"
)

func deferPrint(str string){
  time.Sleep(time.Second*2)
  fmt.Println(str)
}

func main(){
  go deferPrint("hello")
  go deferPrint("world")
  //如果主协程不阻塞，永远不会切换

  time.Sleep(time.Second*4)
}
