package main

import (
	"evento-service/internal/handler"
	"evento-service/internal/models"
	"evento-service/internal/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Configuração da conexão com o banco de dados
	db, err := gorm.Open(postgres.Open("host=db_eventos user=postgres password=1234 dbname=eventos_db port=5432 sslmode=disable"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados", err)
	}
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// Verificar se a tabela de eventos não existe e criar se necessário
	if !db.Migrator().HasTable(&models.Event{}) {
		log.Println("Tabela de eventos não encontrada. Criando...")
		err = db.AutoMigrate(&models.Event{})
		if err != nil {
			log.Fatal("Erro ao migrar modelo", err)
		}
		log.Println("Tabela 'events' criada com sucesso!")
	}

	// Executar a migração e o seed
	if err := db.Exec(`INSERT INTO events (title, description, location, start_at, end_at)
        VALUES ('Evento 1', 'Descrição do Evento 1', 'Local do Evento 1', '2024-04-01 10:00:00', '2024-04-01 12:00:00'),
               ('Evento 2', 'Descrição do Evento 2','987654321', NOW(), NOW() 'Local do Evento 2', '2024-04-05 15:00:00', '2024-04-05 17:00:00');`).Error; err != nil {
		log.Fatal("Erro ao inserir dados de seed:", err)
	}
	log.Println("Migração e seed concluídos com sucesso!")

	// Inicializar repositório e handler
	eventRepo := repository.NewEventRepository(db)
	eventHandler := handler.NewEventHandler(eventRepo)

	// Configuração do roteamento com Gin
	router := gin.Default()
	router.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Definição das rotas
	router.POST("/events", eventHandler.CreateEvent)
	router.GET("/events/:id", eventHandler.GetEventByID)
	router.DELETE("/events/:id", eventHandler.DeleteEvent)
	router.PUT("/events/:id", eventHandler.UpdateEvent)

	log.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
