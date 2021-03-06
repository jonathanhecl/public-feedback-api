package endpoint

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jonathanhecl/public-feedback-api/endpoint/models"
	"github.com/jonathanhecl/public-feedback-api/extras"
)

// HandleGetMessage - Handle Get Message
func HandleGetMessage(w http.ResponseWriter, r *http.Request) {

	messageID := chi.URLParam(r, "id")
	msg, err := ep.db.GetMessage(messageID)
	if err != nil {
		ErrorResponse(w, r, errors.New("Message not found"))
		return
	}

	var res models.GetMessageResponse

	fmt.Println(msg)

	res.Name = msg.Name
	res.Message = msg.Message
	res.ToGroup = msg.ToGroup
	res.CreatedAt = msg.CreatedAt.Unix()
	res.SendedAt = msg.SendedAt.Unix()

	// TODO: Prevent message not sended (SendedAt==0)

	SuccessResponseInterface(w, r, res)

}

// HandleNewMessage - Handle New Message
func HandleNewMessage(w http.ResponseWriter, r *http.Request) {

	var req models.MessageObject

	// Body parser
	err := DecodeRequest(w, r, &req)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	// Validations
	if len(req.Name) == 0 {
		ErrorResponse(w, r, errors.New("Name required"))
		return
	}
	if len(req.Email) == 0 || !extras.ValidateEmail(req.Email) {
		ErrorResponse(w, r, errors.New("E-mail invalid"))
		return
	}
	if len(req.Message) == 0 {
		ErrorResponse(w, r, errors.New("Message required"))
		return
	}
	if len(req.GroupID) == 0 {
		ErrorResponse(w, r, errors.New("Group required"))
		return
	}
	if _, err := ep.db.GetGroup(req.GroupID); err != nil {
		ErrorResponse(w, r, errors.New("Group not found"))
		return
	}

	ip := extras.GetIP(r)
	userAgent := r.UserAgent()

	messageID, err := ep.db.NewMessage(req.Email, req.Name, req.Message, req.GroupID, ip, userAgent)
	if err != nil {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}

	// TODO: Send email to user
	fmt.Println("New message user confirmation pending: ", messageID)
	go EmailUserConfirmation(messageID)

	SuccessResponse(w, r)
}

// HandleRetryConfirmationMessage - Handle Retry Confirmation Message
func HandleRetryConfirmationMessage(w http.ResponseWriter, r *http.Request) {

	var req models.RetryMessageRequest

	// Body parser
	err := DecodeRequest(w, r, &req)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	// Validations
	if len(req.MessageID) == 0 {
		ErrorResponse(w, r, errors.New("MessageID required"))
		return
	}

	msg, err := ep.db.GetMessage(req.MessageID)
	if err != nil {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}

	if msg.ConfirmedAt.Unix() > 0 {
		ErrorResponse(w, r, errors.New("Message already confirmed"))
		return
	}

	// TODO: Resend email to user
	fmt.Println("Resend message user confirmation pending: ", msg.MessageID)
	fmt.Println("Confirmation code:", msg.ConfirmationCode)
	go EmailUserConfirmation(msg.MessageID)

	SuccessResponse(w, r)

}

// HandleConfirmMessage - Handle Confirm Message
func HandleConfirmMessage(w http.ResponseWriter, r *http.Request) {

	var req models.ConfirmMessageRequest

	// Body parser
	err := DecodeRequest(w, r, &req)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	// Validations
	if len(req.MessageID) == 0 {
		ErrorResponse(w, r, errors.New("MessageID required"))
		return
	}
	if len(req.ConfirmationCode) == 0 {
		ErrorResponse(w, r, errors.New("ConfirmationCode required"))
		return
	}

	// confirm
	err = ep.db.ConfirmMessage(req.MessageID, req.ConfirmationCode)
	if err != nil {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}

	// TODO: Confirmed, send to mods
	log.Println("Message user confirmation, send it to mods: ", req.MessageID)
	go EmailModerationWait(req.MessageID)

	SuccessResponse(w, r)

}
