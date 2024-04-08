package main

import (
	"foods/internal/handlers"
	"foods/internal/models"
	"foods/internal/repository"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	foods := []models.Food{
		{Name: "Lasanha", Category: "Massas", Quantity: 20, Price: 25.99, ExpirationAt: time.Now().AddDate(0, 1, 0)},          // Expira em 1 mês
		{Name: "Sushi", Category: "Japonesa", Quantity: 30, Price: 30.50, ExpirationAt: time.Now().AddDate(0, 0, 15)},         // Expira em 15 dias
		{Name: "Hambúrguer", Category: "Fast food", Quantity: 25, Price: 15.99, ExpirationAt: time.Now().AddDate(0, 0, 10)},   // Expira em 10 dias
		{Name: "Salada Caesar", Category: "Saladas", Quantity: 15, Price: 12.75, ExpirationAt: time.Now().AddDate(0, 0, 5)},   // Expira em 5 dias
		{Name: "Pizza", Category: "Italiana", Quantity: 40, Price: 18.99, ExpirationAt: time.Now().AddDate(0, 0, 20)},         // Expira em 20 dias
		{Name: "Churrasco", Category: "Carnes", Quantity: 35, Price: 45.00, ExpirationAt: time.Now().AddDate(0, 2, 0)},        // Expira em 2 meses
		{Name: "Sopa de legumes", Category: "Sopas", Quantity: 10, Price: 9.50, ExpirationAt: time.Now().AddDate(0, 0, 7)},    // Expira em 7 dias
		{Name: "Tacos", Category: "Mexicana", Quantity: 20, Price: 12.99, ExpirationAt: time.Now().AddDate(0, 0, 25)},         // Expira em 25 dias
		{Name: "Frango grelhado", Category: "Carnes", Quantity: 30, Price: 22.50, ExpirationAt: time.Now().AddDate(0, 1, 15)}, // Expira em 1 mês e 15 dias
		{Name: "Ceviche", Category: "Frutos do mar", Quantity: 15, Price: 28.75, ExpirationAt: time.Now().AddDate(0, 0, 12)},  // Expira em 12 dias
	}

	db, err := gorm.Open(postgres.Open("host=db_foods user=postgres password=1234 dbname=foods_db port=5432 sslmode=disable"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

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

	foodRepo := repository.NewFoodRepository(db)
	foodHandler := handlers.NewFoodHandler(foodRepo)

	router := gin.Default()

	router.GET("/foods/all", foodHandler.GetAllFoods)
	router.GET("/foods/:id", foodHandler.GetFoodsByID)
	router.GET("/foods/", foodHandler.GetAllFoodsByCategory)
	router.POST("/foods", foodHandler.CreateFood)
	router.PUT("/foods/:id", foodHandler.UpdateFood)
	router.DELETE("/foods/:id", foodHandler.DeleteFood)

	log.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
