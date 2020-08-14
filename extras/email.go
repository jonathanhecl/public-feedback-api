package extras

import (
	"context"
	"fmt"
	"time"
)

func SendEmail(To string, Subject string, Message string) error {

	fmt.Println("Email sended ", Subject, " to ", To, " with message ", Message)

	return nil

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	from := "test@mail.com"
	msg := ex.mg.NewMessage(from, Subject, Message, To)
	/*
		message.SetTemplate("passwordReset")
		message.AddTemplateVariable("passwordResetLink", "some link to your site unique to your user")
	*/
	_, _, err := ex.mg.Send(ctx, msg)
	if err != nil {
		return err
	}
	return nil

}
