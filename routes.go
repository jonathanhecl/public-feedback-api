package main

import (
	"net/http"
	"time"

	"./endpoint"
	"./extras"

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
			r.Post("/", http.HandlerFunc(endpoint.HandleNewMessage))                    // Post a new message
			r.Post("/confirm", http.HandlerFunc(endpoint.HandleConfirmMessage))         // Confirm a message (with the email)
			r.Get("/resend", http.HandlerFunc(endpoint.HandleRetryConfirmationMessage)) // Resend email to confirm a message
		})

		r.Get("/groups", http.HandlerFunc(endpoint.HandleGetGroupsMessage)) // Get groups

		r.Route("/admin", func(r chi.Router) {
			r.Use(extras.Authenticator)
			r.Get("/groups", http.HandlerFunc(endpoint.HandleAdminGetGroupsMessage))      // Get groups detail
			r.Post("/groups", http.HandlerFunc(endpoint.HandleAdminSetGroupMessage))      // Create or edit a group
			r.Delete("/groups", http.HandlerFunc(endpoint.HandleAdminDeleteGroupMessage)) // Delete a group
		})

	})

	return r

}
