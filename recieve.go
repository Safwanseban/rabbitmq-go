package main

import (
	"log"

	"github.com/Safwanseban/rabbitmq-go/pkg/configs"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := configs.RabbitMqConn()

	if err != nil {
		log.Println(err.Error())
	}
	defer ch.Close()
	err = CreateQueue(ch)
	if err != nil {
		log.Println(err)
	}
}
func CreateQueue(ch *amqp.Channel) error {
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil

}
