package main

import "fmt"

func main() {

	// int は環境依存になる：win32bitならint32になる
	// bit数が違う場合、計算式が成り立たない
	var i int = 100

	fmt.Println(i + 50)

	var ii int64 = 50

	// ↓エラー
	// fmt.Println(i + ii)

	// %T 書式指定子
	// 型変換
	fmt.Printf("%T\n", ii)
	fmt.Printf("%T\n", int32(ii))

}
