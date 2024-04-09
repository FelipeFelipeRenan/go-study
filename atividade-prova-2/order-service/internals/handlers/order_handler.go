package handlers

import (
	"net/http"
	"order-service/internals/models"
	"order-service/internals/repository"

	"github.com/gin-gonic/gin"
)


type OrderHandler struct {
	repo *repository.OrderRepository
}

func NewOrderHandler(repo *repository.OrderRepository) *OrderHandler{
	return &OrderHandler{repo: repo}
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
	c.JSON(http.StatusCreated, order)
}