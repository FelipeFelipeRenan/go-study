package messaging

import (
	"log"

	"github.com/streadway/amqp"
)

func ConsumeFoodCreated(conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"food_created",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for msg := range msgs {
			log.Printf("Received message: %s", msg.Body)
		}
	}()

	log.Println("Consumer started")
	select {}

	return nil
}
