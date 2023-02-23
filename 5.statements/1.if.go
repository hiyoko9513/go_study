package main

import "fmt"

func main() {
	a := 1

	// ベース
	if a == 2 {
		fmt.Println("2だ")
	} else if a == 1 {
		fmt.Println("1だ")
	} else {
		fmt.Println("それ以外")
	}

	// if文の前に簡易文を書ける
	if b := 100; b == 100 {
		fmt.Println("100だ")
	}

	// 簡易文付きの注意点
	// 簡易文で宣言した変数は、if文内でしか利用できない。
	x := 5
	if x := 2; true {
		fmt.Println(x) // 2
	}
	fmt.Println(x) // 5
}
