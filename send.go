package main

import (
	"log"

	"github.com/Safwanseban/rabbitmq-go/pkg/configs"
	"github.com/Safwanseban/rabbitmq-go/pkg/rabbitmq"
)

func main() {
	ch, err := configs.RabbitMqConn()

	if err != nil {
		log.Println(err.Error())
	}
	defer ch.Close()
	if err := rabbitmq.CreateQueue(ch); err != nil {
		log.Println("error creating queue", err)
	}

}
