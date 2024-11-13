package src

import (
	"context"

	"git.fhict.nl/I476237/cashflow-backend/orders/models"
)

type service struct {
	repo OrderRepo
}

func NewService(repo OrderRepo) *service {
	return &service{repo}
}

func (s *service) CreateOrder(ctx context.Context, order models.Order)(*models.Order, error)  {
	return s.repo.Create(ctx, order)
}