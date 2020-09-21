package extras

type exStr struct {
	mailAPIKey string
	mailDomain string
	ps         string
}

var ex *exStr

func InitExtras(mailDomain string, mailAPIKey string, Secret string) {
	ex = &exStr{
		//mg: mailgun.NewMailgun(mailDomain, mailAPIKey),
		mailAPIKey: mailAPIKey,
		mailDomain: mailDomain,
		ps:         Secret,
	}
	return
}
