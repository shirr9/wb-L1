package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	n := 0
	timer := time.After(time.Second * 5)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Ctrl+C done")
			return
		case <-timer:
			fmt.Println("timeout")
			return
		case <-time.After(time.Second):
			fmt.Println(n)
			n++
		}
	}
}

// goroutine exit by syscall.SIGINT
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()
	var wg sync.WaitGroup

	fmt.Println("to stop click Ctrl+C or wait 5 sec for timeout")
	time.Sleep(time.Second)

	wg.Add(1)
	go worker(ctx, &wg)
	wg.Wait()

}
