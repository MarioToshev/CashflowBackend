package src

import (
	"context"

	"git.fhict.nl/I476237/cashflow-backend/orders/models"
)

type OrderService interface {
	CreateOrder(context.Context) (*models.Order, error)
}

type OrderRepo interface {
	Create(context.Context, models.Order) (*models.Order, error)
}