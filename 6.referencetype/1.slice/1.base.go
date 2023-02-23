package main

import "fmt"

func main() {
	var i []int
	fmt.Println(cap(i))

	// 明示的宣言
	var a []int = []int{100, 200}

	// 暗黙的宣言
	b := []string{"A", "B"}
	fmt.Println(b)

	// メイク関数によるスライスの作成
	// make(<型>, <要素数>[, <キャパ(メモリ使用量的な)>)
	l := make([]int, 4, 5)
	fmt.Println(l)

	// 値の代入
	a[0] = 1000
	fmt.Println(a)

	// 値の取り出し
	// 配列[<インデックス番号から>[, <インデックス番号の手前まで>]]
	fmt.Println("値の取り出し START")
	d := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(d)
	fmt.Println(d[0])
	fmt.Println(d[2:4])
	fmt.Println(d[:2])
	fmt.Println(d[2:])
	fmt.Println(d[:])
	fmt.Println(d[len(d)-1])
	fmt.Println(d[1 : len(d)-1])
	fmt.Println("値の取り出し END")

}
