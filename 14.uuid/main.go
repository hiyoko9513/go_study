package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	uuid1, _ := uuid.NewUUID()
	fmt.Println(uuid1.String())

	uuid2, _ := uuid.NewRandom()
	fmt.Println(uuid2.String())
}
