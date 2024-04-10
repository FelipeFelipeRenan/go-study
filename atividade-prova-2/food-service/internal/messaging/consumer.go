package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"foods/internal/repository"
	"log"

	"github.com/streadway/amqp"
)

type MessageHandler struct {
	ch   *amqp.Channel
	repo *repository.FoodRepository
}

type OrderCreatedMessage struct {
	FoodID   int `json:"food_id"`
	Quantity int `json:"quantity"`
}

const (
	queueName    = "order_created_queue"
	exchangeName = "order_exchange"
	routingKey   = "order_created"
)

func NewMessageHandler(ch *amqp.Channel, repo *repository.FoodRepository) *MessageHandler {
	return &MessageHandler{ch: ch, repo: repo}
}

func (h *MessageHandler) ConsumeOrderCreated(ctx context.Context) error {

	err := h.ch.ExchangeDeclare(
		exchangeName,
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

	q, err := h.ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = h.ch.QueueBind(
		q.Name,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	mshs, err := h.ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range mshs {
			var orderMsg OrderCreatedMessage
			if err := json.Unmarshal(msg.Body, &orderMsg); err != nil {
				fmt.Println("Error decoding message", err)
				continue
			}

			err := h.repo.UpdateFoodQuantity(ctx, orderMsg.FoodID, orderMsg.Quantity)
			if err != nil {
				log.Println("Erro ao atualizar a quantidade de alimentos:", err)
			}
		}

	}()
	return nil
}
