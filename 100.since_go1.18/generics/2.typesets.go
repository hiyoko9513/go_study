package main

import "fmt"

// 下記のMyIntのような独自型の大本のint型を見てほしい場合、「~」を先頭に付ける必要がある。
type Number interface {
	~int | int32 | int64 | float64
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type MyInt int

func main() {
	fmt.Println(Max[int](1, 2))
	fmt.Println(Max[float64](1.2, 1.1))
	fmt.Println(Max[MyInt](1, 2))
}
