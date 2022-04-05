package main

import (
	config "go-asynchronous-architecture/mailService/config"
	mail "go-asynchronous-architecture/mailService/mail"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	config.Init()
	var rabbitmqUrl = os.Getenv("RABBITMQ_SERVER")
	conn, err := amqp.Dial(rabbitmqUrl)
	if err != nil {
		log.Fatalf("Error connecting to rabbitmq: '%s'", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error creating channel rabbitmq: '%s'", err)
	}
	defer ch.Close()

	var rabbitmqQueue = os.Getenv("RABBITMQ_QUEUE_NAME")
	q, _ := ch.QueueDeclare(
		rabbitmqQueue, // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)

	msgs, _ := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			_ = mail.SendMail(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
