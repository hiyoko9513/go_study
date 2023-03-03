package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ABC struct{}

type User struct {
	// json:"<フィールド名>,<型>"
	// json:"<フィールド名>,omitempty" // データが存在しない場合、非表示に出来る
	// json:"-" 非表示に出来る
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"-"`
	ABC     *ABC      `json:"A",omitempty` //pointer型にすることで完全非表示にできる
}

// MarshalJSON ❗json.Marshalが呼ばれた時、自動で呼ばれる
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

// UnmarshalJSON ❗json.Unmarshalが呼ばれた時、自動で呼ばれる
func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	// json変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))
	fmt.Printf("%T\n", bs)

	// json → struct
	var u2 User
	if err := json.Unmarshal(bs, &u2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2)
}
