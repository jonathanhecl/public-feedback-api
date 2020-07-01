package endpoint

import (
	"encoding/json"
	"errors"
	"net/http"

	"./models"
	"github.com/go-chi/render"
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
