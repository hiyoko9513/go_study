package main

import "fmt"

// go1.18~~
// anyには何でも入る。
// 今後はinterface型ではなく、anyを使うのがいいかも
func main() {
	var a any
	a = 1
	a = true

	fmt.Println(a)
}
