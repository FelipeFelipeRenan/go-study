package handlers

import (
	"net/http"
	rabbitmq "order-service/internals/messaging"
	"order-service/internals/models"
	"order-service/internals/repository"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)


type OrderHandler struct {
	repo *repository.OrderRepository
	ch *amqp.Channel
}

func NewOrderHandler(repo *repository.OrderRepository, ch *amqp.Channel) *OrderHandler{
	return &OrderHandler{repo: repo, ch: ch}
}

func (h *OrderHandler) CreateOrder(c *gin.Context)  {
	var order models.Order
	if err := c.BindJSON(&order); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to create Order"})
		return
	}

	if err := h.repo.CreateOrder(c, &order); err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	err := rabbitmq.PublishOrderCreated(h.ch, &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error to publish order creation"})
		return
	}
	c.JSON(http.StatusCreated, order)
}