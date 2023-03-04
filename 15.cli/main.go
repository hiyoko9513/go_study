package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

// コマンド実行
// go run 15.cli/main.go --animal cat

func run(args []string) error {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "animal",
			Value: "hiyoko",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.String("animal") == "hiyoko" {
			fmt.Println("ひよこ")
			return nil
		}
		if c.String("animal") == "cat" {
			fmt.Println("ねこ")
			return nil
		}
		return errors.New("not found")
	}

	return app.Run(args)
}

func main() {
	err := run(os.Args)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
