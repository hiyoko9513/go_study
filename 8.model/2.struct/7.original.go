package main

import "fmt"

// MyInt 独自型
type MyInt int

// Print メソッドも作成出来る
func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	// 独自型なので、通常のintでは計算出来ないようになる
	//i := 100
	//fmt.Println(mi + i)

	mi.Print()
}
