package main

import (
	"fmt"
	"time"
)

// Run with: go run guide/examples/07_channels_patterns.go
func main() {
	unbufferedDemo()
	bufferedDemo()
	rangeAndCloseDemo()
	selectDemo()
	directionDemo()
}

// Unbuffered channel: send blocks until a receiver is ready.
func unbufferedDemo() {
	fmt.Println("\n-- unbuffered --")
	ch := make(chan string)
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- "ping"
	}()
	val := <-ch // blocks until send happens
	fmt.Println("got:", val)
}

// Buffered channel: send blocks only when full.
func bufferedDemo() {
	fmt.Println("\n-- buffered --")
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("recv:", <-ch, <-ch)
}

// Closing lets receivers know no more values are coming.
func rangeAndCloseDemo() {
	fmt.Println("\n-- close and range --")
	ch := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println("range recv:", v)
	}
}

// select picks the first ready case; handy for timeouts and multiplexing.
func selectDemo() {
	fmt.Println("\n-- select with timeout --")
	ch := make(chan string)
	go func() {
		time.Sleep(120 * time.Millisecond)
		ch <- "done"
	}()

	select {
	case v := <-ch:
		fmt.Println("got:", v)
	case <-time.After(80 * time.Millisecond):
		fmt.Println("timeout before channel was ready")
	}
}

// Directional channels document intent and prevent misuse.
func directionDemo() {
	fmt.Println("\n-- directional channels --")
	producer := func(out chan<- int) { // send-only
		defer close(out)
		for i := 1; i <= 2; i++ {
			out <- i
		}
	}
	consumer := func(in <-chan int) { // recv-only
		for v := range in {
			fmt.Println("consumed", v)
		}
	}

	ch := make(chan int, 2)
	go producer(ch)
	consumer(ch)
}

