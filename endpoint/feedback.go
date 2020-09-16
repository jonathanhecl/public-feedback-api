package endpoint

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/jonathanhecl/public-feedback-api/endpoint/models"
	"github.com/jonathanhecl/public-feedback-api/extras"
)

// TODO: Responde el mensaje al autor, requiere correo del politico base64
// (guarda IP, User-Agent, ID de correo, correo del politico (base64?))
// El mensaje por defecto es tipo privado. Solo puede enviar una respuesta.

// HandleSendFeedbackMessage - Handle Send Feedback Message
func HandleSendFeedbackMessage(w http.ResponseWriter, r *http.Request) {

	var req models.FeedbackMessageRequest

	// Body parser
	err := DecodeRequest(w, r, &req)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	// Validations
	if len(req.Message) == 0 {
		ErrorResponse(w, r, errors.New("Message required"))
		return
	}

	messageID := chi.URLParam(r, "id")
	msg, err := ep.db.GetMessage(messageID)
	if err != nil {
		ErrorResponse(w, r, errors.New("Original message not found"))
		return
	}

	code := chi.URLParam(r, "code")
	mds, err := ep.db.GetGroup(msg.ToGroup)
	if err != nil {
		ErrorResponse(w, r, errors.New("No groups settings"))
		return
	}

	email := ""
	for m := range mds.Members {
		if code == extras.GenerateMemberLink(msg.MessageID, msg.CreatedAt, mds.Members[m].Email) {
			email = mds.Members[m].Email
			break
		}
	}
	if len(email) == 0 {
		ErrorResponse(w, r, errors.New("Member not found"))
		return
	}

	ip := extras.GetIP(r)
	userAgent := r.UserAgent()

	feedbackID, err := ep.db.NewFeedback(messageID, email, msg.ToGroup, req.Message, ip, userAgent)
	if err != nil {
		if strings.Contains(err.Error(), "Message already replied") {
			ErrorResponse(w, r, err)
		}
		ErrorResponse(w, r, errors.New("Internal error."))
	}

	// TODO: Send feedback email to user
	go EmailFeedbackUser(feedbackID)

	SuccessResponse(w, r)

}
