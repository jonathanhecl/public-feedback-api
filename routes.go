package main

import (
	"net/http"
	"time"

	"./endpoint"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Routes() chi.Router {

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Use(middleware.Timeout(60 * time.Second))

	// Check if everything is good.
	r.Use(middleware.Heartbeat("/ping"))

	r.Group(func(r chi.Router) {

		r.Route("/message", func(r chi.Router) {
			r.Post("/", http.HandlerFunc(endpoint.HandleNewMessage))
			r.Post("/confirm", http.HandlerFunc(endpoint.HandleConfirmMessage))
			r.Get("/resend", http.HandlerFunc(endpoint.HandleRetryConfirmationMessage))
		})

	})

	return r

}
