package main

import (
	"fmt"
)

// クロージャーの実装例
func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := later()

	fmt.Println(f("Golang"))   // => ""
	fmt.Println(f("is"))       // => "Golang"
	fmt.Println(f("awesome!")) // => "is"
}
