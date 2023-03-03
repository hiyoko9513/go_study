package main

import "fmt"

type Vector[T any] []T

type IntVector Vector[int]

func main() {
	var vi Vector[int] = []int{1, 2, 3, 4, 5}
	var vs Vector[string] = []string{"a", "b", "c", "d", "e"}
	var iv IntVector = []int{1, 2, 3, 4, 5}

	fmt.Println(vi)
	fmt.Println(vs)
	fmt.Println(iv)
}
