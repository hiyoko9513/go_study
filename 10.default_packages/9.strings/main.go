package main

import (
	"fmt"
	"strings"
)

// 文字列操作パッケージ

func main() {
	// セパレーターを利用して結合
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1, s2)

	// 文字列検索
	//　検索出来たindexで返される
	i1 := strings.Index("ABCDE", "A")
	i2 := strings.Index("ABCDE", "D")
	fmt.Println(i1, i2)
	// 検索出来た最後のindexを返す
	i3 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i3)
	// 検索文字列のどれかが含まれていたらそのindexを返す
	i4 := strings.IndexAny("ABC", "ABC")
	// 検索文字列のどれかが含まれていたらその最後のindexを返す
	i5 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i4, i5)
	// 検索文字列の文字列で始まるか
	b1 := strings.HasPrefix("ABC", "A")
	// 検索文字列の文字列で終わるか
	b2 := strings.HasSuffix("ABC", "C")
	fmt.Println(b1, b2)
	// 検索文字列が含まれているか
	b3 := strings.Contains("ABC", "B")
	// 検索文字列のどれかが含まれているか
	b4 := strings.ContainsAny("ABCDE", "BD")
	fmt.Println(b3, b4)
	// 検索文字列の出現回数
	i6 := strings.Count("ABCABC", "B")
	// 検索対象の文字数の数＋１が返ってくる
	i7 := strings.Count("ABCABC", "")
	fmt.Println(i6, i7)

	// 文字数を繰り返して結合する
	s3 := strings.Repeat("ABC", 4)
	s4 := strings.Repeat("ABC", 0)
	fmt.Println(s3, s4)

	// 文字列置き換え
	s5 := strings.Replace("AAAAAA", "A", "B", 1)
	s6 := strings.Replace("AAAAAA", "A", "B", -1)
	fmt.Println(s5, s6)

	// 文字列分割
	s7 := strings.Split("A,B,C,D,E", ",")
	fmt.Println(s7)
	// 分割文字列を残して分割
	s8 := strings.SplitAfter("A,B,C,D,E", ",")
	fmt.Println(s8)
	// 分割数を指定出来る 下記では２分割
	s9 := strings.SplitN("A,B,C,D,E", ",", 2)
	fmt.Println(s9)
	// 分割文字列を残して分割数を指定出来る
	s10 := strings.SplitAfterN("A,B,C,D,E", ",", 4)
	fmt.Println(s10)

	// 小文字変換
	s11 := strings.ToLower("ABC")
	s12 := strings.ToLower("E")
	// 大文字変換
	s13 := strings.ToUpper("abc")
	s14 := strings.ToUpper("e")
	fmt.Println(s11, s12, s13, s14)

	// 前後の空白除去
	s15 := strings.TrimSpace("    -    ABC    -    ")
	s16 := strings.TrimSpace("\tABC\r\n")
	s17 := strings.TrimSpace("　　　　ABC　　　　")
	fmt.Println(s15, s16, s17)

	// スペーズで区切られた文字列を分割してスライスで返す
	s18 := strings.Fields("a b c")
	fmt.Println(s18)
	fmt.Println(s18[1])

}
