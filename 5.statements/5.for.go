package main

import "fmt"

func main() {

	// 無限ループ
	// for {
	// 	fmt.Println("るーぷ")
	// }

	// ベース
	i := 0
	for {
		i++
		fmt.Println(i)
		if i == 3 {
			break // 一番近いfor文を抜ける
		}
	}

	// 条件付きfor文
	// trueだとループする
	point := 1
	for point < 10 {
		fmt.Println(point)
		point++
	}

	// 古典的for文
	for i := 0; i < 100; i++ {
		if i == 3 {
			fmt.Println(3)
			continue
		}
		if i > 30 {
			fmt.Println("抜けた")
			break
		}
		fmt.Println(i)
	}
	// ２ずつ足していく場合
	// for i := 0; i < 100; i += 2 {
	// 	fmt.Println(i)
	// }

	// 古典的for文(arrバージョン)
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}
}
