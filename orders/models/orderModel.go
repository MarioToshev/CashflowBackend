package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID              int    `gorm:"primaryKey"`
	UserID          int    `gorm:"index"`
	Ticker          string `gorm:"size:10"`
	Quantity        int
	PriceOfOneShare float64
	TotalPrice      float64
	DateOfOrder     time.Time `gorm:"autoCreateTime"`
	OrderType       string    `gorm:"size:20"`
	StopLoss        float64
	TakeProfit      float64
}