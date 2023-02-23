package main

import "fmt"

func main() {
	// 	ラベル付きfor文
	// 	break <ラベル名>
	// 	ラベルを抜けることができる
LOOP:
	for {
		for {
			for {
				fmt.Println("START")
				break LOOP
			}
			fmt.Println("処理しない")
		}
		fmt.Println("処理しない")
	}
	fmt.Println("END")

LOOP2:
	for i := 0; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j > 1 {
				continue LOOP2
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理しない")
	}
}
