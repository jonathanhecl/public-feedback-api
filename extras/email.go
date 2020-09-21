package extras

import (
	"fmt"
	"net/smtp"
)

func SendEmail(To string, Subject string, Message string) error {

	fmt.Println("Email sended ", Subject, " to ", To, " with message ", Message)

	type SmtpTemplateData struct {
		From    string
		To      string
		Subject string
		Body    string
	}
	email := `From: ` + ex.mailDomain + `
	To: ` + To + `
	Subject: ` + Subject + `

	` + Message

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", ex.mailDomain, ex.mailAPIKey, "smtp.gmail.com"),
		ex.mailDomain, []string{To}, []byte(email))

	if err != nil {
		fmt.Printf("smtp error: %s", err)
	}

	return nil
	/*
		//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		//defer cancel()
		from := "test@mail.com"
		msg := ex.mg.NewMessage(from, Subject, Message, To)

			message.SetTemplate("passwordReset")
			message.AddTemplateVariable("passwordResetLink", "some link to your site unique to your user")

		_, _, err := ex.mg.Send(msg)
		if err != nil {
			return err
		}
		return nil
	*/

}
