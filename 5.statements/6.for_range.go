package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	// range (php的foreach)
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// range (keyの破棄)
	for _, v := range arr {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}
}
