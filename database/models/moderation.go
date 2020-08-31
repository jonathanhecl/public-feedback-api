package models

import "time"

type ModerationObject struct {
	MessageID string       `bson:"id" json:"id"`
	Votes     []VoteObject `bson:"votes" json:"votes"`
}

type VoteObject struct {
	Email     string    `bson:"email" json:"email"`
	IsApprove bool      `bson:"is_approve" json:"is_approve"`
	IP        string    `bson:"ip" json:"ip"`
	UserAgent string    `bson:"user_agent" json:"user_agent"`
	VotedAt   time.Time `bson:"voted_at" json:"voted_at"`
}
