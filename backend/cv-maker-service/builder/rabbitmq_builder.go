package builder

import (
	context2 "context"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"time"
)

func ConnectToRabbitMQ() (*amqp091.Connection, error) {
	var tryCounts int64
	backOff := time.Second
	var connection *amqp091.Connection
	for {
		c, err := amqp091.Dial("amqp://guest:guest@rabbitmq")
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

func GetChannelAndConnectToRabbitMQ() *Emitter {
	connection, err := ConnectToRabbitMQ()
	if err != nil {
		log.Fatalln("Failed to initialize rabbitMQ connection")
	}

	channel, err := connection.Channel()
	if err != nil {
		log.Fatalln("Failed to get channel from connection")
	}

	emitter, err := NewEmitter(channel)
	if err != nil {
		log.Fatalln("Failed to initialize emitter")
	}

	return &emitter
}

type Emitter struct {
	Connection *amqp091.Channel
}

func NewEmitter(channel *amqp091.Channel) (Emitter, error) {
	consumer := Emitter{
		Connection: channel,
	}

	err := declareExchange(channel)
	if err != nil {
		return Emitter{}, err
	}

	return consumer, nil
}

func declareExchange(channel *amqp091.Channel) error {
	return channel.ExchangeDeclare("master-work", amqp091.ExchangeDirect, true, false, false, false, nil)
}

func (e *Emitter) Push(ctx context2.Context, event []byte, severity string) error {
	defer e.Connection.Close()

	err := e.Connection.PublishWithContext(ctx, "master-work", severity, false, false, amqp091.Publishing{
		Body: []byte(event),
	})
	if err != nil {
		return err
	}

	return nil
}
