package main

import (
	//f "fmt"
	. "fmt" // 推奨されないやり方ではある

	// go mod ini <モジュール名>　を実行する必要がある
	// <モジュール名>/foo
	"example.com/7.package/foo"
)

func appName() string {
	// 関数で定義された変数、定数は関数でないのみ利用可能
	const AppName = "Go App"
	var version = "1.0"
	return AppName + "" + version
}

func main() {
	// minはプライベートなので参照出来ない
	//f.Println(foo.min)
	Println(foo.ReturnMin())

	Println(appName())
}
