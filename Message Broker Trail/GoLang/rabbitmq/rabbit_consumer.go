package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Service interface {
	Connect() error
	Consume()
}

type RabbitMQ struct {
	Conn *amqp.Connection
}

func (r *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ")

	var err error
	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Connected to RabbitMQ")

	ch, err = r.Conn.Channel()
	if err != nil {
		return err
	}

	_, err = ch.queueDeclare(
		"order_notify",
		true,
		false,
		false,
		false,
		nil,
	)

	return nil
}

// Consume reads messages - from message queue
func (r *RabbitMQ) Consume() {
	msgs, err := r.Channel.Consume(
		"order_notify",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	for msg := range msgs {
		fmt.Printf("Recieved message: %s \n", msg.Body)
	}
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
