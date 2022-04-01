package kafka

import (
	"fmt"
	"os"

	controller "api-go-elasticsearch/consumer/controllers/pix"
	rabbitmq "api-go-elasticsearch/consumer/messages/rabbitmq"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ListenMessages() {
	var kafkaServer string = os.Getenv("KAFKA_BOOTSTRAP_SERVER")
	var kafkaGroupId string = os.Getenv("KAFKA_GROUP_ID")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"group.id":          kafkaGroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	var kafkaTopic string = os.Getenv("KAFKA_TOPIC")
	c.SubscribeTopics([]string{kafkaTopic, "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			errorToPersist := controller.PersistData(msg.Value)
			if errorToPersist == nil {
				rabbitmq.PublishNotification(msg.Value)
			}
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
