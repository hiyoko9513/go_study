package main

import (
	"fmt"
)

func basic() {
	ch1 := make(chan int, 20)
	ch2 := make(chan string, 20)
	ch2 <- "APPLE"

	// selectの意義
	// 下記のコードを実行しようとするとdeadlockを起こす
	// ch1にはデータが入っておらず、取り出すことができないため
	// deadlockするとプログラム全体が終了してしまう。
	// それを解消するためにselectを利用する
	//e1 := <-ch1
	//e2 := <-ch2
	//fmt.Println(e2)
	//fmt.Println(e1)

	// 💢selectの中の処理でチャネルの操作でなくてはならない。
	// 送受信どちらでも可
	// 💢実行される順番はランダムになっている。下記のコメントアウトを外して複数回実行すればわかる。
	//ch1 <- 100
	select {
	case v := <-ch1:
		fmt.Println(v + 1000)
	case v := <-ch2:
		fmt.Println(v + ".inc")
	default:
		fmt.Println("実行されない")
	}
}

// 実装例
func sample() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	ch3 := make(chan int, 5)

	go func() {
		for {
			i1 := <-ch1
			ch2 <- i1 * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch2
			ch3 <- i2 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch1 <- n:
			n++
		case i3 := <-ch3:
			fmt.Println("recieve", i3)
		default:
			if n > 100 {
				//close(ch1)
				break L
			}
		}
	}
}

func main() {
	basic()
	sample()
}
