package endpoint

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/jonathanhecl/public-feedback-api/extras"
)

// HandleModerationApproved - Handle ModerationApproved
func HandleModerationApproved(w http.ResponseWriter, r *http.Request) {

	messageID := chi.URLParam(r, "id")
	msg, err := ep.db.GetMessage(messageID)
	if err != nil {
		ErrorResponse(w, r, errors.New("Message not found"))
		return
	}

	if msg.ConfirmedAt.Unix() == 0 {
		ErrorResponse(w, r, errors.New("Message is not confirmed"))
		return
	}
	if msg.ClosedAt.Unix() != 0 {
		ErrorResponse(w, r, errors.New("Message already closed"))
		return
	}

	code := chi.URLParam(r, "code")
	mds, err := ep.db.GetGroup("MOD")
	if err != nil {
		ErrorResponse(w, r, errors.New("No groups settings"))
		return
	}

	email := ""
	for m := range mds.Members {
		if code == extras.GenerateModeratorLink(msg.MessageID, msg.CreatedAt, mds.Members[m].Email) {
			email = mds.Members[m].Email
			break
		}
	}
	if len(email) == 0 {
		ErrorResponse(w, r, errors.New("Moderator not found"))
		return
	}

	ip := extras.GetIP(r)
	userAgent := r.UserAgent()

	err = ep.db.SetModerationVote(messageID, email, true, ip, userAgent)
	if err != nil {
		if strings.Contains(err.Error(), "Moderator already voted") {
			ErrorResponse(w, r, err)
		}
		ErrorResponse(w, r, errors.New("Internal error."))
	}
	SuccessResponse(w, r)

}

// HandleModerationDisapproved - Handle ModerationDisapproved
func HandleModerationDisapproved(w http.ResponseWriter, r *http.Request) {

	MessageID := chi.URLParam(r, "id")
	msg, err := ep.db.GetMessage(MessageID)
	if err != nil {
		ErrorResponse(w, r, errors.New("Message not found"))
		return
	}

	if msg.ConfirmedAt.Unix() == 0 {
		ErrorResponse(w, r, errors.New("Message is not confirmed"))
		return
	}
	if msg.ClosedAt.Unix() != 0 {
		ErrorResponse(w, r, errors.New("Message already closed"))
		return
	}

	code := chi.URLParam(r, "code")

	mds, err := ep.db.GetGroup("MOD")
	if err != nil {
		ErrorResponse(w, r, errors.New("No groups settings"))
		return
	}

	mod := ""
	for m := range mds.Members {
		if code == extras.GenerateModeratorLink(msg.MessageID, msg.CreatedAt, mds.Members[m].Email) {
			mod = mds.Members[m].Email
			break
		}
	}
	if len(mod) == 0 {
		ErrorResponse(w, r, errors.New("Moderator not found"))
		return
	}

	ip := extras.GetIP(r)
	userAgent := r.UserAgent()

	err = ep.db.SetModerationVote(MessageID, mod, false, ip, userAgent)
	if err != nil {
		if strings.Contains(err.Error(), "Moderator already voted") {
			ErrorResponse(w, r, err)
		}
		ErrorResponse(w, r, errors.New("Internal error."))
	}

	SuccessResponse(w, r)

}
