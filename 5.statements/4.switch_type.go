package main

import "fmt"

func anything(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 40000)
	}
}

func main() {
	// if文を利用してかけるが、Switchの方がより簡単に書けるため、Switchで書くことが推奨されている
	anything(1)
	anything("AAA")

	var x interface{} = 3

	// カターサーション
	// 型の復元
	// int型をfloatで復元しようとするとエラーが起きる
	// 第２返り値を設定することで、エラーは防ぐことができる
	// ❌エラー
	// f := x.(float64)
	// fmt.Println(f)
	// 🟡エラー回避
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("何もない")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s)
	} else {
		fmt.Println("対応していない型")
	}

	// 値を使わない場合
	switch x.(type) {
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't Know")
	}

	// 値を使いたい場合
	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println("I don't Know")
	}
}
