package main

import "fmt"

// ポインタとは、値型に分類されるデータ構造(基本型、参照型、構造体）のメモリ上のアドレスと型の情報のことです。
// Goではこれを使って、データ構造を間接的に参照、操作ができる様になっている。
// 💢参照型は元からその機能を含んでいる為、基本的には不要となっています。
func double(i int) {
	i = i * 2
}

func doubleForPointer(i *int) {
	*i = *i * 2
}

func main() {
	var n int = 100
	fmt.Println(n)
	// アドレスの取得
	fmt.Println(&n)
	// 0xc000136008

	// 返り値を受け取っていないので値に変化なし
	double(n)
	fmt.Println(n)

	// ポインタ型の明示的宣言
	var p *int = &n
	fmt.Println(&p)
	// 0xc000234567
	fmt.Println(p)
	// 0xc000136008
	// 実態の表示
	fmt.Println(*p)
	// 100

	n = 200
	fmt.Println(*p)
	// 200
	*p = 300
	fmt.Println(n)
	// 300

	doubleForPointer(&n)
	fmt.Println(n)

	doubleForPointer(p)
	fmt.Println(*p)

	fmt.Println(&n)
	fmt.Println(*&n)
	fmt.Println(&*&n)
}
