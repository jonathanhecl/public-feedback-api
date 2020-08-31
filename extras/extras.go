package extras

import "github.com/mailgun/mailgun-go"

type exStr struct {
	mg *mailgun.MailgunImpl
	ps string
}

var ex *exStr

func InitExtras(mailDomain string, mailAPIKey string, Secret string) {
	ex = &exStr{
		mg: mailgun.NewMailgun(mailDomain, mailAPIKey),
		ps: Secret,
	}
	return
}
