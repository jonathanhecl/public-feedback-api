package endpoint

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/jonathanhecl/public-feedback-api/endpoint/models"
)

func DecodeRequest(w http.ResponseWriter, r *http.Request, req interface{}) error {

	if r.Body == nil {
		return errors.New("Without body")
	}
	if json.NewDecoder(r.Body).Decode(&req) != nil {
		return errors.New("Decode error")
	}
	return nil

}

func ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

	render.Status(r, 500)
	render.JSON(w, r, models.ErrorObject{Message: err.Error()})

}

func SuccessResponse(w http.ResponseWriter, r *http.Request) {

	render.Status(r, 200)
	render.JSON(w, r, models.SuccessObject{Success: true})

}

func SuccessResponseInterface(w http.ResponseWriter, r *http.Request, s interface{}) {

	render.Status(r, 200)
	render.JSON(w, r, s)

}

func PixelResponse(w http.ResponseWriter, r *http.Request) {

	render.Status(r, 200)
	w.Header().Set("Content-Type", "image/gif")
	output, _ := base64.StdEncoding.DecodeString("R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs=")
	//io.WriteString(w, string(output))
	render.Respond(w, r, output)

}
