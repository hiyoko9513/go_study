package main

import "fmt"

// 非推奨❌

func main() {
	// recoverはpanicで起こったエラーを回復させることができる。
	// 基本デファーと一緒に用いる
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	// 強制プログラム終了ができる
	panic("runtime error")
	fmt.Println("START")
}
