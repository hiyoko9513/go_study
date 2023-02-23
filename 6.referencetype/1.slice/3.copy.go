package main

import "fmt"

func main() {
	var a []int = []int{100, 200}

	// 💢注意： スライスの場合、下記のような代入をすると同じアドレスを参照してしまう。
	// 💢注意： この特性は参照型のみ
	// b := a
	// b[0] = 1000
	// fmt.Println(a)

	d := []int{1, 2, 3, 4, 5, 6}

	// copy(<参照先>, <参照元>) 参照元とは別のアドレスに値をコピーすることができる。
	// 返り値はコピーに成功した数
	// 前から順番に置き換える
	copy1 := copy(d, a)
	fmt.Println(copy1)
	fmt.Println(d)
}
