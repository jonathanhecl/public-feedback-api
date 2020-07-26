package endpoint

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"../extras"
	"./models"
)

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

	messageID, err := ep.db.NewMessage(req.Email, req.Message, req.GroupID)
	if err != nil {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}

	// send email
	fmt.Println("SEND")
	fmt.Println(messageID)

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

	// resend email
	fmt.Println("RESEND")
	fmt.Println(msg)

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

	log.Println("CONFIRMED")

	SuccessResponse(w, r)

}
