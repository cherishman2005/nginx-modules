/*
Race method returns a response as soon as one of the callbacks in an iterable resolves with the value that is not an error, otherwise last error is returne
 */
package main

import (
	"context"
	"errors"
	"log"
	"time"
	"github.com/vardius/gollback"
)

func main() {
	log.Printf("main start")
	r, err := gollback.Race(
		context.Background(),
		func(ctx context.Context) (interface{}, error) {
			time.Sleep(3 * time.Second)
			return 1, nil
		},
		func(ctx context.Context) (interface{}, error) {
			return nil, errors.New("failed")
		},
		func(ctx context.Context) (interface{}, error) {
			return 3, nil
		},
	)
	log.Printf("r=%v, err=%v", r, err)
	log.Printf("main end")
}
