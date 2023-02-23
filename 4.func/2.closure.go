package main

func main() {
	// 無名関数の定義
	f := func(x, y int) int {
		return x + y
	}
	f(1, 2)

	// 無名関数をそのまま実行
	func(x, y int) int {
		return x + y
	}(1, 2)

}
