package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	go printNumbers()

	fmt.Println("Main function is running................1")
	// time.Sleep(2 * time.Second)
	
	go print("Raghav")
	fmt.Println("Main function is running.................2")
	go print("GoLang")
	
	fmt.Println("Main function is running.................3")
	fmt.Println("Main function is running...")
	time.Sleep(2 * time.Second)
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
