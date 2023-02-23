package main

import "fmt"

// 複数のgoroutine間でのデータの受け渡し
// 送受信、それぞれの専用のチャネルを宣言できる
// make(<チェネル型>[, <キャパ>])
// キャパの数をこえてデータを保持できない。エラーが起こるので注意

// 💢データを入れた順にデータを受け取る
// 💢チャネルからデータを取り出すとチャネルのデータもなくなる

func main() {
	// 受信専用のチャネル
	// var ch2 <-chan int
	// 送信専用のチャネル
	// var ch3 chan<- int

	// 送受信可能なチャネルの宣言
	var ch1 chan int // nil
	fmt.Println(cap(ch1))
	ch1 = make(chan int, 20)
	fmt.Println(cap(ch1))

	// 暗示的チャネルの作成 これなら一行ですむ
	// c1 := make(chan int)
	c2 := make(chan int, 10)

	// データの送信
	c2 <- 9
	fmt.Println(len(c2))
	// データの受信
	i := <-c2
	fmt.Println(i)
	fmt.Println(len(c2))

	// 💢データを入れた順にデータを受け取る
	// 💢チャネルからデータを取り出すとチャネルのデータもなくなる
	c2 <- 2
	c2 <- 3
	c2 <- 4
	fmt.Println("len", len(c2))

	fmt.Println(<-c2)
	fmt.Println("len", len(c2))

	fmt.Println(<-c2)
	fmt.Println("len", len(c2))

	fmt.Println(<-c2)
	fmt.Println("len", len(c2))

}
