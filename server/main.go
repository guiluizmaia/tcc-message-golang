package main

import (
	"encoding/json"
	"log"

	Service "github.com/guiluizmaia/tcc-message-golang/server/services"
	"github.com/streadway/amqp"
)

func main() {
	connectRabbitMQ, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)

	}

	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer channelRabbitMQ.Close()

	_, err = channelRabbitMQ.QueueDeclare(
		"queueRequest", // queue name
		true,           // durable
		false,          // auto delete
		false,          // exclusive
		false,          // no wait
		nil,            // arguments
	)

	if err != nil {
		log.Fatalf("Could not queue declare: %v", err)
	}

	messages, err := channelRabbitMQ.Consume(
		"queueRequest", // queue name
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no local
		false,          // no wait
		nil,            // arguments
	)

	if err != nil {
		log.Fatalf("Could not queue consume: %v", err)
	}

	service := Service.NewUserService()

	forever := make(chan bool)

	for message := range messages {
		userCreated, _ := service.CreateUser(message.Body)
		userByte, _ := json.Marshal(userCreated)

		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(userByte),
		}

		channelRabbitMQ.Publish(
			"",
			"queueResponse",
			false,
			false,
			message,
		)
	}

	<-forever
}
