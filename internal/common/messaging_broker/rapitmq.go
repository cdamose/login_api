package messagingbroker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"login_api/internal/communication_svc/container"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	Message     string `json:"message"`
	PhoneNumber string `json:"phone_number"`
}
type MessageHandler func([]byte)

type RappitMQBroker struct {
	done        chan struct{}
	application container.Application
}

func NewRappitMQBroker(application container.Application) *RappitMQBroker {
	return &RappitMQBroker{
		done:        make(chan struct{}),
		application: application,
	}
}

func (rq *RappitMQBroker) Subscribe(topic string, handler MessageHandler) {
	conn, err := amqp.Dial("amqp://test:test@rabbitmq:5672/")
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		topic,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		failOnError(err, "Failed to declare a queue")
		return
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		failOnError(err, "Failed to register a consumer")
		return
	}

	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
			var data Message
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println("BEFORE SENING SMS")
			fmt.Println(data.Message)
			fmt.Println(data.PhoneNumber)
			i, err := rq.application.CommunicationApplication.SendSMS(context.Background(), "+919677892850", "from servcie")
			fmt.Println(i)
			fmt.Println(err)
		}
	}()
	log.Printf(" started reciving message on topic '%s'. To exit press CTRL+C", topic)
	<-rq.done
}

func (rq *RappitMQBroker) Stop() {
	close(rq.done)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
