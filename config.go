package main

import "time"

var (
	serverName        = "PublicFeedback Core API"
	serverVer         = "0.0.5"
	logPath           = "development.log"
	minModApproves    = 2
	messageExpiration = time.Hour * 24
	port              = "443"
	mongoDB           = ""
	googleCert        = ""
	googleGroup       = ""
	secret            = "default"
	mailDomain        = ""
	mailAPIKey        = ""
)
