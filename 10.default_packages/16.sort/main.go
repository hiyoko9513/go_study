package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

func normal() {
	i := []int{5, 3, 2, 4, 5, 6, 4, 8, 9, 8, 7, 10}
	s := []string{"a", "z", "j"}
	fmt.Println(i, s)

	sort.Ints(i)
	sort.Strings(s)
	fmt.Println(i, s)
}

func normal2() {
	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}

	fmt.Println(el)

	// 標準
	// 安定ソートを保証しない
	// 下記のコードの場合、名前順ソートが破棄される
	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value })
	fmt.Println("---------")
	fmt.Println(el)
	fmt.Println("---------")

	// 安定ソート
	// 下記コードでは、名前順を維持しながら、値順を保っている
	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })
	fmt.Println("---------")
	fmt.Println(el)
	fmt.Println("---------")
}

func main() {
	//normal()
	normal2()
}
