package main

import (
	"fmt"
	"sync"
	"time"
)

var st2 struct{ A, B, C int }

var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	// ❗処理詳細
	// mutex.lockされると
	// １つのgoroutineが処理を完了しないと次のgoroutineが動作しないような仕組みになっている。
	// unlockされると次のgoroutineが動作を開始する
	// 💢非同期処理が実質停止状態になるため、範囲をしっかり考えて利用する必要がある。
	mutex.Lock()

	st2.A = n
	time.Sleep(time.Microsecond)
	st2.B = n
	time.Sleep(time.Microsecond)
	st2.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st2)

	mutex.Unlock()
}

func main() {
	mutex = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}
	for {
	} // プログラムが停止しないように
}
