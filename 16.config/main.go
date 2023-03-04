package main

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

// そのた利用方法について
// https://pkg.go.dev/github.com/gookit/config/v2#section-readme
// jsonの型でも利用できる
func main() {
	config.WithOptions(config.ParseEnv)

	// add driver for support yaml content
	config.AddDriver(yaml.Driver)
	err := config.LoadFiles("example.yml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("config data: \n %#v\n", config.Data())
}
