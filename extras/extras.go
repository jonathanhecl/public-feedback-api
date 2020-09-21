package extras

type exStr struct {
	mailAPIKey string
	mailDomain string
	ps         string
	webDomain  string
}

var ex *exStr

func InitExtras(mailDomain string, mailAPIKey string, secret string, webDomain string) {
	ex = &exStr{
		//mg: mailgun.NewMailgun(mailDomain, mailAPIKey),
		mailAPIKey: mailAPIKey,
		mailDomain: mailDomain,
		ps:         secret,
		webDomain:  webDomain,
	}
	return
}
