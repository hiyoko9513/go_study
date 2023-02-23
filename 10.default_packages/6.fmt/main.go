package main

import (
	"fmt"
	"os"
)

func main() {
	// 改行
	fmt.Println("表示")

	// 標準
	print("Hello")
	// 改行
	println("Hello!", "www")

	// print系
	fmt.Print("Hello")
	fmt.Println("Hello!")
	fmt.Println("Hello", "world!") // 連結
	fmt.Println("Hello", "world!")
	fmt.Printf("%s\n", "Hello")
	fmt.Printf("%#v\n", "Hello")

	// sprint系
	fmt.Sprint("Hello")
	fmt.Sprintf("Hello")
	fmt.Sprintln("Hello")

	// fsprint系 書き込み先を指定できる
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

}
