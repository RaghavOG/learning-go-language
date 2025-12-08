package main

import (
	"fmt"
	"sync"
	"time"
)

// Run with: go run guide/examples/06_goroutines_basics.go
func main() {
	// Start a goroutine (lightweight concurrent function call)
	go say("fire-and-forget once") // will run, but main could exit first without sync

	// Use WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		sayWithDelay("worker A", 80*time.Millisecond)
	}()
	go func() {
		defer wg.Done()
		sayWithDelay("worker B", 30*time.Millisecond)
	}()
	wg.Wait() // wait for A and B

	// Channels move data between goroutines safely
	msgs := make(chan string) // unbuffered channel synchronizes send/recv
	go func() {
		msgs <- "ping"
	}()
	val := <-msgs // blocks until value is sent
	fmt.Println("got from channel:", val)

	// Buffered channel decouples send/recv up to capacity
	buf := make(chan int, 2)
	buf <- 1
	buf <- 2
	// buf <- 3 // would block if uncommented (buffer full)
	fmt.Println(<-buf, <-buf)

	// Closing a channel signals no more values
	data := make(chan int)
	go produceNumbers(data, 3)
	for n := range data { // loop ends when channel is closed
		fmt.Println("range recv:", n)
	}
}

func say(msg string) {
	fmt.Println(msg)
}

func sayWithDelay(msg string, d time.Duration) {
	time.Sleep(d)
	fmt.Println(msg)
}

// Sends 1..limit then closes channel
func produceNumbers(out chan<- int, limit int) {
	defer close(out)
	for i := 1; i <= limit; i++ {
		out <- i
	}
}

