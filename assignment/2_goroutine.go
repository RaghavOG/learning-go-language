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

