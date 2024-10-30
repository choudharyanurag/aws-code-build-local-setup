package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (appConfig *ServerConfig) routes() http.Handler {
	mux := chi.NewRouter()

	// Read config and add more cors options here
	corsOptions := cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
	}
	mux.Use(cors.Handler(corsOptions))
	mux.Use(middleware.Heartbeat("/ping"))

	return mux
}
