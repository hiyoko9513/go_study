package main

import (
	"fmt"
	"math"
)

func main() {
	// 円周率
	fmt.Println(math.Pi)

	// ２の平方根
	fmt.Println(math.Sqrt2)

	// それぞれの方の最大値最小値
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.SmallestNonzeroFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.SmallestNonzeroFloat64)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)

	// 絶対値
	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100))

	// 累乗
	fmt.Println(math.Pow(0, 2))
	fmt.Println(math.Pow(2, 2))

	// 平方根
	fmt.Println(math.Sqrt(2))
	// 立方根
	fmt.Println(math.Cbrt(8))

	// 最大値、最小値を返す
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))

	// 小数点以下切り捨て
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))

	// 引数の数値より小さい最大の整数を返す
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))

	// 引数の数値より大きい最小の整数を返す
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(-3.5))

	// todo まとめる
	//n := math.Sqrt(-1)
	//fmt.Println(math.IsNaN(n))
	//math.Inf(0)
	//math.Inf(-1)
	//math.NaN()
}
