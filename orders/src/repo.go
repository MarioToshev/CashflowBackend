package src

import (
	"context"
	"log"

	"git.fhict.nl/I476237/cashflow-backend/orders/initializers"
	"git.fhict.nl/I476237/cashflow-backend/orders/models"
)


type repo struct {
}

func NewRepo() *repo {
	return &repo{}
}

func (o *repo) Create(ctx context.Context, order models.Order) (*models.Order, error) {
	result := initializers.GetDB().Create(&order)
	if result.Error != nil {
		log.Println("Error creating order")
		return nil, result.Error
	}
	return &order, nil
}

func MigrateTableOrders(ctx context.Context) {
	log.Println("starting migration")
	if !initializers.GetDB().Migrator().HasTable(&models.Order{}) {
		log.Println("Migrating table orders..")
		initializers.DB.AutoMigrate(&models.Order{})
	}
}



func GetOrder(id int) (*models.Order, error) {
	var order models.Order
	err := initializers.GetDB().First(&order, id).Error
	return &order, err
}

func UpdateOrder(order *models.Order) error {
	return initializers.GetDB().Save(order).Error
}

func DeleteOrder(id int) error {
	return initializers.GetDB().Delete(&models.Order{}, id).Error
}