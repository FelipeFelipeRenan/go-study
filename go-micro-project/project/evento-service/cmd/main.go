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
)

func main() {
	db, err := gorm.Open(postgres.Open("host=db user=postgres password=1234 dbname=events_db port=5432 sslmode=disable"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info),})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados", err)
	}
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	err = db.AutoMigrate(&models.Event{})
	if err != nil {
		log.Fatal("Erro ao migrar modelo", err)
	}

	eventRepo := repository.NewEventRepository(db)
	eventHandler := handler.NewEventHandler(eventRepo)

	router := gin.Default()

	router.POST("/events", eventHandler.CreateEvent)
	router.GET("/events/:id", eventHandler.GetEventByID)

	log.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
