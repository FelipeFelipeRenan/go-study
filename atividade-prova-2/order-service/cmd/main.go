package main

import (
	"log"
	"order-service/internals/handlers"
	"order-service/internals/models"
	"order-service/internals/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=db_orders user=postgres password=1234 dbname=orders_db port=5432 sslmode=disable"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados", err)

		defer func(){
			if sqlDB, err := db.DB(); err == nil{
				sqlDB.Close()
			}
		}()

		if !db.Migrator().HasTable(&models.Order{}){
			log.Println("Order table")
			err = db.AutoMigrate(&models.Order{})
			if err != nil {
				log.Fatal("Erro ao migrar modelo", err)
			}
			log.Println("Tabela Order criada com sucesso")
		}

		ordeRepo := repository.NewOrderRepository(db)
		orderHandler := handlers.NewOrderHandler(ordeRepo)

		router := gin.Default()

		router.POST("/orders", orderHandler.CreateOrder)
	}
}