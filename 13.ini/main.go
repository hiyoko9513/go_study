package main

import (
	"fmt"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port int
}

var Config ConfigList

// 自動で走る
func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port: cfg.Section("server").Key("port").MustInt(8080),
	}
}

func main() {
	fmt.Println(Config.Port)
}
