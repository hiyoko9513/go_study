package main

import "fmt"

/*
func 関数名(引数　引数の型) 戻り値型 {
	関数の中身
	return 返す値
}
*/

// Plus 複数引数が存在し、型が同一の場合下記のように省略できる
func Plus(x, y int) int {
	return x + y
}

// Div 複数の返り値がある場合は、（）
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// Double 返り値の変数事態を指定することが出来る
func Double(price int) (result int) {
	result = price * 2
	return
}

// Noreturn noreturn 返り値なし
func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	r1 := Plus(10, 20)
	fmt.Println(r1)

	r2, _ := Div(30, 20)
	fmt.Println(r2)

	r3 := Double(100)
	fmt.Println(r3)

	Noreturn()
}
