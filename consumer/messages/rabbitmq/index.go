package rabbitmq

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func PublishNotification(dataToSend []byte) error {
	var rabbitmqUrl string = os.Getenv("RABBITMQ_SERVER")
	conn, err := amqp.Dial(rabbitmqUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	if err != nil {
		return err
	}
	defer ch.Close()

	var rabbitmqQueue string = os.Getenv("RABBITMQ_QUEUE_NAME")
	// We create a Queue to send the message to.
	q, err := ch.QueueDeclare(
		rabbitmqQueue, // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(dataToSend),
		},
	)

	// If there is an error publishing the message, a log will be displayed in the terminal.
	failOnError(err, "Failed to publish a message")
	if err != nil {
		return err
	}

	log.Printf(" [x] Congrats, sending message...")
	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
