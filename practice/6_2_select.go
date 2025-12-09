package main

import (
    "fmt"
    "time"
)

func worker(name string, ch chan string) {
    time.Sleep(time.Duration(len(name)) * 200 * time.Millisecond)
    ch <- name + " finished"
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go worker("Raghav", ch1)
    go worker("GoLang", ch2)

    select {
    case msg := <-ch1:
        fmt.Println("Winner:", msg)
    case msg := <-ch2:
        fmt.Println("Winner:", msg)
    }
}
