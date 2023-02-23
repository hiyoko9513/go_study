package main

import "fmt"

// 基本利用❌❌❌

func main() {
	fmt.Println("A")
	goto L
	fmt.Println("B")
L:
	fmt.Println("C")

	for {
		for {
			for {
				fmt.Println("Goto")
				goto D
			}
		}
	}
D:
	fmt.Println("end")

}
