package extras

import "github.com/mailgun/mailgun-go"

type exStr struct {
	mg *mailgun.MailgunImpl
}

var ex *exStr

func InitExtras(mailDomain string, mailAPIKey string) {
	ex = &exStr{
		mg: mailgun.NewMailgun(mailDomain, mailAPIKey),
	}
	return
}
