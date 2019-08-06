package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"safebox.jerson.dev/api/modules/config"
)

func init() {
	err := config.ReadDefault()
	if err != nil {
		panic(err)
	}
}

func main() {

	s := gocron.NewScheduler()
	s.Every(1).Day().At("11:00").Do(sendMails)
	<-s.Start()
}

func sendMails() {
	fmt.Println("I am runnning task.")
}
