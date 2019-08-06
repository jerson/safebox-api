package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"safebox.jerson.dev/api/modules/commands"
	"safebox.jerson.dev/api/modules/config"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/queue"
	"strconv"
)

func init() {
	err := config.ReadDefault()
	if err != nil {
		panic(err)
	}

}

func main() {
	ctx := context.NewContextSingle("queue")
	defer ctx.Close()
	err := run(ctx)
	if err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {

	log := ctx.GetLogger("run")
	conn, err := queue.StartServer()
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := queue.GetChannel(conn)
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := queue.GetQueue(ch, "default")
	if err != nil {
		return err
	}

	messages, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for msg := range messages {

			var err error
			var message queue.Message

			err = json.Unmarshal(msg.Body, &message)
			if err != nil {
				log.Error(err)
				_ = msg.Ack(true)
				continue
			}

			ctx := context.NewContextSingle("task")
			log.Infof("new task: %s => %s", message.Name, message.Token)

			switch message.Name {
			case "email:location":

				userID, err := strconv.Atoi(message.Params["userID"])
				if err != nil {
					log.Error(err)
					break
				}
				err = commands.EmailLocation(ctx, int64(userID))
				if err != nil {
					log.Error(err)
					break
				}
				log.Info("success EmailLocation")
				break
			default:
				log.Warnf("not defined: %s", message.Name)
				break
			}

			_ = msg.Ack(true)
			ctx.Close()
		}
	}()

	log.Info(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}
