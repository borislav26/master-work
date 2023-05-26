package event

import (
	context2 "context"
	"github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	Connection *amqp091.Connection
	QueueName  string
}

func NewEmitter(connection *amqp091.Connection, queueName string) (Emitter, error) {
	consumer := Emitter{
		Connection: connection,
		QueueName:  queueName,
	}
	err := consumer.setup()
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
