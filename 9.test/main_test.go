package main

import "testing"

var Debug bool = false

// testから始まる関数をテスト用の関数とされる
// $ go test
// $ go test -v
// $ go test -v ./...  // フォルダー直下のすべてをテスト
// $ go test -cover ./...  // テストのcover率がわかる（関数すべてをテストを書いていると１００％となる）
func TestIsOne(t *testing.T) {
	if Debug {
		t.Skip("スキップする")
	}
	i := 1
	v := IsOne(i)

	if !v {
		t.Errorf("%d != %d", i, 1)
	}
}
