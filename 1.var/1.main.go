package main

import "fmt"

/*
ヴァリアブル＝変数

明示的宣言
var 変数名 型 = 値

暗黙的
変数名 := 値

定義した変数は利用する必要がある
*/

var i int = 100
var s string = "Golang"

var t, f bool = true, false

var (
	ii int    = 1000
	ss string = "Go"
)

var i2 int

var sss string

// i3 := 100

func main() {
	fmt.Println(i2)
	i2 = 150
	fmt.Println(i2)

	fmt.Println(sss)
	sss = "Go!!"
	fmt.Println(sss)

	i2 = 200
	fmt.Println(i2)

	i3 := 200
	fmt.Println(i3)

	// i3 := 300

	i3 = 300

	// i3 = "string"

}
