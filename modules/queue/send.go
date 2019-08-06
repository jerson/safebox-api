package queue

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"safebox.jerson.dev/api/modules/util"
)

//SendTask ...
func SendTask(task Task) (string, error) {

	token := util.UniqueID()
	priority := task.Message.Priority
	if priority < 1 {
		priority = 4
	}

	conn, err := StartServer()
	if err != nil {
		return token, err
	}
	defer conn.Close()

	ch, err := GetChannel(conn)
	if err != nil {
		return token, err
	}
	defer ch.Close()

	task.Message.Token = token

	body, err := json.Marshal(task.Message)
	if err != nil {
		return token, err
	}

	q, err := GetQueue(ch, task.QueueName)
	if err != nil {
		return token, err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			Priority:    priority,
			ContentType: "application/json",
			Body:        body,
		})

	return token, err
}

// SendEmailLocationTask ...
func SendEmailLocationTask(userID string) (string, error) {
	task := Task{
		QueueName: "default",
		Message: Message{
			Priority: 1,
			Name:     "email:location",
			Params: map[string]string{
				"userID": userID,
			},
		},
	}
	return SendTask(task)

}
