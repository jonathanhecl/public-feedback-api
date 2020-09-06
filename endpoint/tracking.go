package endpoint

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jonathanhecl/public-feedback-api/extras"
)

// HandleTrackingPixel - Handle TrackingPixel
func HandleTrackingPixel(w http.ResponseWriter, r *http.Request) {

	messageID := chi.URLParam(r, "id")
	msg, err := ep.db.GetMessage(messageID)
	if err != nil {
		PixelResponse(w, r)
		return
	}

	code := chi.URLParam(r, "code")
	mds, err := ep.db.GetGroup(msg.ToGroup)
	if err != nil {
		PixelResponse(w, r)
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
		PixelResponse(w, r)
		return
	}

	ip := extras.GetIP(r)
	userAgent := r.UserAgent()

	err = ep.db.SetTracking(messageID, msg.ToGroup, email, ip, userAgent)
	if err != nil {
		PixelResponse(w, r)
		return
	}

	PixelResponse(w, r)

}
