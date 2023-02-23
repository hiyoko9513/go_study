package main

import (
	"fmt"
	"time"
)

func receiver2(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	// チャネルのクローズ
	close(ch1)
	fmt.Println(<-ch1)
	// 第２引数にチェネルのオープン状態がわかるようになっている
	// 💢注意；クローズしていても値が空になるまでは使えるチャネルとされるため、trueが帰ってきて、値も正常に取り出せる
	i, ok := <-ch1
	fmt.Println(i, ok)
	i, ok = <-ch1
	fmt.Println(i, ok)

	ch2 := make(chan int, 20)
	go receiver2("1.goroutine", ch2)
	go receiver2("2.goroutine", ch2)
	go receiver2("3.goroutine", ch2)
	j := 0
	for j < 100 {
		ch2 <- j
		j++
	}
	close(ch2)
	time.Sleep(3 * time.Second) // プログラム終了しないようにしている

}
