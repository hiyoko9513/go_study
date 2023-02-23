package main

import (
	"fmt"
)

func main() {
	// ベース
	// データの型は揃えないとエラーになる
	n := 4
	// n := "4" ❌
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("どれでもない")
	}

	// 簡易文付き
	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("どれでもない")
	}

	// 式を使ってケース式を書くことができる
	// この場合、Switchに変数を送る必要がなくなる
	// switch {} ← 変数を送る必要❌
	n2 := 6
	switch {
	case n2 > 0 && n2 < 4:
		fmt.Println("0 < n2 < 4")
	case n2 > 3 && n2 < 6:
		fmt.Println("3 < n2 < 6")
	}

	// ケースに式と値の２パターンを混在させることは❌
	/*
		switch n := 2; n {
		case 1, 2, 3:
			fmt.Println("0 < n2 < 4")
		case n > 3 && n < 6:
			fmt.Println("3 < n2 < 6")
		}
	*/

}
