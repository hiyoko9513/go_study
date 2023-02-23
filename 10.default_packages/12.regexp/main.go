package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 標準
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match)

	// Compile
	// ❗大量の処理を行い場合はこちらの方がいい。予めコンパイルすることが出来る
	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match)
	// 基本こちらの方がいい
	// 動的に正規表現のパターンを変化させる場合、上記の方を利用
	re2 := regexp.MustCompile("A")
	match = re2.MatchString("ABC")
	fmt.Println(match)

	// フラグ一覧
	//
	// i 大文字小文字を区別しない
	// m マルチモード（^と&が文頭文末に加えて、行頭行末にマッチ）
	// s .が\にマッチ
	// U 最小マッチへ変換(a*はa*?へ、a+はa+?へ)
	re3 := regexp.MustCompile(`(?i)abc`)
	match = re3.MatchString("ABC")
	fmt.Println(match)

	// 幅を持たない正規表現
	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")
	fmt.Println(match)
	match = re4.MatchString("  ABC  ")
	fmt.Println(match)

	// すべてマッチ
	re5 := regexp.MustCompile(".")
	match = re5.MatchString("ABC")
	fmt.Println(match)
	match = re5.MatchString("\n")
	fmt.Println(match)

	// 繰り返しのパターン
	re6 := regexp.MustCompile("a+b*")
	fmt.Println(re6.MatchString("ab"))
	fmt.Println(re6.MatchString("a"))
	fmt.Println(re6.MatchString("aaaabbbbbbbb"))
	fmt.Println(re6.MatchString("b"))

	re7 := regexp.MustCompile("A+?A+?X")
	fmt.Println(re7.MatchString("AAX"))
	fmt.Println(re7.MatchString("AAAAAAXXXX"))

	// 文字クラス
	re8 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re8.MatchString("Y"))

	re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	fmt.Println(re9.MatchString("ABC"))
	fmt.Println(re9.MatchString("abcdefg"))

	re10 := regexp.MustCompile(`[^0-9A-Za-z_]`)
	fmt.Println(re10.MatchString("ABC"))
	fmt.Println(re10.MatchString("あ"))

	// グループ
	re1 = regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re1.MatchString("abcxyz"))
	fmt.Println(re1.MatchString("ABCXYZ"))
	fmt.Println(re1.MatchString("ABCxyz"))
	fmt.Println(re1.MatchString("ABCabc"))

	////////////////////////////////////////
	// 文字列の取得
	////////////////////////////////////////
	// １箇所取得
	fs1 := re1.FindString("AAAAABCXYZZZZ")
	fmt.Println(fs1)
	// 複数箇所取得
	fs2 := re1.FindAllString("ABCXYZABCXYZ", -1)
	fmt.Println(fs2)

	// 分割
	fmt.Println(re1.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1))
	re1 = regexp.MustCompile(`\s+`)
	fmt.Println(re1.Split("aaaaaaaaaa     bbbbbbb  cccccc", -1))

	// 置き換え
	fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ","))
	re1 = regexp.MustCompile(`、|。`)
	fmt.Println(re1.ReplaceAllString("私は、Golangを使用する、プログラマー。", ""))

	// sample
	re1 = regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
	0123-456-7889
	111-222-333
	556-787-899
	`
	ms := re1.FindAllStringSubmatch(s, -1)
	for _, v := range ms {
		fmt.Println(v)
	}
	fmt.Println(re1.ReplaceAllString("Tel: 000-111-222", "$3-$2-$1"))
	fmt.Println(re1.ReplaceAllString("Tel: 000-111-222", "あ-い-う"))

}
