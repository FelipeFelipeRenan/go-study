package main

import (
	"log"
	"net/http"
	"order-service/internals/handlers"
	"order-service/internals/models"
	"order-service/internals/repository"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Inicialização do banco de dados
	db, err := gorm.Open(postgres.Open("host=db_orders user=postgres password=1234 dbname=orders_db port=5432 sslmode=disable"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados", err)
	}
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// Verificação e migração do modelo de Order
	if !db.Migrator().HasTable(&models.Order{}) {
		log.Println("Tabela Order não existe, criando...")
		if err := db.AutoMigrate(&models.Order{}); err != nil {
			log.Fatal("Erro ao migrar modelo", err)
		}
		log.Println("Tabela Order criada com sucesso")
	}

	// Conexão com RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal("Erro ao conectar ao RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Erro ao criar canal", err)
	}
	defer ch.Close()

	log.Println("Conectado ao RabbitMQ")

	// Inicialização do repositório e handlers
	orderRepo := repository.NewOrderRepository(db)
	orderHandler := handlers.NewOrderHandler(orderRepo, ch)

	// Configuração do servidor HTTP
	router := gin.Default()
	router.POST("/orders", orderHandler.CreateOrder)

	log.Println("Servidor iniciado em http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
