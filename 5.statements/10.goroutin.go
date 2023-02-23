package main

import (
	"fmt"
)

// 並行処理

func sub() {
	i := 0
	for {
		fmt.Println(i)
		i++
	}
}

func main() {
	// go文を利用することで暗黙的に並行処理をしてくれる
	// クロージャーも利用できる
	go sub()

	for {
		fmt.Println("main")
	}
}
