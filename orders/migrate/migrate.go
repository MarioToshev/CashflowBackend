package migrate

import (
	"log"

	"git.fhict.nl/I476237/cashflow-backend/orders/initializers"
	"git.fhict.nl/I476237/cashflow-backend/orders/models"
)

func init() {
	initializers.ConnectToDB()
}

func MigrateTables() {	
	log.Println("starting migration")
	log.Println(initializers.DB)	
	initializers.DB.AutoMigrate(&models.Order{})
}