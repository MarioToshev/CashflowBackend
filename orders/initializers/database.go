package initializers

import (
	"log"

	"git.fhict.nl/I476237/cashflow-backend/orders/models"
	_ "github.com/lib/pq"
	"github.com/subinoybiswas/goenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	if DB != nil {
		return DB
	}
	ConnectToDB()
	DB.AutoMigrate(&models.Order{})
	return DB
}

func ConnectToDB() {
	var err error
	dsn := GetConnString()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
	}
	if DB != nil {
		log.Println("Connected to database")

	}
}

func GetConnString() string {
	value, err := goenv.GetEnvFrmFile("DB_CONN_STR")
	if err != nil {
		log.Fatal("Failed to get connection string")
	}
	return value
}
