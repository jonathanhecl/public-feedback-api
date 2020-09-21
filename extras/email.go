package extras

import (
	"fmt"
	"net/smtp"
)

// Tutorial: https://devanswers.co/create-application-specific-password-gmail/
// SMTP Password: https://security.google.com/settings/security/apppasswords
// Enable Access: https://accounts.google.com/DisplayUnlockCaptcha

func SendEmail(To string, Subject string, Message string) error {

	fmt.Println("Email sended ", Subject, " to ", To, " with message ", Message)

	email := "From: " + ex.mailDomain
	email = "\nTo: " + To
	email = "\nSubject: " + Subject
	email = "\n\n" + Message

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", ex.mailDomain, ex.mailAPIKey, "smtp.gmail.com"),
		ex.mailDomain, []string{To}, []byte(email))

	if err != nil {
		fmt.Printf("smtp error: %s", err)
	} else {
		fmt.Println("Sended OK")
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
