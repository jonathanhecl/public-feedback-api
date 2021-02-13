package main

import "time"

var (
	serverName        = "PublicFeedback Core API"
	serverVer         = "0.0.19"
	logPath           = "development.log"
	messageExpiration = time.Hour * 24
	// Default
	apiDomain   = "localhost"
	brandTitle  = "PublicFeedback"
	googleCert  = ""
	googleGroup = ""
	mailAPIKey  = "" // future
	mailDomain  = "" // future
	minApproved = 1
	mongoDB     = ""
	port        = "8080"
	secret      = "default"
	webDomain   = "localhost"
)
