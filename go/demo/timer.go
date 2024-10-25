package main

import (
	"log"
	"time"
)

func timerFun() {
	timer := time.NewTimer(3 * time.Second)

	log.Print("Timer start")
	// 模拟在定时器触发前就跳出模块
	return

	select {
	case <-timer.C:
		log.Print("Timer triggered")
	default:
		log.Print("Timer not triggered")
	}

	log.Print("return")
}

func main() {
	log.Print("main start")
	go timerFun()
	time.Sleep(10*time.Second)

	log.Print("main exit")
}
