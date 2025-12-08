package main

import (
	"fmt"
	"sync"
	"time"
)

// Run with: go run guide/examples/05_concurrency.go
func main() {
	// Goroutine + WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("worker 1 start")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("worker 1 done")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("worker 2 start")
		time.Sleep(50 * time.Millisecond)
		fmt.Println("worker 2 done")
	}()
	wg.Wait()

	// Channels and select
	msgs := make(chan string, 1) // buffered
	msgs <- "hello"
	fmt.Println(<-msgs)

	data := make(chan int)
	quit := make(chan struct{})
	go producer(data, quit)

	for {
		select {
		case n := <-data:
			fmt.Println("got", n)
			if n >= 3 {
				close(quit)
				return
			}
		case <-time.After(200 * time.Millisecond):
			fmt.Println("timeout")
		}
	}
}

func producer(out chan<- int, quit <-chan struct{}) {
	defer close(out)
	for i := 1; ; i++ {
		select {
		case <-quit:
			return
		case out <- i:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

