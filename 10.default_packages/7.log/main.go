package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	log.Print("Log\n")
	log.Println("Log2")
	log.Printf("Log%d\n", 3)

	//	log.Fatal("Exit", err)
	//	log.Fatalln("Exit", err)
	//	log.Fatalf("%s;%s", "Exit", err)
	//
	//	log.Panic("Exit", err)
	//	log.Panicln("Exit", err)
	//	log.Panicf("%s;%s", "Exit", err)

	f, err := os.Create("test.log")
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.Println("ファイルに書き込む")

	// 標準フォーマット
	log.SetFlags(log.LstdFlags)
	log.Println("A")

	// マイクロ秒を追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")

	// 時刻とファイルの行番号
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	log.SetFlags(log.Ltime | log.Llongfile)
	log.Println("D")

	log.SetFlags(log.LstdFlags)

	// ログprefixを設定
	log.SetPrefix("[LOG]")
	log.Println("E")

	// loggerの生成
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	_, err = os.Open("not found")
	if err != nil {
		logger.Fatalln("message")
	}
}
