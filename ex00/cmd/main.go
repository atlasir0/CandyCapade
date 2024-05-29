package main

import (
	"ex00/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/buy_candy", handlers.BuyCandy)

	log.Println("Server starting on port 3333")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatalf("could not start server: %s", err.Error())
	}
}
