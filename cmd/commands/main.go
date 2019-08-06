package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli"
	"os"
	"safebox.jerson.dev/api/modules/config"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/queue"
	"safebox.jerson.dev/api/repositories"
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
			Name:     "email:location",
			Category: "email",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "username, u",
				},
			},
			Action: func(c *cli.Context) error {

				ctx := context.NewContextSingle("command")
				defer ctx.Close()
				log := ctx.GetLogger("action")
				repo := repositories.NewUserRepository(ctx)
				user, err := repo.FindOneByUsername(c.String("username"))
				if err != nil {
					return err
				}

				token, err := queue.SendEmailLocationTask(fmt.Sprintln(user.ID))
				if err != nil {
					return err
				}

				log.Infof("Result: %s", token)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
