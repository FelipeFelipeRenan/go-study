package messaging

import (
	"log"
	"order-service/internals/models"

	"github.com/streadway/amqp"
)

const (
	orderExchangeName  = "order_exchange"
	orderQueueName = "order_created"
)

func PublishOrderCreated(ch *amqp.Channel, order *models.Order) error {
	err := ch.ExchangeDeclare(
		orderExchangeName, 
		"direct", 
		true, 
		false, 
		false, 
		false, 
		nil,
	)
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		orderQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		orderQueueName,
		"",
		orderExchangeName,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	orderJSON, err := order.ToJSON()
	if err != nil {
		log.Fatal("Erro ao converter pedido para JSON:", err)
	}
	err = ch.Publish(
		orderExchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: []byte(orderJSON),
		},
	)
	if err != nil {
		return err
	}
	return nil	
}
