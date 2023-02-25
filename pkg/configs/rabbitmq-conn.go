package configs

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func RabbitMqConn() (*amqp.Channel, error) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}


	return ch, nil

}
