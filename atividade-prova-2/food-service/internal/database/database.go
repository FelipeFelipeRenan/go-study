package database

import (
	"foods/internal/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabase() (*gorm.DB, error){
	foods := []models.Food{
		{Name: "Lasanha", Category: "Massas", Quantity: 20, Price: 25.99, ExpirationAt: time.Now().AddDate(0, 1, 0)},
		{Name: "Sushi", Category: "Japonesa", Quantity: 30, Price: 30.50, ExpirationAt: time.Now().AddDate(0, 0, 15)},
		// Adicione outros alimentos aqui
	}

	db, err := gorm.Open(postgres.Open("host=db_foods user=postgres password=1234 dbname=foods_db port=5432 sslmode=disable"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados", err)
		return nil, err
	}

	// Drop and create the table
	db.Migrator().DropTable(&models.Food{})
	if err := db.AutoMigrate(&models.Food{}); err != nil {
		log.Fatal("Erro ao migrar modelo", err)
		return nil, err
	}

	log.Println("Tabela 'food' criada com sucesso")

	for _, f := range foods {
		if err := db.Create(&f).Error; err != nil {
			log.Fatal("Erro ao inserir dados de seed para alimentos:", err)
		}
	}
	log.Println("Seed de dados para participantes concluído com sucesso!")

	return db, nil
}