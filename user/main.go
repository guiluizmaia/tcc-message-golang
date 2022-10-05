package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
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

	_, err = channelRabbitMQ.QueueDeclare(
		"queueResponse",
		true,
		false,
		false,
		false,
		nil,
	)

	quant := 0
	amount := 205000
	bytesRequest := 0
	timeStart := time.Now()
	timeFinish := time.Now()
	success := 0
	error := 0

	for i := 0; i < amount; i++ {
		bodySend, _ := json.Marshal(map[string]string{
			"id":          uuid.New().String(),
			"name":        "nameFake",
			"lastName":    "lastNameFake",
			"age":         "50",
			"document":    "33333333333",
			"address":     "Rua X",
			"nationality": "Nationality Fake",
			"motherName":  "MotherNameFake",
			"fatherName":  "FatherName",
			"gender":      "Gender Fake",
			"birthday":    "12/07/2000",
			"email":       "test@test.com",
		})

		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(bodySend),
		}

		channelRabbitMQ.Publish(
			"",
			"queueRequest",
			false,
			false,
			message,
		)
	}

	messages, err := channelRabbitMQ.Consume(
		"queueResponse",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Could not queue consume: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			quant += 1
			// fmt.Print(string(message.Body), "\n")

			if message.Body != nil {
				success += 1
				bytesRequest = len(message.Body)
			}

			if message.Body == nil {
				error += 1
			}

			if quant == amount {
				timeFinish = time.Now()

				timeofRequests := timeFinish.Sub(timeStart)
				log.Printf("Amount of requests: %d", amount)
				log.Printf("Bytes of one request: %d", bytesRequest)
				log.Printf("Time of all requests: %s", timeofRequests)
				log.Printf("Errors: %d", error)
				log.Printf("Success: %d", success)
			}

		}
	}()

	<-forever
}
