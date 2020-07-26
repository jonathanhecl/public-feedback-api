package models

import "time"

type MessageObject struct {
	MessageID        string    `bson:"id" json:"id"`
	Email            string    `bson:"email" json:"email"`
	ToGroup          string    `bson:"to_group" json:"to_group"`
	Message          string    `bson:"message" json:"message"`
	CreatedAt        time.Time `bson:"created_at" json:"created_at"`
	ConfirmationCode string    `bson:"confirmation_code" json:"confirmation_code"`
	ConfirmedAt      time.Time `bson:"confirmed_at" json:"confirmed_at"`
	AdmittedBy       string    `bson:"admitted_by" json:"admitted_by"`
	AdmittedAt       time.Time `bson:"admitted_at" json:"admitted_at"`
	SendedAt         time.Time `bson:"sended_at" json:"sended_at"`
}
