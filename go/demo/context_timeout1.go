package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	ch := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "data"
	}()

	select {
	case <-ctx.Done():
		return "timeout", ctx.Err() // 返回超时或取消错误
	case result := <-ch:
		return result, nil
	}
}

func main() {
	ret,err := fetchData(context.Background())
	fmt.Printf("result=%s, err=%v\n", ret, err)
}
