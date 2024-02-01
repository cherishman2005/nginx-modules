package main

import (
	"math/rand"
	"fmt"
	"time"
)


func main() {
    sleep := time.Second
	// Add some randomness to prevent creating a Thundering Herd
	jitter := time.Duration(rand.Int63n(int64(sleep)))
	fmt.Println(jitter)
}
