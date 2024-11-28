package main

import (
    "fmt"
    "log"
    "time"
)

func run(task_id, sleeptime int, ch chan string) {

    time.Sleep(time.Duration(sleeptime) * time.Second)
    ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
    return
}

func main() {
    input := []int{3, 2, 10, 1}
    ch := make(chan string)
    startTime := time.Now()
    log.Println("Multirun start")
    for i, sleeptime := range input {
        go run(i, sleeptime, ch)
    }

    for range input {
        log.Println(<-ch)
    }

    endTime := time.Now()
    log.Printf("Multissh finished. Process time %s. Number of tasks is %d", endTime.Sub(startTime), len(input))
}
