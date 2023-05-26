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

type Emitter struct {
	Connection *amqp091.Connection
	QueueName  string
}

func NewEmitter(connection *amqp091.Connection, queueName string, channel *amqp091.Channel) (Emitter, error) {
	consumer := Emitter{
		Connection: connection,
		QueueName:  queueName,
	}
	err := consumer.setup()
	if err != nil {
		return Emitter{}, err
	}

	err = declareExchange(channel)
	if err != nil {
		return Emitter{}, err
	}
	
	return consumer, nil
}

func (e *Emitter) setup() error {
	channel, err := e.Connection.Channel()
	if err != nil {
		return err
	}
	return declareExchange(channel)
}

func declareExchange(channel *amqp091.Channel) error {
	return channel.ExchangeDeclare("master-work", amqp091.ExchangeDirect, true, false, false, false, nil)
}

func (e *Emitter) Push(event, severity string) error {
	ch, err := e.Connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.Tx()
	if err != nil {
		return err
	}

	ctx := context2.Background()
	err = ch.PublishWithContext(ctx, "master-work", severity, false, false, amqp091.Publishing{
		Body: []byte(event),
	})
	if err != nil {
		return err
	}

	err = ch.TxCommit()
	if err != nil {
		return err
	}

	return nil
}