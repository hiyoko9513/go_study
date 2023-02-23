package main

import "fmt"

// User エラーは無視していい
type User struct {
	Name string
	Age  int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUserForPointer(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7)

	// 実際はポインター型を関数に渡す事が多い
	user8 := &User{}
	fmt.Println(user8)

	UpdateUser(user1)
	UpdateUserForPointer(user8)

	fmt.Println(user1)
	fmt.Println(user8)
}
