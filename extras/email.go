package extras

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

// Tutorial: https://devanswers.co/create-application-specific-password-gmail/
// SMTP Password: https://security.google.com/settings/security/apppasswords
// Enable Access: https://accounts.google.com/DisplayUnlockCaptcha

func SendEmail(To string, Subject string, Message string) error {

	fmt.Println("Email ", Subject, " to ", To, " with message ", Message)

	m := gomail.NewMessage()
	m.SetHeader("From", ex.mailDomain)
	m.SetHeader("To", To)
	m.SetHeader("Subject", Subject)
	m.SetBody("text/html", Message)

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, ex.mailDomain, ex.mailAPIKey)
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("SendEmail->SMTP error: %s", err)
	}

	fmt.Println("Sended OK")

	return nil
	/*
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
