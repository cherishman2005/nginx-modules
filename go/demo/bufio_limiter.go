package main

import (
        "fmt"
        "io"
        "strings"
)

func main() {
        // 创建一个字符串作为底层读取器
        reader := strings.NewReader("Hello, World!")

        // 创建一个LimitedReader，限制最多读取5个字节
        limitReader := io.LimitReader(reader, 5)

        // 使用LimitedReader进行读取操作
        buffer := make([]byte, 10)
        n, err := limitReader.Read(buffer)

        if err != nil && err != io.EOF {
                fmt.Println("读取错误:", err)
                return
        }

        fmt.Println("读取的字节数:", n)
        fmt.Println("读取的内容:", string(buffer[:n]))
}
