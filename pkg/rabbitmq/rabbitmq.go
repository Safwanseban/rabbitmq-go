package rabbitmq

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateQueue(ch *amqp.Channel) error {

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := "i am safwan"
	err = ch.PublishWithContext(

		ctx, "", q.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil

}
