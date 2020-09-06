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

		// PAGINA DE ENVIO DE MENSAJE del usuario

		r.Get("/groups", http.HandlerFunc(endpoint.HandleGetGroupsMessage)) // Get groups use ISO 2 chars for country

		r.Route("/message", func(r chi.Router) {
			r.Post("/", http.HandlerFunc(endpoint.HandleNewMessage))                    // Post a new message
			r.Post("/confirm", http.HandlerFunc(endpoint.HandleConfirmMessage))         // Confirm a message (with the email)
			r.Get("/resend", http.HandlerFunc(endpoint.HandleRetryConfirmationMessage)) // Resend email to confirm a message
		})

		// PAGINA de respuesta de aprobado/desaprobado

		//api.populeaks.com/348983748734/moderation/approved/38473847-54545-6565656/ APROBAR
		//api.populeaks.com/348983748734/moderation/disapproved/38473847-54545-6565656/ DESAPROBAR/RECHAZAR

		r.Route("/moderation", func(r chi.Router) {
			r.Get("/{id}/approved/{code}", http.HandlerFunc(endpoint.HandleModerationApproved))
			r.Get("/{id}/disapproved/{code}", http.HandlerFunc(endpoint.HandleModerationDisapproved))
		})

		//<img src="https://www.populeaks.com/tracking/38473847-54545-6565656/abc8347837483/pixel.gif" /> 1x1
		r.Route("/tracking", func(r chi.Router) {
			r.Get("/{id}/{code}/pixel.gif", http.HandlerFunc(endpoint.HandleTrackingPixel))
		})

		// PAGINA de RESPUESTA del politico

		// https://www.populeaks.com/feedback/38473847-54545-6565656/?cp=abc8347837483 RESPONDER
		r.Route("/feedback", func(r chi.Router) { // TODO:
			r.Get("/{id}/{code}/", http.HandlerFunc(endpoint.HandleGetGroupsMessage))  // TODO: Lee el mensaje del usuario, requiere correo del politico base64 (guarda IP, User-Agent, ID de correo, correo del politico (base64?))
			r.Post("/{id}/{code}/", http.HandlerFunc(endpoint.HandleGetGroupsMessage)) // TODO: Responde el mensaje al autor, requiere correo del politico base64 (guarda IP, User-Agent, ID de correo, correo del politico (base64?)) El mensaje por defecto es tipo privado. Solo puede enviar una respuesta.
		})

		/*
			https://stackoverflow.com/questions/39543845/golang-iris-web-serve-1x1-pixel

						iris.Get("/", func(ctx *iris.Context) {
				        img := image.NewRGBA(image.Rect(0, 0, 1, 1)) //We create a new image of size 1x1 pixels
				        img.Set(0, 0, color.RGBA{255, 0, 0, 255}) //set the first and only pixel to some color, in this case red

				        err := gif.Encode(ctx.Response.BodyWriter(), img, nil) //encode the rgba image to gif, using gif.encode and write it to the response
				        if err != nil {
				            panic(err) //if we encounter some problems panic
				        }
				        ctx.SetContentType("image/gif") //set the content type so the browser can identify that our response is actually an gif image
				    })
					iris.Listen(":8080")
		*/

	})

	return r

}
