package main

import (
	"fmt"
	"sync"
)

func square(n int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := n * n
	out <- result
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	out := make(chan int, len(nums)) // enough buffer
	var wg sync.WaitGroup
	wg.Add(len(nums))

	// --- FAN OUT ---
	for _, n := range nums {
		go square(n, out, &wg)
	}

	// Close channel AFTER all goroutines finish
	go func() {
		wg.Wait()
		close(out)
	}()

	// --- FAN IN ---
	for result := range out {
		fmt.Println(result)
	}
}
