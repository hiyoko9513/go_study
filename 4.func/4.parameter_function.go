package main

import (
	"fmt"
)

// 関数を引数にとる関数
func callFunction(f func()) {
	f()
}

func main() {
	callFunction(func() {
		fmt.Println("I'm a function")
	}) // => "I'm a function"
}
