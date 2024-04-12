//	"go-micro.dev/v4"

package main

import (

	"foods/internal/handlers"
	"foods/internal/models"
	"foods/internal/repository"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"go-micro.dev/v4"
	"go-micro.dev/v4/transport"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	foods := []models.Food{
		{Name: "Lasanha", Category: "Massas", Quantity: 20, Price: 25.99, ExpirationAt: time.Now().AddDate(0, 1, 0)},
		{Name: "Sushi", Category: "Japonesa", Quantity: 30, Price: 30.50, ExpirationAt: time.Now().AddDate(0, 0, 15)},
		// Adicione outros alimentos aqui
	}

	db, err := gorm.Open(postgres.Open("host=db_foods user=postgres password=1234 dbname=foods_db port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados", err)
	}
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()
	
	// Drop and create the table
	db.Migrator().DropTable(&models.Food{})
	if err := db.AutoMigrate(&models.Food{}); err != nil {
		log.Fatal("Erro ao migrar modelo", err)
	}

	log.Println("Tabela 'food' criada com sucesso")
	for _, f := range foods {
		if err := db.Create(&f).Error; err != nil {
			log.Fatal("Erro ao inserir dados de seed para alimentos:", err)
		}
	}
	log.Println("Seed de dados para participantes concluído com sucesso!")

	// Crie um repositório e manipulador de alimentos
	foodRepo := repository.NewFoodRepository(db)
	foodHandler := handlers.NewFoodHandler(foodRepo)

	// Defina endpoints Go kit para suas operações
	endpoints := handlers.MakeEndpoints(foodHandler)

	// Crie um serviço Go kit que usa esses endpoints
	service := micro.NewService(
		micro.Name("foods"),
		micro.Version("latest"),
		micro.Address(":8080"),
	)
	service.Init()

	// Implemente um servidor HTTP do Gin
	router := gin.Default()

	// Roteie as solicitações HTTP para os endpoints do Go kit usando o httptransport
	router.POST("/foods", gin.WrapH(httptransport.NewServer(
		endpoints.CreateFoodEndpoint,
		handlers.DecodeCreateFoodRequest,
		handlers.EncodeResponse,
	)))

	router.GET("/foods/all", gin.WrapH(httptransport.NewServer(
		endpoints.GetAllFoodsEndpoint,
		handlers.DecodeGetAllFoodsRequest,
		handlers.EncodeResponse,
	)))

	router.GET("/foods/:id", gin.WrapH(httptransport.NewServer(
		endpoints.GetFoodsByIDEndpoint,
		handlers.DecodeGetFoodsByIDRequest,
		handlers.EncodeResponse,
	)))

	router.GET("/foods/", gin.WrapH(httptransport.NewServer(
		endpoints.GetAllFoodsByCategoryEndpoint,
		handlers.DecodeGetAllFoodsByCategoryRequest,
		handlers.EncodeResponse,
	)))

	router.PUT("/foods/:id", gin.WrapH(httptransport.NewServer(
		endpoints.UpdateFoodEndpoint,
		handlers.DecodeUpdateFoodRequest,
		handlers.EncodeResponse,
	)))

	router.DELETE("/foods/:id", gin.WrapH(httptransport.NewServer(
		endpoints.DeleteFoodEndpoint,
		handlers.DecodeDeleteFoodRequest,
		handlers.EncodeResponse,
	)))

	// Inicie o serviço Go kit em uma goroutine separada
	go func() {
		if err := service.Run(); err != nil {
			log.Fatal("Failed to run Go kit service", err)
		}
	}()

	// Inicie o servidor Gin para lidar com as solicitações HTTP
	if err := router.Run(":8081"); err != nil {
		log.Fatal("Failed to run Gin server", err)
	}
}
