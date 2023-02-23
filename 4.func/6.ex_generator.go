package main

import (
	"fmt"
)

// ジェネレーターの実装例
// 何らかのルールに従って連続した値を返し続ける仕組み
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	int1 := integers()

	fmt.Println(int1()) // => "1"
	fmt.Println(int1()) // => "2"
	fmt.Println(int1()) // => "3"

	int2 := integers()
	fmt.Println(int2()) // => "1"
}
