package main

import (
	"fmt"
	"sync"
)

func square(num int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- num * num
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 5) // buffered to avoid blocking sends

	nums := []int{2, 4, 6, 8, 10}
	wg.Add(len(nums))
	for _, n := range nums {
		go square(n, ch, &wg)
	}

	// Wait for all goroutines to finish sending, then close channel.
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Receive exactly 5 results (or range over ch until closed).
	for result := range ch {
		fmt.Println(result)
	}
}

