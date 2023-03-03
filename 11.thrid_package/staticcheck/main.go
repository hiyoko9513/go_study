package main

import "fmt"

func main() {
	bad_name := bad_name()
	fmt.Println("Hello World", bad_name)
}

func bad_name() string {
	return "bad_name"
}
