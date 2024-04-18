package rmq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func (conn *RMQConnection) SendMessage(
	ctx context.Context,
	eventBody any,
	entity string,
	event string,
) error {
	ch, err := conn.conn.Channel()
	if err != nil {
		fmt.Println("failed to open channel")
		return err
	}

	body, err := json.Marshal(eventBody)
	if err != nil {
		fmt.Println("failed to send message")
		return err
	}
	err = ch.PublishWithContext(ctx,
		"notifications",
		"library."+entity+"."+event,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		return err
	}
	return nil
}
