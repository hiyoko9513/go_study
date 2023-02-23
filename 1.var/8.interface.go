package main

import "fmt"

func main() {
	// 「interface{}」波括弧までが型
	// 超汎用的な型
	// データ特有の演算は出来ない（例：数字計算ができない）
	// 初期値はnil
	var x interface{}
	fmt.Println(x)

	x = 1
	x = 3.14
	x = "A"
	x = [...]int{3, 4, 5, 6}
	fmt.Println(x)
	// x = 2
	// var y interface{} = 3
	// fmt.Println(x + y)
}
