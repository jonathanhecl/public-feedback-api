package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jonathanhecl/public-feedback-api/database"
	"github.com/jonathanhecl/public-feedback-api/endpoint"

	"github.com/jonathanhecl/public-feedback-api/extras"
)

func main() {
	var err error

	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = port
	}
	MONGODB := os.Getenv("MONGODB")
	if len(MONGODB) == 0 {
		MONGODB = mongoDB
	}
	GOOGLECERT := os.Getenv("GOOGLECERT")
	if len(GOOGLECERT) == 0 {
		GOOGLECERT = googleCert
	}
	GOOGLEGROUP := os.Getenv("GOOGLEGROUP")
	if len(GOOGLEGROUP) == 0 {
		GOOGLEGROUP = googleGroup
	}
	SECRET := os.Getenv("SECRET")
	if len(SECRET) == 0 {
		SECRET = secret
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

	// MongoDB
	db := database.InitDatabase(MONGODB, GOOGLECERT, GOOGLEGROUP)
	defer database.CloseDatabase(db)
	extras.InitExtras(MAILDOMAIN, MAILAPIKEY, SECRET)
	endpoint.InitEndpoint(db, minModApproves)

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
