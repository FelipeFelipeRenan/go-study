package main

import (
	"context"
	"foods/internal/database"
	"foods/internal/handlers"
	"foods/internal/repository"
	"foods/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	//"go-micro.dev/v4"
)

func main() {

	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}
	// Crie um repositório e manipulador de alimentos
	foodRepo := repository.NewFoodRepository(db)
	foodService := service.NewFoodService(foodRepo)
	foodHandler := handlers.NewFoodHandler(foodService)

	// Implemente um servidor HTTP do Gin
	router := gin.Default()

	setupRoutes(router, foodHandler)

	// Inicie o servidor Gin para lidar com as solicitações HTTP
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}


	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to run Gin server", err)
		}
	}()

	log.Println("Servidor HTTP iniciado em http://localhost:8080")

	// Aguarde um sinal de interrupção para encerrar o servidor
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("Encerrando o servidor...")

	// Encerre o servidor HTTP do Gin
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Failed to shutdown Gin server", err)
	}

	log.Println("Servidor HTTP do Gin encerrado com sucesso")

	log.Println("Servidor encerrado")
}

func setupRoutes(router *gin.Engine, foodHandler *handlers.FoodHandler){
	router.GET("/foods", foodHandler.GetAllFoods)
}
