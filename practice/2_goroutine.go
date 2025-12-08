package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3) // We will run 3 goroutines

	go func() {
		printNumbers()
		wg.Done()
	}()

	fmt.Println("Main function is running................1")

	go func() {
		print("Raghav")
		wg.Done()
	}()

	fmt.Println("Main function is running.................2")

	go func() {
		print("GoLang")
		wg.Done()
	}()

	fmt.Println("Main function is running.................3")

	wg.Wait() // Wait for all goroutines to end
	fmt.Println("Main function is running.................4")
}

func print(msg string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(200 * time.Millisecond)
	}
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}
}
