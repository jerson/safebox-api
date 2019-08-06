package main

import (
	"github.com/urfave/cli"
	"os"
	"safebox.jerson.dev/api/modules/config"
)

func init() {
	err := config.ReadDefault()
	if err != nil {
		panic(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "safebox-commands"
	app.Usage = "comands for safebox API"
	app.Version = config.Vars.Version

	app.Commands = []cli.Command{
		{
			Name:     "send:email:location",
			Category: "email",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "username, u",
				},
			},
			Action: func(c *cli.Context) error {

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
