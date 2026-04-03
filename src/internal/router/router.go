package router

import (
	"net/http"

	"github.com/travis8artley/distributed-api-services/src/internal/handlers"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.Health)
	return mux
}
