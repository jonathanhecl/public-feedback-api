package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/render"
	"github.com/jonathanhecl/public-feedback-api/database"
	"github.com/jonathanhecl/public-feedback-api/endpoint"

	"github.com/jonathanhecl/public-feedback-api/extras"
)

var started = time.Now()

func main() {
	var err error

	WEBDOMAIN := os.Getenv("WEBDOMAIN")
	if len(WEBDOMAIN) == 0 {
		WEBDOMAIN = webDomain
	}

	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = port
	}
	MONGODB := os.Getenv("MONGODB")
	if len(MONGODB) == 0 {
		MONGODB = mongoDB
	}
	SECRET := os.Getenv("SECRET")
	if len(SECRET) == 0 {
		SECRET = secret
	}
	GOOGLECERT := os.Getenv("GOOGLECERT")
	if len(GOOGLECERT) == 0 {
		GOOGLECERT = googleCert
	}
	GOOGLEGROUP := os.Getenv("GOOGLEGROUP")
	if len(GOOGLEGROUP) == 0 {
		GOOGLEGROUP = googleGroup
	}
	MINAPPROVED, err := strconv.Atoi(os.Getenv("MINAPPROVED"))
	if err != nil {
		MINAPPROVED = minApproved
	}
	MAILDOMAIN := os.Getenv("MAILDOMAIN")
	if len(MAILDOMAIN) == 0 {
		MAILDOMAIN = mailDomain
	}
	MAILAPIKEY := os.Getenv("MAILAPIKEY")
	if len(MAILAPIKEY) == 0 {
		MAILAPIKEY = mailAPIKey
	}

	fmt.Println(serverName + " v" + serverVer)
	fmt.Println("Min. Approved: ", MINAPPROVED)

	time.Sleep(5 * time.Second)

	// MongoDB
	db := database.InitDatabase(MONGODB, GOOGLECERT, GOOGLEGROUP)
	defer database.CloseDatabase(db)
	extras.InitExtras(MAILDOMAIN, MAILAPIKEY, SECRET, WEBDOMAIN)
	endpoint.InitEndpoint(db, MINAPPROVED)

	go func() {
		db.LoadGroups()
		for range time.Tick(30 * time.Minute) {
			db.LoadGroups()
		}
	}()

	// Routes
	r := Routes()

	// Listen and Server
	fmt.Println("Ready... Listen " + PORT + " port...")
	err = http.ListenAndServe(":"+PORT, extras.LogRequest(r)) // HTTP
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HandleGetStatus(w http.ResponseWriter, r *http.Request) {

	type mStatus struct {
		Server  string `json:"server"`
		Version string `json:"version"`
		Online  string `json:"online"`
	}

	render.Status(r, 200)
	render.JSON(w, r, mStatus{
		Server:  serverName,
		Version: serverVer,
		Online:  time.Since(started).String(),
	})
}
