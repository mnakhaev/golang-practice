package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	resultCh := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	services := []string{"Taxi A", "Taxi B", "Taxi C", "Taxi D"}

	var winner string
	wg := sync.WaitGroup{}

	for i := range services {
		svc := services[i]

		wg.Add(1)
		go func() {
			orderTaxi(ctx, svc, resultCh)
			wg.Done()
		}()

		go func() {
			winner = <-resultCh
			cancel() // now need to abort all other orders since we have the winner
		}()
	}
	wg.Wait()
	log.Printf("found taxi in %q", winner)
}

func orderTaxi(ctx context.Context, svc string, resultCh chan string) {
	time.Sleep(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			log.Printf("stopped search in service name %q with message = %v", svc, ctx.Err())
			return
		default:
			if rand.Float64() > 0.75 {
				resultCh <- svc
				return
			}
		}
	}
}
