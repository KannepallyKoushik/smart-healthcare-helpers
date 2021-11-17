package main

import (
	"fmt"

	"github.com/KannepallyKoushik/rabbitmq-test/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	fmt.Println("Hey User!")

	rmq := rabbitmq.NewRabbitMQService()

	app := App{Rmq: rmq}

	err := app.Rmq.Connect()

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error setting up application.")
		fmt.Println(err)
	}
}
