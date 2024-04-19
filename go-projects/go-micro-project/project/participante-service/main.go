// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"participante-service/internal/handlers"
	"participante-service/internal/models"
	"participante-service/internal/repository"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=db_participantes user=postgres password=1234 dbname=participantes_db port=5432 sslmode=disable"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados", err)
	}
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	if !db.Migrator().HasTable(&models.Participante{}) {
		log.Println("Tabela de participantes")
		err = db.AutoMigrate(&models.Participante{})
		if err != nil {
			log.Fatal("Erro ao migrar modelo", err)
		}
		log.Println("Tabela 'participantes' criada com sucesso!")
	// Seed de dados para participantes
	if err := db.Exec(`
		INSERT INTO participantes (name, email, phone)
		VALUES ('Participante 1', 'participante1@example.com'),
		       ('Participante 2', 'participante2@example.com');
	`).Error; err != nil {
		log.Fatal("Erro ao inserir dados de seed para participantes:", err)
	}
	log.Println("Seed de dados para participantes concluído com sucesso!")
	}

	participantRepo := repository.NewParticipanteRepository(db)
	participantHandler := handlers.NewParticipanteHandler(participantRepo)

	router := gin.Default()

	router.POST("/participants", participantHandler.CreateParticipante)
	router.GET("/participants/:id", participantHandler.GetParticipanteByID)
	router.DELETE("/participants/:id", participantHandler.DeleteParticipante)
	router.PUT("/participants/:id", participantHandler.UpdateParticipante)

	log.Println("Servidor iniciado em http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
