package main

import "fmt"

// User エラーは無視していい
type User struct {
	Name string
	Age  int
}

// NewUser コンストラクタの命名はNewから始まる。このようなコンストラクタ関数はよく利用される
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("user1", 10)

	fmt.Println(user1)
	fmt.Println(*user1)
}
