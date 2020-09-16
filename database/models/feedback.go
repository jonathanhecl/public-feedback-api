package models

import "time"

type FeedbackObject struct {
	FeedbackID string    `bson:"id" json:"id"`
	MessageID  string    `bson:"message_id" json:"message_id"`
	Email      string    `bson:"email" json:"email"`
	IP         string    `bson:"ip" json:"ip"`
	UserAgent  string    `bson:"user_agent" json:"user_agent"`
	ToGroup    string    `bson:"to_group" json:"to_group"`
	Message    string    `bson:"message" json:"message"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
}
