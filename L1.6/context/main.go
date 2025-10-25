package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	n := 0
	for {
		select {
		case <-ctx.Done():
			time.Sleep(time.Millisecond)
			fmt.Println("context done")
			return
		default:
			fmt.Println(n)
			n++
			time.Sleep(time.Millisecond)
		}
	}
}

// goroutine exit via context
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(ctx, &wg)
	wg.Wait()
}
