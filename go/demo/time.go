package main

import (
    "fmt"
    "time"
)

// 获取年月日
func DateYmdFormat() string {
    tm := time.Now()
    return tm.Format("20060102")
}

func main() {
    timeStr := time.Now().Format("2006-01-02 15:04:05") //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
    fmt.Println(timeStr)                                //打印结果：2017-04-11 13:24:04

    fmt.Println(DateYmdFormat())
}

