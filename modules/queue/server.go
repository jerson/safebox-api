package queue

import (
	"github.com/streadway/amqp"
	"safebox.jerson.dev/api/modules/config"
)

//StartServer ...
func StartServer() (*amqp.Connection, error) {
	conn, err := amqp.Dial(config.Vars.RabbitMQ.Server)
	if err != nil {
		return nil, err
	}
	return conn, nil

}

//GetChannel ...
func GetChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	q, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return q, nil

}

//GetQueue ...
func GetQueue(ch *amqp.Channel, name string) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return q, err
	}
	return q, nil

}
