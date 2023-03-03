package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}
type List []Entry

// //////////////////////////////////////////////////////////
// ❗この３つの関数を用意すれば、独自のカスタムソートを作成出来る
func (l List) Len() int {
	return len(l)
}
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// メインカスタム部分
func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return (l[i].Name < l[j].Name)
	} else {
		return (l[i].Value < l[j].Value)
	}
}

////////////////////////////////////////////////////////////

func main() {

	m := map[string]int{"ada": 1, "hoge": 4, "basha": 3, "poeni": 3}
	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}

	//sort.SliceStable(lt, func(i, j int) bool { return lt[i].Name < lt[j].Name })
	//fmt.Println(lt)

	// ソート
	sort.Sort(lt)
	fmt.Println(lt)

	// リバーズ（ソートをひっくり返す）
	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt)

}
