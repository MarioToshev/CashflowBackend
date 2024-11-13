package main

import (
	"log"
	"net/http"
)

const (
	httpPort = ":8080"
)

func main() {
	mux := http.NewServeMux()	
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("server running on port %s", httpPort)

	if err := http.ListenAndServe(httpPort, mux); err != nil {
		panic(err)
	}

}