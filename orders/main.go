package main

import (
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("Order service is initialization started")
}

func main() {
	port := ":8081"
	fmt.Printf("Listening on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Failed to listen on port %s: %v\n", port, err)
	}
	// repo := src.NewRepo()
	// service := src.NewService(repo)
	// fmt.Println("Order service is running")
	// fmt.Println("Creating order...")
	// result, err := service.CreateOrder(context.Background(), models.Order{
	// 	ID:              0,
	// 	UserID:          0,
	// 	Ticker:          "",
	// 	Quantity:        0,
	// 	PriceOfOneShare: 0,
	// 	TotalPrice:      0,
	// 	DateOfOrder:     time.Time{},
	// 	OrderType:       "",
	// 	StopLoss:        0,
	// 	TakeProfit:      0,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)
}