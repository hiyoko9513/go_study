package main

import "fmt"

// パッケージを呼び出した時、mainより先に呼び出される関数

func init() {
	fmt.Println("init")
}

// ❌２つに分けるのは非推奨
// 上から順に実行される
func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("main")
}
