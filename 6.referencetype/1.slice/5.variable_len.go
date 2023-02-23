package main

import "fmt"

// Sum 可変長の引数をとる
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6))
	fmt.Println(Sum())

	// スライスの可変長化
	sl := []int{10, 20, 30}
	fmt.Println(Sum(sl...))
}
