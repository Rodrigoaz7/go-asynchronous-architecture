package main

import (
	config "go-asynchronous-architecture/consumer/config"
	kafka "go-asynchronous-architecture/consumer/messages/kafka"
)

func main() {
	config.Init()
	kafka.ListenMessages()
}
