package handlers

import (
	"foods/internal/models"
	"foods/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	repo *repository.FoodRepository
}

func NewFoodHandler(repo *repository.FoodRepository) *FoodHandler{
	return &FoodHandler{repo: repo}
}


func (h *FoodHandler) CreateFood(c *gin.Context){
	var food models.Food
	if err := c.BindJSON(&food); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to create Food"})
		return
	}
	if err := h.repo.CreateFood(c, &food); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Food"})
	} 
	c.JSON(http.StatusCreated, food)
}

func (h *FoodHandler) GetAllFoods(c *gin.Context){
	foods, err := h.repo.GetAllFoods(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Erro ao buscar alimentos"})
		return
	}

	c.JSON(http.StatusOK, foods)
}

func (h *FoodHandler) GetAllFoodsByCategory(c *gin.Context){
	category := c.Query("category")

	if category == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"O parâmetro 'category' é obrigatório"})
		return
	}

	foods, err := h.repo.GetAllFoodsByCategory(c,category)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "Erro ao buscar alimentos por categoria"})
		return
	}

	if len(foods) == 0{
		c.JSON(http.StatusNotFound, gin.H{"error":"Não foram encontrada alimentos para a categoria indicada"})
		return
	}
	c.JSON(http.StatusOK, foods)
}


func (h *FoodHandler) GetFoodsByID(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Food ID"})
	}
	Food, err := h.repo.GetFoodsByID(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
		return
	}
	c.JSON(http.StatusOK, Food)
}

func (h *FoodHandler) UpdateFood(c *gin.Context){
	var Food models.Food
	if err := c.BindJSON(&Food); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid food ID"})
		return
	}
	existingFood, err := h.repo.GetFoodsByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"Food not found"})
		return
	}
	
	existingFood.Name = Food.Name
	existingFood.Category = Food.Category
	existingFood.Quantity = Food.Quantity
	existingFood.Price = Food.Price
	existingFood.ExpirationAt = Food.ExpirationAt

	if err := h.repo.UpdateFood(c, existingFood); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update food"})
		return
	}
	c.JSON(http.StatusOK, existingFood)
}


func (h *FoodHandler) DeleteFood(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Invalid Food ID"})
		return
	}
	err = h.repo.DeleteFood(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to delete food"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Food deleted successfully"})
}

