package main

import (
	"fmt"
)

// 関数を返す関数
func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func main() {
	f := returnFunc()
	f() // => "I'm a function"
}
