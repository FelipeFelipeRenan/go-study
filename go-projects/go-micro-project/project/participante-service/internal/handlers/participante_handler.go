package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"participante-service/internal/models"
	"participante-service/internal/repository"

	// Importe o pacote swaggo para as anotações
)

// ParticipanteHandler gerencia os handlers dos Participanteos.
type ParticipanteHandler struct {
	repo *repository.ParticipanteRepository
}

// NewParticipanteHandler cria uma nova instância de ParticipanteHandler.
func NewParticipanteHandler(repo *repository.ParticipanteRepository) *ParticipanteHandler {
	return &ParticipanteHandler{repo: repo}
}

// @Summary Cria um novo Participanteo
// @Description Cria um novo Participanteo com os dados fornecidos
// @Tags Participanteos
// @Accept json
// @Produce json
// @Param requestBody body Participante true "Dados do Participanteo a ser criado"
// @Success 201 {object} Participante
// @Router /Participantes [post]
func (h *ParticipanteHandler) CreateParticipante(c *gin.Context) {
	var Participante models.Participante
	if err := c.BindJSON(&Participante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.repo.CreateParticipante(c, &Participante)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Participante"})
		return
	}
	c.JSON(http.StatusCreated, Participante)
}

// @Summary Retorna um Participanteo por ID
// @Description Retorna um Participanteo com base no ID fornecido
// @Tags Participanteos
// @Accept json
// @Produce json
// @Param id path int true "ID do Participanteo a ser retornado"
// @Success 200 {object} Participante
// @Router /Participantes/{id} [get]
func (h *ParticipanteHandler) GetParticipanteByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Participante ID"})
		return
	}
	Participante, err := h.repo.GetParticipanteByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Participante not found"})
		return
	}
	c.JSON(http.StatusOK, Participante)
}

// @Summary Atualiza um Participanteo
// @Description Atualiza um Participanteo com base nos dados fornecidos
// @Tags Participanteos
// @Accept json
// @Produce json
// @Param id path int true "ID do Participanteo a ser atualizado"
// @Param requestBody body Participante true "Novos dados do Participanteo"
// @Success 200 {object} Participante
// @Router /Participantes/{id} [put]
func (h *ParticipanteHandler) UpdateParticipante(c *gin.Context) {
	var Participante models.Participante
	if err := c.BindJSON(&Participante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.repo.UpdateParticipante(c, &Participante)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Participante"})
		return
	}
	c.JSON(http.StatusOK, Participante)
}

// @Summary Exclui um Participanteo
// @Description Exclui um Participanteo com base no ID fornecido
// @Tags Participanteos
// @Accept json
// @Produce json
// @Param id path int true "ID do Participanteo a ser excluído"
// @Success 200 {string} string "Participanteo excluído com sucesso"
// @Router /Participantes/{id} [delete]
func (h *ParticipanteHandler) DeleteParticipante(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Participante ID"})
		return
	}
	err = h.repo.DeleteParticipante(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Participante"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Participante deleted successfully"})
}