package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(3 * time.Second)
        ch1 <- "Message from ch1"
    }()

    go func() {
        time.Sleep(4 * time.Second)
        ch2 <- "Message from ch2"
    }()

    select {
    case msg := <-ch1:
        fmt.Println("Received:", msg)
    case msg := <-ch2:
        fmt.Println("Received:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout!")
    }
}



/*

2) ADD TIMEOUT USING SELECT
If ch doesn't send anything in 2 seconds → timeout triggers.

3) RANDOM CASE WHEN MULTIPLE ARE READY
If 2 channels are ready at the same time, Go picks randomly among ready cases.

4) DEFAULT CASE — NON-BLOCKING SELECT




*/