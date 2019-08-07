package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jasonlvhit/gocron"
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
	ctx := context.NewContextSingle("cron")
	defer ctx.Close()
	log := ctx.GetLogger("main")

	ch := make(chan string)
	log.Info("running ")

	go func() {
		s := gocron.NewScheduler()
		s.Every(1).Day().At(config.Vars.Cron.TimeEmail).Do(sendMails)
		<-s.Start()
	}()

	go func() {
		s := gocron.NewScheduler()
		s.Every(1).Minutes().Do(deleteAccessToken)
		<-s.Start()
	}()
	<-ch

}

func deleteAccessToken() {
	ctx := context.NewContextSingle("deleteAccessToken")
	defer ctx.Close()
	log := ctx.GetLogger("deleteAccessToken")
	repo := repositories.NewAccessTokenRepository(ctx)
	err := repo.DeleteExpired()
	if err != nil {
		log.Error(err)
	} else {

		log.Info("removed all expired tokens")
	}

}

func sendMails() {
	ctx := context.NewContextSingle("sendMails")
	defer ctx.Close()
	log := ctx.GetLogger("sendMails")

	err := sendMailsPage(ctx, 0)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("send all emails")
	}
}

func sendMailsPage(ctx context.Context, offset int) error {
	log := ctx.GetLogger("sendMailsPage")
	repo := repositories.NewUserRepository(ctx)

	limit := 100
	list, err := repo.ListLocationEnabled(offset, limit, "id", "asc")
	if err != nil {
		return err
	}

	if len(list.Items) < 1 {
		log.Infof("finished")
		return nil
	}

	for _, user := range list.Items {

		token, err := queue.SendEmailLocationTask(fmt.Sprint(user.ID))
		if err != nil {
			return err
		}

		log.Infof("Send to user: %s ==> %s", user.Username, token)
	}

	return sendMailsPage(ctx, offset+limit)
}
