package handlers

import (
	"context"
	"foods/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	service service.Service
}

func NewFoodHandler(s service.Service) *FoodHandler {
	return &FoodHandler{service: s}
}

func (h *FoodHandler) GetAllFoods(c *gin.Context) {

	foods, err := h.service.GetAllFoods(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch foods"})
		return
	}
	c.JSON(http.StatusOK, foods)
}
