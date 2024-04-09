package messaging

import (
	"fmt"

	"github.com/streadway/amqp"
)

func PublishFoodCreated(conn *amqp.Connection, foodID int) error  {
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
	body := fmt.Sprintf("New food created with ID: %d", foodID)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})

		if err != nil {
			return err
		}
		return nil
}