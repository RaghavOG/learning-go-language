package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan string)

	// Simulate two servers with different latencies.
	go func() {
		time.Sleep(1 * time.Second)
		result <- "Server A"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		result <- "Server B"
	}()

	select {
	case res := <-result:
		fmt.Println("Fastest server:", res)
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("Timeout occurred")
	}
}

