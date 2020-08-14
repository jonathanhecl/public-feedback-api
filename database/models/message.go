package models

import "time"

type MessageObject struct {
	MessageID        string    `bson:"id" json:"id"`
	Name             string    `bson:"name" json:"name"`
	Email            string    `bson:"email" json:"email"`
	IP               string    `bson:"ip" json:"ip"`
	UserAgent        string    `bson:"user_agent" json:"user_agent"`
	ToGroup          string    `bson:"to_group" json:"to_group"`
	Message          string    `bson:"message" json:"message"`
	ConfirmationCode string    `bson:"confirmation_code" json:"confirmation_code"`
	CreatedAt        time.Time `bson:"created_at" json:"created_at"`
	ConfirmedAt      time.Time `bson:"confirmed_at" json:"confirmed_at"` // 0 = not confirmed
	SendedAt         time.Time `bson:"sended_at" json:"sended_at"`       // 0 = not sended
	ClosedAt         time.Time `bson:"closed_at" json:"closed_at"`       // 0 = not clossed
}
