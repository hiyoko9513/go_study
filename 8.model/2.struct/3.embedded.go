package main

import "fmt"

// User エラーは無視していい
type User struct {
	Name string
	Age  int
}

type T struct {
	User User
}

// SetNameForPointer ポインター型と上記のメソッドの値型の２種類を使うことが推奨されていない
// 💢基本的には構造体に定義するメソッドはポインター型にするべき
func (u *User) SetNameForPointer(name string) {
	u.Name = name
}

// 型省略が出来る
//type T struct {
//	User
//}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)

	fmt.Println(t.User)

	fmt.Println(t.User.Name)

	// 型を省略した場合、下記のコードでも参照できる
	//fmt.Println(t.Name)

	t.User.SetNameForPointer("A")
	fmt.Println(t)
}
