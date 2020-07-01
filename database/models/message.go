package models

import "time"

type MessageObject struct {
	MessageID        string    `bson:"id" json:"id"`
	Email            string    `bson:"email" json:"email"`
	Message          string    `bson:"message" json:"message"`
	ConfirmationCode string    `bson:"confirmation_code" json:"confirmation_code"`
	CreatedAt        time.Time `bson:"created_at" json:"created_at"`
	ConfirmedAt      time.Time `bson:"confirmed_at" json:"confirmed_at"`
}
