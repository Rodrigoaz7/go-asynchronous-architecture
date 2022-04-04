// https://developer.okta.com/blog/2021/04/23/elasticsearch-go-developers-guide#how-to-build-a-console-menu-in-go
// https://karmi.github.io/gotalks/go-elasticsearch-files/saved_resource.html
// https://www.youtube.com/watch?v=3arH8SgCdIs
// https://x-team.com/blog/set-up-rabbitmq-with-docker-compose/
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
