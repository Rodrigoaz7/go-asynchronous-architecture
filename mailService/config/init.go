package config

import "os"

func Init() {
	os.Setenv("HOST", "http://localhost")
	os.Setenv("LOCAL_CONSUMER_PORT", ":8080")
	os.Setenv("LOCAL_PUBLISHER_PORT", ":8090")
	os.Setenv("KAFKA_BOOTSTRAP_SERVER", "host.docker.internal:9094")
	os.Setenv("KAFKA_GROUP_ID", "1")
	os.Setenv("KAFKA_TOPIC", "pix_transactions")
	os.Setenv("RABBITMQ_SERVER", "amqp://guest:guest@rabbitmq:5672/")
	os.Setenv("RABBITMQ_QUEUE_NAME", "pix-queue")
	os.Setenv("ELASTICSEARCH_INDEX_NAME", "pix")
	os.Setenv("ELASTICSEARCH_HOST", "http://elasticsearch01")
	os.Setenv("ELASTICSEARCH_PORT", ":9200")
	os.Setenv("MAIL_HOST", "smtp.gmail.com")
	os.Setenv("MAIL_PORT", "587")
	os.Setenv("MAIL_USERNAME", "rodrigo.aze7@gmail.com")
	os.Setenv("MAIL_PASSWORD", "fumvmcnbqglsnpvi")
}
