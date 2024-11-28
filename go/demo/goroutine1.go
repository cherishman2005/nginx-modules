package main

import (
    "fmt"
    "log"
    "runtime/debug"
    "sync"
    "time"
    "context"
)

var wg = sync.WaitGroup{}

func run(task_id, sleeptime int, ch chan string) {
    defer func() {
        wg.Done()
        if r := recover(); r != nil {
            log.Printf("Recovered from panic: %v\n%s\n", r, debug.Stack())
        }
    }()
    time.Sleep(time.Duration(sleeptime) * time.Second)
    ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
    return
}

func multirun() {
    input := []int{3, 2, 10, 1}
    ch := make(chan string, len(input))
    startTime := time.Now()
    log.Println("Multirun start")

    ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
    defer func() {
        cancel()
        for range input {
            log.Println("time is up:", <-ch)
        }
    }()

    wg.Add(len(input))
    for i, sleeptime := range input {
        go run(i, sleeptime, ch)
    }
    wg.Wait()
    //close(ch)
    for range input {
        log.Println(<-ch)
    }

    select {
    case <- ctx.Done():
      log.Printf("multirun ctx err: %v", ctx.Err())
      return
    }

    endTime := time.Now()
    log.Printf("Multissh finished. Process time %s. Number of tasks is %d", endTime.Sub(startTime), len(input))
}

func main() {
    log.Println("main start")
    go multirun()
    time.Sleep(15 * time.Second)
    log.Println("main end")
}
