package main

import (
	"fmt"
	"sync"
	"time"
)

func work(task string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(task, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go work("Emails", &wg)
	go work("Notifications", &wg)
	go work("Database Sync", &wg)

	wg.Wait()
}


// This function will receive a pointer to a WaitGroup.
/***


Because WaitGroup ka internal counter shared hota hai across goroutines.

If you passed it by value:

each goroutine would get its own copy

wg.Done() would update the copied version

original wg in main() kabhi reduce nahi hota

wg.Wait() → deadlock (infinite wait)

/// This is why we pass a pointer to the WaitGroup to ensure all goroutines
work on the same shared counter.
***/