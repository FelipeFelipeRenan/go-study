package handler

import (
	"net/http"
	"strconv"

	"evento-service/internal/models"
	"evento-service/internal/repository"
	"github.com/gin-gonic/gin"
	// Importe o pacote swaggo para as anotações
)

// EventHandler gerencia os handlers dos eventos.
type EventHandler struct {
	repo *repository.EventRepository
}

// NewEventHandler cria uma nova instância de EventHandler.
func NewEventHandler(repo *repository.EventRepository) *EventHandler {
	return &EventHandler{repo: repo}
}

// @Summary Cria um novo evento
// @Description Cria um novo evento com os dados fornecidos
// @Tags Eventos
// @Accept json
// @Produce json
// @Param requestBody body Event true "Dados do evento a ser criado"
// @Success 201 {object} Event
// @Router /events [post]
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.repo.CreateEvent(c, &event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}
	c.JSON(http.StatusCreated, event)
}

// @Summary Retorna um evento por ID
// @Description Retorna um evento com base no ID fornecido
// @Tags Eventos
// @Accept json
// @Produce json
// @Param id path int true "ID do evento a ser retornado"
// @Success 200 {object} Event
// @Router /events/{id} [get]
func (h *EventHandler) GetEventByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := h.repo.GetEventByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

// @Summary Atualiza um evento
// @Description Atualiza um evento com base nos dados fornecidos
// @Tags Eventos
// @Accept json
// @Produce json
// @Param id path int true "ID do evento a ser atualizado"
// @Param requestBody body Event true "Novos dados do evento"
// @Success 200 {object} Event
// @Router /events/{id} [put]
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	var event models.Event
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.repo.UpdateEvent(c, &event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}
	c.JSON(http.StatusOK, event)
}

// @Summary Exclui um evento
// @Description Exclui um evento com base no ID fornecido
// @Tags Eventos
// @Accept json
// @Produce json
// @Param id path int true "ID do evento a ser excluído"
// @Success 200 {string} string "Evento excluído com sucesso"
// @Router /events/{id} [delete]
func (h *EventHandler) DeleteEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	err = h.repo.DeleteEvent(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
