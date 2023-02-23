package main

import "fmt"

func main() {
	// 条件
	fmt.Println(1 == 1)
	fmt.Println(1 == 2)
	fmt.Println(4 >= 8)
	fmt.Println(true != false)

	// 複数条件
	fmt.Println(true && false == true)
	fmt.Println(true || false == true)

	// 条件の否定
	fmt.Println(!true)
	fmt.Println(!false)
}
