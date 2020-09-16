package main

import "time"

var (
	serverName        = "PublicFeedback Core API"
	serverVer         = "0.0.8"
	logPath           = "development.log"
	messageExpiration = time.Hour * 24
	// Default
	port        = "8080"
	mongoDB     = ""
	secret      = "default"
	googleCert  = ""
	googleGroup = ""
	minApproved = 1
	mailDomain  = "" // future
	mailAPIKey  = "" // future
)
