package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const (
	specConfig = "config.json"
)

func main() {
	app := cli.NewApp()
	app.Name = "kamoc"
	app.Commands = []cli.Command{
		CreateCommand,
	}
	app.Run(os.Args)
}

var (
	StateCommand = cli.Command{
		Name:  "state",
		Usage: "Usage state!",
		Action: func(c *cli.Context) error {
			fmt.Println("state", c.Args().First())
			return nil
		},
	}

	CreateCommand = cli.Command{
		Name:  "create",
		Usage: "create a container",
		Action: func(c *cli.Context) error {
			fmt.Println("create", c.Args().First())
			return nil
		},
	}

	StartCommand = cli.Command{
		Name:  "start",
		Usage: "Usage start!",
		Action: func(c *cli.Context) error {
			fmt.Println("start", c.Args().First())
			return nil
		},
	}

	KillCommand = cli.Command{
		Name:  "kill",
		Usage: "Usage kill!",
		Action: func(c *cli.Context) error {
			fmt.Println("kill", c.Args().First())
			return nil
		},
	}

	DeleteCommand = cli.Command{
		Name:  "delete",
		Usage: "Usage delete!",
		Action: func(c *cli.Context) error {
			fmt.Println("delete", c.Args().First())
			return nil
		},
	}
)
