package main

import (
    "bufio"
    "fmt"
    "io"
    "strings"
    "time"
)

func main() {
    //reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
    reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com. It is the home of gophers"), 14)
    line, err := reader.Peek(4)
    fmt.Printf("err:%v, line:%s\n", err, line)
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
