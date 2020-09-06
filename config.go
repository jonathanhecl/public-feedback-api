package main

import "time"

var (
	serverName        = "PublicFeedback Core API"
	serverVer         = "0.0.6"
	logPath           = "development.log"
	minModApproves    = 2
	messageExpiration = time.Hour * 24
	port              = "8080"
	mongoDB           = ""
	googleCert        = ""
	googleGroup       = ""
	secret            = "default"
	mailDomain        = ""
	mailAPIKey        = ""
)
