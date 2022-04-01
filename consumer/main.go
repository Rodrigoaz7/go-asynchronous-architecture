// https://developer.okta.com/blog/2021/04/23/elasticsearch-go-developers-guide#how-to-build-a-console-menu-in-go
// https://karmi.github.io/gotalks/go-elasticsearch-files/saved_resource.html
// https://www.youtube.com/watch?v=3arH8SgCdIs
// https://x-team.com/blog/set-up-rabbitmq-with-docker-compose/
package main

import (
	config "api-go-elasticsearch/consumer/config"
	kafka "api-go-elasticsearch/consumer/messages/kafka"
)

func main() {
	config.Init()
	kafka.ListenMessages()

	// msgs, err := rabbitmq.Connect()
	// if err != nil {
	// 	panic(err)
	// }

	// forever := make(chan bool)

	// go func() {
	// 	for d := range msgs {
	// 		log.Printf("Received a message: %s", d.Body)
	// 		controller.Create(d.Body)
	// 	}
	// }()

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// <-forever
}
