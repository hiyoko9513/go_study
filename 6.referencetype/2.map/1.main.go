package main

import "fmt"

// 連想配列
// map[<keyの型>]<valueの型>{<key>: <value>,,,,,,}
// make 空のmapを作成できる
// delete(<map>, <key>) mapのkeyの削除

func main() {
	// 明示的宣言
	var m = map[string]int{"a": 100, "b": 200}
	fmt.Println(m)

	// 暗示的宣言
	m2 := map[string]int{"a": 100, "b": 200}
	fmt.Println(m2)

	// 複数map
	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	/*
		var m3 map[string]int
		m3["pc"] = 5000
		fmt.Println(m3)
	*/

	m4 := make(map[int]string)
	m4[1] = "US"
	m4[2] = "JAPAN"
	m4[3] = "CHINA"
	fmt.Println(m4)

	// 存在しないkeyは、空で返ってくる
	s := m["a"]
	fmt.Println(s)
	s1 := m["c"]
	fmt.Println(s1)

	// エラーハンドリング
	s2, ok := m["b"]
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(s2)

	s3, ok := m["c"]
	fmt.Println(s3)

	m["b"] = 300
	m["c"] = 500

	// mapの削除
	delete(m, "c")

	fmt.Println(len(m))

	// map for
	for k, v := range m {
		fmt.Println(k, v)
	}

	// valueにスライス
	m5 := map[int][]string{
		1: {"A"},
		2: {"B", "C"},
		3: {"D", "E"},
	}
	fmt.Println(m5)

	// valueにmap
	m6 := map[int]map[float64]string{
		1: {3.14: "円周率"},
	}
	fmt.Println(m6)

}
