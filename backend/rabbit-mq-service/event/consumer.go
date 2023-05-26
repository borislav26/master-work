package event

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type Consumer struct {
	Connection *amqp091.Connection
	QueueName  string
}

func NewConsumer(connection *amqp091.Connection, queueName string) (Consumer, error) {
	consumer := Consumer{
		Connection: connection,
		QueueName:  queueName,
	}
	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}
	return consumer, nil
}

func (c *Consumer) setup() error {
	channel, err := c.Connection.Channel()
	if err != nil {
		return err
	}
	return declareExchange(channel)
}

func (c *Consumer) Listen(topics []string, conf chan amqp091.Confirmation) error {
	ch, err := c.Connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := declareQueue(ch, "master-work")
	if err != nil {
		return err
	}

	for _, topic := range topics {
		err = ch.QueueBind(queue.Name, topic, "master-work", false, nil)
		if err != nil {
			return err
		}
	}

	messages, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	//conf1 := make(chan amqp091.Confirmation)

	forever := make(chan bool)
	go func() {
		for d := range messages {
			log.Println("this is body: %v", string(d.Body))
			d.Ack(true)
			log.Println("this is id", d.MessageId)
			log.Println("this is before", conf)
			ch.NotifyPublish(conf)
			log.Println("this is after", conf)
			//conf1 <- amqp091.Confirmation{DeliveryTag: 1, Ack: true}
		}
	}()
	<-forever
	return nil
}
