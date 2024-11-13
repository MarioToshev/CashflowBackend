package main

import "net/http"

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("Post /order/{orderId}", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	// createOrder
}