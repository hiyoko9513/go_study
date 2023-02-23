package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 注意
	// アンダースコア変数の意味
	// 宣言はするものの利用はしないという意味
	// goでは一度宣言した変数は利用しないと怒られるため

	var s string = "100"

	// string to int
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)

	h := "Hello World"
	// string()がない場合、バイト型で表示される
	fmt.Println(string(h[0]))

	var ii int = 321
	var ss string
	// int to string
	ss = strconv.Itoa(ii)
	fmt.Println(ss)

}
