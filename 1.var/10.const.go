package main

import "fmt"

func main() {

	// 先頭を大文字にすることで、パブリックなものとして認識する
	// 小文字の場合はローカル
	// 変数名、関数名、クラス名など、それぞれと同様
	const Pi = 3.14

	const zero = 0.0

	const (
		URL      = "http://xxx.jp"
		SiteName = "test"
	)

	const (
		A = 1
		B
		C
		D = "A"
		E
		F
	)

	// エラーになる
	// var big int = 9223372036854775807 + 1
	// エラーにならない
	const big = 9223372036854775807 + 1

	fmt.Println(Pi, URL, SiteName)
	fmt.Println(big - 1)

	// iota
	// 連続する数字を返す
	const (
		c0 = iota // c0 == 0
		c1 = iota // c1 == 1
		c2 = iota // c2 == 2
	)

}
