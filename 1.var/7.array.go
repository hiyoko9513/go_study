package main

import "fmt"

func main() {
	// 注意
	// 配列型は要素数を変更できない
	// スライス型は要素数を変更できる

	// 明示的な宣言
	var arr1 [2]int
	fmt.Println(arr1)

	// 要素数3の2文字のみ代入
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	// 暗示的な宣言
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 要素の数を自動でカウントしてくれる
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)

	arr5 := [5]int{1, 2, 3}
	fmt.Println(arr5)

	// 配列の取り出し
	fmt.Println(arr4[0])
	fmt.Println(arr4[2-1])

	// 配列の再代入
	arr4[1] = "F"
	fmt.Println(arr4)

	// 要素数の取得
	fmt.Println(len(arr4))

	// fmt.Println(arr1 == arr5)
	// fmt.Println(arr2 == arr3)

}
