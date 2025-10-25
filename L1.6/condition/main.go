package main

import (
	"fmt"
	"sync/atomic"
)

// goroutine exit by condition
func main() {
	var done atomic.Bool
	c := make(chan int)
	go func() {
		cnt := 0
		for v := range c {
			if cnt < 10 {
				fmt.Println(v)
			}
			cnt++
			if cnt == 10 {
				done.Store(true)
			}
		}
	}()

	for i := 0; i < 100; i++ {
		if done.Load() {
			break
		}
		c <- i
	}
	close(c)
}
