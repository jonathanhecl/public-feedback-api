package main

import "time"

var (
	serverName        = "PublicFeedback Core API"
	serverVer         = "0.0.11"
	logPath           = "development.log"
	messageExpiration = time.Hour * 24
	// Default
	webDomain   = "localhost"
	port        = "8080"
	mongoDB     = ""
	secret      = "default"
	googleCert  = ""
	googleGroup = ""
	minApproved = 1
	mailDomain  = "" // future
	mailAPIKey  = "" // future
)
