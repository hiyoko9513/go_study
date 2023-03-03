package main

import (
	"fmt"
	"net/url"
)

func main() {
	// url解析
	u, _ := url.Parse("http://example.com/search?a=1&b=2#top")
	fmt.Println("Scheme:" + u.Scheme)
	fmt.Println("Host:" + u.Host)
	fmt.Println("Path:" + u.Path)
	fmt.Println("RawQuery:" + u.RawQuery)
	fmt.Println("Fragment:" + u.Fragment)
	fmt.Println(u.IsAbs())
	fmt.Println(u.Query())

	fmt.Println("--------------------")

	// url生成
	url := &url.URL{}
	url.Scheme = "https:"
	url.Host = "google.com"
	q := url.Query()
	q.Set("q", "Golang")
	url.RawQuery = q.Encode()
	fmt.Println(url)

}
