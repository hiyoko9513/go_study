package main

import (
	"fmt"
	"sync"
)

func main() {
	//wgを生成
	wg := new(sync.WaitGroup)
	// 待ち受けるgoroutineの数を設定
	// ❗Done関数の実行待ち数
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		wg.Done()
	}()

	wg.Wait()
}
