package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(timer *time.Timer, wg *sync.WaitGroup) {
	defer wg.Done()
	n := 0
	for {
		select {
		case <-timer.C:
			fmt.Println("timer done")
			return
		default:
			fmt.Println(n)
			n++
			time.Sleep(time.Millisecond)
		}
	}
}

// goroutine exit via timer
func main() {
	duration := 1 * time.Second
	timer := time.NewTimer(duration)
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(timer, &wg)

	wg.Wait()
}
