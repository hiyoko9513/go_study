package main

import (
	"fmt"
	"log"
	"os"
)

func exit() {
	// ❗exitするとプログラムを終了する。deferも破棄される。
	defer func() {
		fmt.Println("hiyoko")
	}()
	os.Exit(1)
}

func logFatal() {
	// Openファイルを読み込む
	// log.Fatalln エラーログを出力して、エラー落ちさせる。
	_, err := os.Open("not found")
	if err != nil {
		log.Fatalln(err)
	}
}

// args 実行手順が少しややこしい(💢この関数のみ実行出来る状態にしておく。カレントディレクトリであること。)
// １，$ go build -o <ビルドファイル名（任意）> <os.Argsが使われているファイル main.go>
// ２，$ ./<ビルドファイル名> 1 2 3
func args() {
	// ❗コマンドラインに引数をとる
	// インデックス０はコマンド名が入るので注意
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	fmt.Printf("length=%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}

func openRead() {
	f, err := os.Open("test.log")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

	nn, err := f.ReadAt(bs, 10)
	fmt.Println(nn)
	fmt.Println(string(bs))

	//////////////////////////////

	//offset, err := f.Seek(10, os.SEEK_SET)
	//offset, err = f.Seek(-2, os.SEEK_CUR)
	//offset, err = f.Seek(0, os.SEEK_END)

	fi, err := f.Stat()
	fmt.Println(fi.Name())
	fi.Size()
	fi.Mode()
	fi.ModTime()
	fi.IsDir()
}

// createWrite
// 💢createは既存ファイルが存在している場合は、削除して作り直されるので注意
func createWrite() {
	f, _ := os.Create("foo.txt")
	f.Write([]byte("Hello\n"))
	f.WriteAt([]byte("Golang"), 7)
	f.Seek(0, os.SEEK_END) // offset位置の変更（カーソル位置の変更と認識でいいかも）
	f.WriteString("Yaah")
}

func fileOpen() {
	//os.O_RDONLY 読み込み専用
	//os.O_WRONLY 書き込み専用
	//os.O_RDWR 読み書き
	//os.O_APPEND ファイル末尾に追記
	//os.O_CREATE ファイルがなければ作成
	//os.O_TRUNC 可能であれば空にする
	// 複数渡せる
	f, err := os.OpenFile("foo.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}

func remove() {
	_ = os.Remove("foo.txt")
	_ = os.Remove("foo")
	_ = os.RemoveAll("foo")
}

func rename() {
	_ = os.Rename("foo.txt", "bar.txt")
	_ = os.Rename("bar.txt", "foo.txt")
	_ = os.Rename("foo", "bar")
	_ = os.Rename("foo.txt", "bar/foo.txt")
}

func main() {
	exit()
	//logFatal()
	//args()
	//openRead()
	//createWrite()
	//fileOpen()
	//remove()

	// todo まとめる
	//dir, err := os.Getwd()
	//fmt.Println(dir)
	//
	//err = os.Chdir("path/to/dir")
	//
	//err = os.Mkdir("test", 0775)
	//err = os.MkdirAll("foo/fooo/foooo", 0775)
	//
	//f, err = os.Open(".")
	//if err != nil {
	//	log.Println(err)
	//}
	//defer f.Close()
	//
	//fis, err := f.Readdir(0)
	//for _, fi := range fis {
	//	if fi.IsDir() {
	//		fmt.Println(fi)
	//	}
	//}
	//
	//host, err := os.Hostname()
	//fmt.Println(host)
	//
	//for _, v := range os.Environ() {
	//	fmt.Println(v)
	//}
	//
	//fmt.Println(os.Getpid())
	//os.Getppid()
	//os.Getuid()
	//os.Geteuid()
	//os.Getgid()
	//os.Getegid()
}
