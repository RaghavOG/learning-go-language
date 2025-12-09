package main

import (
	"fmt"
	"time"
)

// helper to visualize buffer state after every send/receive
func printBufferState(ch chan int, label string) {
	fmt.Printf("[%s] Buffer size: %d / %d\n", label, len(ch), cap(ch))
}

func main() {
	ch := make(chan int, 3) // BUFFERED CHANNEL OF SIZE 3

	// SENDER 1
	go func() {
		fmt.Println("G1: Sending 100...")
		ch <- 100
		fmt.Println("G1: Sent 100!")
	}()

	// SENDER 2
	go func() {
		fmt.Println("G2: Sending 200...")
		ch <- 200
		fmt.Println("G2: Sent 200!")
	}()

	// SENDER 3
	go func() {
		fmt.Println("G3: Sending 300...")
		ch <- 300
		fmt.Println("G3: Sent 300!")
	}()

	// SENDER 4
	go func() {
		fmt.Println("G4: Sending 400 (will block if buffer full)...")
		ch <- 400
		fmt.Println("G4: Sent 400! (only prints after a receive happens)")
	}()

	// Let senders run
	time.Sleep(1 * time.Second)

	// Show buffer state BEFORE receiving
	printBufferState(ch, "Before receiving")

	// Now start receiving
	fmt.Println("\n--- RECEIVING VALUES ---")
	val := <-ch
	fmt.Println("Main: Received:", val)
	printBufferState(ch, "After receiving 1")

	time.Sleep(1 * time.Second)

	// Uncomment these to drain more and see effect
	
	val = <-ch
	fmt.Println("Main: Received:", val)
	printBufferState(ch, "After receiving 2")

	val = <-ch
	fmt.Println("Main: Received:", val)
	printBufferState(ch, "After receiving 3")

	val = <-ch
	fmt.Println("Main: Received:", val)
	printBufferState(ch, "After receiving 4")
	

	time.Sleep(1 * time.Second)
	fmt.Println("\nProgram End")
}
