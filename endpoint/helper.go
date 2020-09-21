package endpoint

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
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

func ParseTemplate(filename string, data map[string]string) string {

	t, err := template.ParseFiles("./templates/" + filename + ".html")
	if err != nil {
		fmt.Println("ParseTemplate->ParseFiles: ", err)
		return ""
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println("ParseTemplate->Execute: ", err)
		return ""
	}

	return buf.String()

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
	io.WriteString(w, string(output))

}
