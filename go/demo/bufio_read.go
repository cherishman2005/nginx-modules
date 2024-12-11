package main

import (
    "bufio"
    "fmt"
    "io"
    "strings"
    "time"
)

//func main() {
//    //reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
//    reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com. It is the home of gophers"), 14)
//    line, err := reader.Peek(4)
//    fmt.Printf("err:%v, line:%s\n", err, line)
//    data := make([]byte, 30)
//    n, err := io.ReadFull(reader, data)
//    fmt.Printf("n:%d, err:%v, data:%v\n", n, err, string(data))
//
//    time.Sleep(5*time.Second)
//}

func main() {
    //reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
    reader := bufio.NewReaderSize(strings.NewReader("http"), 14)
    go func() {
        time.Sleep(2*time.Second)
        reader = bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
    }()
    for i:=0; i<4; i++ {
        line, err := reader.Peek(5)
        fmt.Printf("err:%v, line:%s\n", err, line)
        time.Sleep(1*time.Second)
    }

    frameHeader := make([]byte, 5)
    frameSize, err := io.ReadFull(reader, frameHeader)
    fmt.Printf("frameSize:%d, err:%v, data:%v\n", frameSize, err, string(frameHeader))

    data := make([]byte, 30)
    n, err := io.ReadFull(reader, data)
    fmt.Printf("n:%d, err:%v, data:%v\n", n, err, string(data))

    time.Sleep(5*time.Second)
}

func Peek(reader *bufio.Reader) {
    line, _ := reader.Peek(14)
    fmt.Printf("%s\n", line)
    // time.Sleep(1)
    fmt.Printf("%s\n", line)
}
