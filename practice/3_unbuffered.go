package main 

import (
	"fmt"
	"time"
)


func main(){
	ch := make(chan int)

go func() {
    fmt.Println("Sending...1")
    ch <- 100
    fmt.Println("Sent!")
}()
go func() {
    fmt.Println("Sending...2")
    ch <- 100
    fmt.Println("Sent!")
}()
go func() {
    fmt.Println("Sending...3")
    ch <- 100
    fmt.Println("Sent!")
}()

time.Sleep(1 * time.Second)
fmt.Println("Receiving...")
fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)
time.Sleep(1 * time.Second)
}








/**

Buffered = dabba jisme kuch storage space hai
Unbuffered = dabba jisme ZERO space hai

So unbuffered channel is like:
	Sender puts a message
	BUT there is no place to keep it
	So sender blocks until receiver takes it

Buffered channel is like:
	Sender puts message in a box
	Box has some capacity
	Sender continues until the box is full
	Receiver removes whenever needed


Unbuffered Channel (no storage → blocking)
Buffered Channel (has storage → non-blocking until full)

Use unbuffered when:
	You want strict synchronization
	Sender must wait for receiver
	Real-time hand-off required
	Order matters

Use buffered when:
	Work comes fast → must not block
	Workers are slower → need queue
	You want decoupling
	Parallelism required

***/