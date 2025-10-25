package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	n := 0
	for {
		select {
		case <-done:
			fmt.Println("exit")
			return
		default:
			fmt.Println(n)
			n++
			time.Sleep(time.Millisecond)
		}
	}
}

// goroutine exit via notification channel
func main() {
	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(done, &wg)
	time.Sleep(10 * time.Millisecond)
	close(done)
	wg.Wait()
}
