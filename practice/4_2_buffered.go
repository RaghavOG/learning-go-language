package main

import (
	"fmt"
	"time"
)

func printBuffer(ch chan int, label string) {
	fmt.Printf("\n[%s]\n", label)
	fmt.Printf("Buffer (%d/%d): ", len(ch), cap(ch))

	// Drain ONLY up to capacity items safely
	tmp := make([]int, 0, cap(ch))
	count := len(ch)

	for i := 0; i < count; i++ {
		val := <-ch
		tmp = append(tmp, val)
		fmt.Printf("[%d] ", val)
	}

	// Print remaining empty slots
	for i := count; i < cap(ch); i++ {
		fmt.Printf("[ ] ")
	}
	fmt.Println()

	// Restore values safely
	for _, v := range tmp {
		ch <- v
	}
}

func main() {
	ch := make(chan int, 3)

	go func() {
		fmt.Println("G1: Sending 100...")
		ch <- 100
		fmt.Println("G1: Stored 100 → slot filled")
	}()

	go func() {
		fmt.Println("G2: Sending 200...")
		ch <- 200
		fmt.Println("G2: Stored 200 → slot filled")
	}()

	go func() {
		fmt.Println("G3: Sending 300...")
		ch <- 300
		fmt.Println("G3: Stored 300 → slot filled")
	}()

	go func() {
		fmt.Println("G4: Sending 400 (will block if full)...")
		ch <- 400
		fmt.Println("G4: Stored 400 → became unblocked!")
	}()

	time.Sleep(1 * time.Second)
	printBuffer(ch, "Before receiving")

	fmt.Println("\nMain: Receiving one...")
	val := <-ch
	fmt.Println("Main: Received:", val)

	printBuffer(ch, "After receiving")

	time.Sleep(1 * time.Second)
	fmt.Println("End")
}
