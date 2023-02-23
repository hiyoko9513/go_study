package main

import "fmt"

func main() {
	ch3 := make(chan int, 4)
	ch3 <- 100
	ch3 <- 200
	ch3 <- 300
	close(ch3) // チャネルでforを利用する時はcloseしないとdeadlockする
	for i := range ch3 {
		fmt.Println(i)
	}
}
