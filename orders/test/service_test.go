package test

import (
	"context"
	"log"
	"testing"
	"time"

	"git.fhict.nl/I476237/cashflow-backend/orders/models"
	"git.fhict.nl/I476237/cashflow-backend/orders/src"
	"git.fhict.nl/I476237/cashflow-backend/orders/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)


func TestT_MCreateOrder(t *testing.T) {
	log.Println("starting migration")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := mock.NewMockOrderRepo(ctrl)
	service := src.NewService(mockOrderRepo)

	order := &models.Order{
		ID:              1,
		UserID:          1,
		Ticker:          "sda",
		Quantity:        1,
		PriceOfOneShare: 1,
		TotalPrice:      1,
		DateOfOrder:     time.Time{},
		OrderType:       "d",
		StopLoss:        1,
		TakeProfit:      1,
	}

	ctx := context.Background()
	mockOrderRepo.EXPECT().Create(ctx, *order).Return(order, nil)
	result, err := service.CreateOrder(ctx, *order)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}


func TestCreateOrder(t *testing.T) {
	// Initialize gomock controller
	ctr := gomock.NewController(t)
	defer ctr.Finish()

	mockRepo := mock.NewMockOrderRepo(ctr)
	service := src.NewService(mockRepo)

	order := models.Order{
		ID:              0,
		UserID:          0,
		Ticker:          "AAPL",
		Quantity:        10,
		PriceOfOneShare: 150,
		TotalPrice:      1500,
		DateOfOrder:     time.Now(),
		OrderType:       "buy",
		StopLoss:        0,
		TakeProfit:      0,
	}

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&order, nil)

	ctx := context.Background()
	result, err := service.CreateOrder(ctx, order)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, order.ID, result.ID)
	assert.Equal(t, order.Ticker, result.Ticker)
}
