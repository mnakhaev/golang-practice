package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	doWork(ctx)
}

func doWork(ctx context.Context) {
	// although we redefined context, timeout will be still taken from parent context
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			// will be executed after timeout in parent context
			log.Printf("new CTX was done: %v", ctx.Err())
			return
		default:
			log.Println("working")
			time.Sleep(1 * time.Second)
		}
	}

}
