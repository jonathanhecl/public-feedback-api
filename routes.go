package main

import (
	"net/http"
	"time"

	"github.com/jonathanhecl/public-feedback-api/endpoint"

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

		r.Get("/status", http.HandlerFunc(HandleGetStatus))

		r.Get("/groups", http.HandlerFunc(endpoint.HandleGetGroupsMessage)) // Get groups use ISO 2 chars for country

		r.Route("/message", func(r chi.Router) {
			r.Post("/", http.HandlerFunc(endpoint.HandleNewMessage))                    // Post a new message
			r.Post("/confirm", http.HandlerFunc(endpoint.HandleConfirmMessage))         // Confirm a message (with the email)
			r.Get("/resend", http.HandlerFunc(endpoint.HandleRetryConfirmationMessage)) // Resend email to confirm a message
			r.Get("/{id}", http.HandlerFunc(endpoint.HandleGetMessage))                 // Get a message
		})

		//api.populeaks.com/348983748734/moderation/approved/38473847-54545-6565656/ APPROVED
		//api.populeaks.com/348983748734/moderation/disapproved/38473847-54545-6565656/ DISAPPROVED

		r.Route("/moderation", func(r chi.Router) {
			r.Get("/{id}/approved/{code}", http.HandlerFunc(endpoint.HandleModerationApproved))
			r.Get("/{id}/disapproved/{code}", http.HandlerFunc(endpoint.HandleModerationDisapproved))
		})

		//<img src="https://www.populeaks.com/tracking/38473847-54545-6565656/abc8347837483/pixel.gif" /> 1x1
		r.Route("/tracking", func(r chi.Router) {
			r.Get("/{id}/{code}/pixel.gif", http.HandlerFunc(endpoint.HandleTrackingPixel))
		})

		// https://www.populeaks.com/feedback/38473847-54545-6565656/?cp=abc8347837483 RESPONDER
		r.Route("/feedback", func(r chi.Router) {
			r.Post("/{id}/{code}", http.HandlerFunc(endpoint.HandleSendFeedbackMessage))
		})

	})

	return r

}
