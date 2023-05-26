package event

import "github.com/rabbitmq/amqp091-go"

func declareExchange(channel *amqp091.Channel) error {
	return channel.ExchangeDeclare("master-work", amqp091.ExchangeDirect, true, false, false, false, nil)
}

func declareQueue(channel *amqp091.Channel, queueName string) (amqp091.Queue, error) {
	return channel.QueueDeclare("master-work", true, false, false, false, nil)
}
