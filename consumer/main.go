// https://developer.okta.com/blog/2021/04/23/elasticsearch-go-developers-guide#how-to-build-a-console-menu-in-go
// https://karmi.github.io/gotalks/go-elasticsearch-files/saved_resource.html
// https://www.youtube.com/watch?v=3arH8SgCdIs
// https://x-team.com/blog/set-up-rabbitmq-with-docker-compose/
package main

import (
	config "go-asynchronous-architecture/consumer/config"
	kafka "go-asynchronous-architecture/consumer/messages/kafka"
)

func main() {
	config.Init()
	kafka.ListenMessages()
}
