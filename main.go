package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"./database"
	"./endpoint"

	"./extras"

	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte(tokenSecret), nil)
}

func main() {
	var err error

	fmt.Println(serverName + " v" + serverVer)

	// MongoDB
	db := database.InitDatabase(mongoUri, mongoDb, googleCert, groupSpreadsheet)
	defer database.CloseDatabase(db)
	extras.InitExtras(mailDomain, mailAPIKey, adminPassword)
	endpoint.InitEndpoint(db)

	go func() {
		db.LoadGroups()
		for range time.Tick(30 * time.Minute) {
			db.LoadGroups()
		}
	}()

	// Routes
	r := Routes()

	// Listen and Server
	fmt.Println("Ready... Listen " + portHTTPS + " port...")
	err = http.ListenAndServeTLS(":"+portHTTPS, "server.crt", "server.key", extras.LogRequest(r)) // HTTPS
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
