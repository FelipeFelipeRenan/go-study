package main

import (
	"evento-service/internal/handler"
	"evento-service/internal/models"
	"evento-service/internal/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {
	db, err := gorm.Open(postgres.Open("host=db user=postgres password=1234 dbname=events_db port=5432 sslmode=disable"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados", err)
	}
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// verificar se tabela de eventos nao existe
	if !db.Migrator().HasTable(&models.Event{}) {
		log.Println("Tabela de eventos")
		err = db.AutoMigrate(&models.Event{})
		if err != nil {
			log.Fatal("Erro ao migrar modelo", err)
		}
		log.Println("Tabela 'events' criada com sucesso!")
	}

	// Executar a migração e o seed
	if err := db.Exec(`INSERT INTO events (title, description, location, start_at, end_at)
        VALUES ('Evento 1', 'Descrição do Evento 1', 'Local do Evento 1', '2024-04-01 10:00:00', '2024-04-01 12:00:00'),
               ('Evento 2', 'Descrição do Evento 2', 'Local do Evento 2', '2024-04-05 15:00:00', '2024-04-05 17:00:00');`).Error; err != nil {
		log.Fatal("Erro ao inserir dados de seed:", err)
	}
	log.Println("Migração e seed concluídos com sucesso!")

	eventRepo := repository.NewEventRepository(db)
	eventHandler := handler.NewEventHandler(eventRepo)

	router := gin.Default()

	router.POST("/events", eventHandler.CreateEvent)
	router.GET("/events/:id", eventHandler.GetEventByID)
	router.DELETE("/events/:id", eventHandler.DeleteEvent)
	router.PUT("/events/:id", eventHandler.UpdateEvent)

	router.Use()

	log.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
