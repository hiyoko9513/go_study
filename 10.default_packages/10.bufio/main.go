package main

import (
	"bufio"
	"fmt"
	"os"
)

// バッファをもたせる
// 動作させてばわかる
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr)
	}

}
