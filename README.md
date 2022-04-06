# go-asynchronous-architecture
Recently I've been studying a little bit about asynchronous systems and Service-Oriented Architecture. Then I decided to make a simple small project to establish my knowledge. 

This project is a small and very simple Pix simulator.  Here we can use softwares as kafka and rabbitmq to comunicate asynchronous between three microservices.

I'm using Elasticsearch as my database only for study purperses, but of course, it's not the ideal for real scenarios. 

The microservices are all developed in Golang language and I'm using docker-compose file to deploy all the libraries that I need.

# Table of contents

<!--ts-->
   * [About](#go-asynchronous-architecture)
   * [Table of contents](#table-of-contents)
   * [Technologies](#technologies)
   * [Features](#features)
   * [Architecture](#architecture)
   * [How to use](#how-to-use)
     * [Environment variables](#environment-variables)
     * [Running](#running)
     * [Monitoring](#monitoring)
<!--te-->

# Technologies
- [Golang](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Elasticsearch](https://www.elastic.co/pt/what-is/elasticsearch)
- [Kibana](https://www.elastic.co/pt/kibana/)
- [Kafka](https://kafka.apache.org/)
- [RabbitMq](https://www.rabbitmq.com/)

# Features

- [x] Pix Transaction Register Through Kafka
- [x] Notification of Event with RabbitMQ
- [x] Email Service with Queue of events

# Architecture
![image](https://user-images.githubusercontent.com/27520422/161871813-6aaf15e9-2ea5-4b21-a316-cfeb27a62e1d.png)

Basically, the database used to store the data is elasticsearch. As said before, it's only for studies purpeses. The services are called "publisher", "consumer" and "mailservice". The publisher is the only one open to recieve a http post request from client. When it does gets a post request, the publisher send the data to a kafka broker. As this is a simple study projetct, I created only one topic with one partition (just one consumer) and one replication factory. 

When the stream data reaches the broker, the "consumer" service recieve this message. The first operation is to persist the data to the elasticsearch. If everything goes well, then the consumer will send a notification message's event to a rabbitmq queue.

Finally, the mail service will recieve from the rabbitmq's queue the success message and it will fire the emails.

# How to use
You only need the docker installed in your machine. If you use windows, it's recommended that you use WSL console and install docker client in it.
```docker
docker-compose up
```

## Environment variables
The data configuration of all services are defined in a file located at config/init.go. It should have this format for all three services:
```golang
package config

import "os"

func Init() {
	os.Setenv("LOCAL_CONSUMER_PORT", ":8080")
	os.Setenv("LOCAL_PUBLISHER_PORT", ":8090")
	os.Setenv("KAFKA_BOOTSTRAP_SERVER", "host.docker.internal:9094")
	os.Setenv("KAFKA_GROUP_ID", "kafka's group id here")
	os.Setenv("KAFKA_TOPIC", "kafka's topic here")
	os.Setenv("RABBITMQ_SERVER", "amqp://guest:guest@rabbitmq:5672/")
	os.Setenv("RABBITMQ_QUEUE_NAME", "rabbitmq's queue here")
	os.Setenv("ELASTICSEARCH_INDEX_NAME", "elasticsearch's index here")
	os.Setenv("ELASTICSEARCH_HOST", "http://elasticsearch01")
	os.Setenv("ELASTICSEARCH_PORT", ":9200")
}
```
Each service uses some of the variables above, you should use only the ones that the service needs.

Ps: In the Dockerfile files, the images are being exposed to specific ports. You need to adapte that to your local environment variables.

## Running
After all docker's conteiners are up, you can init the project accessing by your LOCAL-PUBLISHER-PORT. A HTTP GET requisition will return a json array with all pix transactions registered in elasticsearch database. 

To save some pix transaction into database, you only need to make a HTTP POST request to your publish server with the body: 

```json
  {
    "target_account": "",
    "source_account": "",
    "target_mail": "",
    "source_mail": "",
    "value": 
}
```

If everything goes ok, the data is saved into elasticsearch database and a mail is fired to the target and source emails.

## Monitoring
Some images imported in the docker-compose file (such as kibana01, control-center and rabbitmq it self) can be used to monitoring the data path locally. When a POST request is fired into publisher service, you can monitoring de Kafka Broker through url http://localhost:9021.

Kafka's Broker Monitoring:
![kafka](https://user-images.githubusercontent.com/27520422/161872793-e3d9e009-895e-4890-960d-03ffd5bd2db6.png)

Once the consumer's service recieve the stream data, it will save it into elasticsearch database. You can see all the data with kibana through url http://localhost:5601. After that, you can see the notification message sent into rabbitmq queue though url http://localhost:15672.

Kibana devtools:
![kibana](https://user-images.githubusercontent.com/27520422/161872828-bf94515b-b4a0-448d-9c9f-20a456786052.png)

RabbitMQ's Queue Monitoring:
![rabbitmq](https://user-images.githubusercontent.com/27520422/161872873-be12d0dd-29a6-4c33-9b75-1505724da0be.png)

Notification Email:
![mail](https://user-images.githubusercontent.com/27520422/161872893-257b69c0-2faf-4e6b-9c8e-4e10b4c7a8db.png)







