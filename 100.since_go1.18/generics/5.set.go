package main

import "fmt"

// comparable
// 型の比較が可能なもののみ入る制約
// 型比較出来ない型：例；MAP、Slice、関数
type Set[T comparable] map[T]struct {
}

func NewSet[T comparable](xs ...T) Set[T] {
	s := make(Set[T])
	for _, x := range xs {
		s.Add(x)
	}
	return s
}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s Set[T]) Remove(x T) {
	delete(s, x)
}

func main() {
	s := NewSet(1, 2, 3, 4, 5, 6, 7)
	s.Remove(1)
	fmt.Println(s)
}
