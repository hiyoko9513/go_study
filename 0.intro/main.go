/*
ソースフォルダーの場所
GOPATH

簡易実行
go run main.go

ビルド(mac)
go build -o main main.go

ビルド(win)
go build -o main.exe main.go
*/

// 全ての要素はパッケージに属する
// パッケージの宣言は１つのみ
package main

// フォーマットパッケージ
// パッケージ一覧 https://pkg.go.dev/std
import (
	"fmt"
	"time"
)

/*
main関数 エントリーポイント
Goが実行するのはメインパッケージの中のメイン関数の中
*/
func main() {
	fmt.Println("hello world")
	fmt.Println(time.Now())

	// Printfについて
	// https://blog.y-yuki.net/entry/2017/05/02/000000
	fmt.Printf("My Name is %v", "Go")
}
