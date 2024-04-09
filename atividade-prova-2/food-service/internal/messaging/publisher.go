package messaging

import (
	"foods/internal/models"

	"github.com/streadway/amqp"
)

func PublishFoodCreated(ch *amqp.Channel, food *models.Food) error {

	exchangeName := "food_exchange"

	err := ch.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)

	if err != nil {
		return err
	}

	queueName := "food_created"
	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)

	err = ch.QueueBind(queueName, "", exchangeName, false, nil)
	if err != nil {
		return err
	}

	msgBody := []byte(food.Name)
	err = ch.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msgBody),
		})

	if err != nil {
		return err
	}
	return nil
}
