package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "cli for hello app"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "name, n",
			Usage: "name of the input",
			Value: "World",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("Hello %s !!!", name)
		return nil
	}

	app.Run(os.Args)
}
