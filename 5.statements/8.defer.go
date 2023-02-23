package main

import (
	"fmt"
)

// 関数終了時に実行される処理

func TestDefer() {
	// 複数デファーを登録した場合は、下から順に実行される
	defer fmt.Println("END")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("START")
}

func main() {
	TestDefer()

	// ファイル書き込み
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// file.Write([]byte("Hello\n"))

	// クロージャーなども書ける
	defer func() {
		fmt.Println("Func A")
		fmt.Println("Func B")
		fmt.Println("Func C")
	}()
}
