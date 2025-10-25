package main

import (
	"fmt"
	"runtime"
	"sync"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	n := 0
	for {
		if n == 10 {
			fmt.Println("goexit done")
			runtime.Goexit()
		}
		fmt.Println(n)
		n++
	}
}

// goroutine exit by goexit
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(&wg)
	wg.Wait()
}
