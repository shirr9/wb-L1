package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic: ", r)
		}
	}()
	defer wg.Done()

	n := 0
	for {
		if n == 10 {
			panic("n became 10")
		}
		fmt.Println(n)
		n++
	}
}

// goroutine exit by panic
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(&wg)
	wg.Wait()
}
