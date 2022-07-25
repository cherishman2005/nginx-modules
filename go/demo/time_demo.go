
package main
import (
	"fmt"
	"time"
)

func getSendTime() int64 {
	// 设定发送时间，多久聚合一次发送
	// 一天
	next := time.Now().Add(time.Hour*24)
    next = time.Date(next.Year(), next.Month(), next.Day(), 9, 0, 0, 0, next.Location())
	stime := next.Unix()

	return stime
}

func getSendTime1() int64 {
	// 设定发送时间，多久聚合一次发送
	// 一天
	//stime := time.Now().Add(time.Hour*24).Unix()
	//stime = stime - stime % int64(time.Hour.Seconds()*24) + int64(time.Hour*9)

	// 2小时
	stime := time.Now().Add(time.Hour*2).Unix()
	stime = stime - stime % int64(time.Hour.Seconds()*2)
	return stime
}


func main() {
    t := getSendTime()
    fmt.Println("24h=", t)
    
    t1 := getSendTime1()
    fmt.Println("2h=", t1)
}