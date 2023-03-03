package main

import "fmt"

// ジェネリクス
// 型を抽象化してコードの再利用に繋がる。
// 型安全性の維持

func PrintSliceInt(si []int) {
	for _, i := range si {
		fmt.Println(i)
	}
}

func PrintSliceString(ss []string) {
	for _, s := range ss {
		fmt.Println(s)
	}
}

// 従来の記述
func old(si []int, ss []string) {
	PrintSliceInt(si)
	PrintSliceString(ss)
}

// PrintSlice ジェネリクス
func PrintSlice[T any](slice []T) {
	for _, v := range slice {
		fmt.Println(v)
	}
}

func basic(si []int, ss []string) {
	PrintSlice[int](si)
	PrintSlice[string](ss)
}

func f[T fmt.Stringer](xs []T) []string {
	var result []string
	for _, v := range xs {
		result = append(result, v.String())
	}
	return result
}

type MyInt int

func (i MyInt) String() string {
	return fmt.Sprintf("%d", int(i))
}

func basic2() {
	fmt.Println(f([]MyInt{1, 2}))
}

func main() {
	//si := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//ss := []string{"a", "b", "c", "d", "e"}
	//old(si, ss)
	//basic(si, ss)
	basic2()
}
