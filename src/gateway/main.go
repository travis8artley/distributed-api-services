package main

import (
	"log"
	"net/http"

	"github.com/travis8artley/distributed-api-services/src/internal/router"
)

func main() {
	mux := router.NewMux()
	log.Println("gateway online on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
