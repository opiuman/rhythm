package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"gopkg.in/mgo.v2"
)

func main() {
	var (
		ampqUrl = flag.String("amqpUrl", "amqp://guest:guest@localhost:5672/", "RabbitMQ Url")
		qName   = flag.String("qName", "API", "queueName")
		msUrl   = flag.String("msUrl", "localhost:27017", "mongodb server Url")
		dbName  = flag.String("dbName", "msgdb", "mongodb name")
		cName   = flag.String("cName", "msg", "mongo collection name")
	)
	flag.Parse()
	conn, err := amqp.Dial(*ampqUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		*qName, // name
		false,  // durable
		false,  // delete when usused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	session, err := mgo.Dial(*msUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(*dbName).C(*cName)

	go func() {
		for d := range msgs {
			log.Printf("Received a message")

			var m map[string]interface{}
			err := json.Unmarshal(d.Body, &m)
			if err != nil {
				log.Println("Failed to unmarshall the message")
			} else {
				err = c.Insert(m)
				if err != nil {
					log.Println("Failed to insert the message into Mongodb")
				}
			}

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
