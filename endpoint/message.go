package endpoint

import (
	"errors"
	"fmt"
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
	messageID, err:= ep.db.NewMessage(req.Email, req.Message)
	if err!=nil {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}

	// send email
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

	m, err:=ep.db.GetMessage(req.MessageID)
	if err!=nil {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}

	if m.ConfirmedAt.Unix() > 0 {
		ErrorResponse(w, r, errors.New("Message already confirmed"))
		return
	}

	// resend email
	fmt.Println(m.MessageID)

	SuccessResponse(w, r)

}

// HandleConfirmMessage - Handle Confirm Message
func HandleConfirmMessage(w http.ResponseWriter, r *http.Request) {

	var req models.HandleConfirmMessageRequest

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
	err=ep.db.ConfirmMessage(req.MessageID, req.ConfirmationCode)
	if err!=nil {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}

	SuccessResponse(w, r)

}