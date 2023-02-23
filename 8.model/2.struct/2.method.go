package main

import "fmt"

// User エラーは無視していい
type User struct {
	Name string
	Age  int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

// SetNameForPointer ポインター型と上記のメソッドの値型の２種類を使うことが推奨されていない
// 💢基本的には構造体に定義するメソッドはポインター型にするべき
func (u *User) SetNameForPointer(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	user1.SetName("A")
	user1.SayName()

	user1.SetNameForPointer("A")
	user1.SayName()

	user2 := &User{Name: "user2"}
	user2.SetNameForPointer("B")
	user2.SayName()

}
