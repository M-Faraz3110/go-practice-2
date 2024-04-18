package rmq

import "github.com/rabbitmq/amqp091-go"

type RMQConnection struct {
	conn *amqp091.Connection
}

func NewRMQConnection(connString string) *RMQConnection {
	conn, err := amqp091.Dial(connString)
	if err != nil {
		panic("unable to dial rabbit mq connection")
	}
	return &RMQConnection{conn: conn}
}
