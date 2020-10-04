package extras

type exStr struct {
	mailAPIKey string
	mailDomain string
	ps         string
	webDomain  string
	apiDomain  string
}

var ex *exStr

func InitExtras(mailDomain string, mailAPIKey string, secret string, webDomain string, apiDomain string) {
	ex = &exStr{
		//mg: mailgun.NewMailgun(mailDomain, mailAPIKey),
		mailAPIKey: mailAPIKey,
		mailDomain: mailDomain,
		ps:         secret,
		webDomain:  webDomain,
		apiDomain:  apiDomain,
	}
	return
}
