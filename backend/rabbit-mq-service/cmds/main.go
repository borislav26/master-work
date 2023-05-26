package main

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"rabbit-mq-service/event"
	"time"
)

func main() {
	rabbitMQ, err := connect()
	if err != nil {
		log.Fatalln("Failed to connect to rabbitMQ service", err)
	}
	defer rabbitMQ.Close()
	consumer, err := event.NewConsumer(rabbitMQ, "master-work")
	if err != nil {
		log.Fatalln("Failed to initialize consumer", err)
	}

	emitter, err := event.NewEmitter(rabbitMQ, "master-work")
	if err != nil {
		log.Fatalln("Failed to initialize emitter", err)
	}

	ch, err := rabbitMQ.Channel()
	defer ch.Close()
	conf := make(chan amqp091.Confirmation)

	go func() {
		for {
			select {
			case <-conf:
				log.Println("this is confirmation of message")
				continue
			}
		}
	}()

	emitter.Push("this is event", "log.INFO")
	consumer.Listen([]string{"log.ERROR", "log.INFO", "log.WARNING"}, conf)

}

func connect() (*amqp091.Connection, error) {
	var tryCounts int64
	backOff := time.Second
	var connection *amqp091.Connection
	for {
		c, err := amqp091.Dial("amqp://guest:guest@localhost")
		if err != nil {
			log.Println("Trying to connect to rabbitmq service")
			tryCounts++
		} else {
			connection = c
			break
		}
		if tryCounts > 5 {
			log.Println("Failed to connect to rabbitMQ service because ", err.Error())
			return nil, err
		}
		backOff = time.Duration(math.Pow(float64(tryCounts), 2)) * time.Second
		time.Sleep(backOff)
	}
	return connection, nil
}
