package main

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func main() {
	// 疑似乱数・単調
	ulid1 := ulid.Make()
	println(ulid1.String())

	// より高度
	// 💢goroutineでの並行生成は安全性にかけるためmutexなど利用する必要がありそう
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid2, _ := ulid.New(ms, entropy)
	println(ulid2.String())
}
