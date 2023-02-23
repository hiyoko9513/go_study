package main

import "fmt"

// User エラーは無視していい
type User struct {
	Name string
	Age  int
}

// Users ポインター型に出来る。値型でも出来る。
type Users []*User

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{}

	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)

	for _, u := range users {
		fmt.Println(*u)
	}

	users2 := make([]*User, 0) // make関数で作成も出来る
	users2 = append(users2, &user1, &user2)

	for _, v := range users2 {
		fmt.Println(*v)
	}
}
