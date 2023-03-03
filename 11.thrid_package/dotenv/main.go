package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	envPath := "./11.thrid_package/dotenv/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Printf("読み込みエラー: %v", err)
	}
	hiyoko := os.Getenv("HIYOKO")
	fmt.Println(hiyoko)
}
