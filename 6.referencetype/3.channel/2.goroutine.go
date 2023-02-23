package main

import (
	"fmt"
	"time"
)

func receiver(c chan int) {
	for {
		fmt.Println(<-c)
	}
}

func main() {
	var ch1 chan int
	var ch2 chan int
	ch1 = make(chan int)
	ch2 = make(chan int)

	go receiver(ch1)
	go receiver(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}
}
