package models

import "time"

type TrackingObject struct {
	MessageID string                 `bson:"id" json:"id"`
	GroupID   string                 `bson:"group_id" json:"group_id"`
	Members   []MemberTrackingObject `bson:"members" json:"members"`
	SendedAt  time.Time              `bson:"sended_at" json:"sended_at"`
}

type MemberTrackingObject struct {
	Email         string    `bson:"email" json:"email"`
	IP            string    `bson:"ip" json:"ip"`
	UserAgent     string    `bson:"user_agent" json:"user_agent"`
	Readed        int32     `bson:"readed" json:"readed"`
	FirstReadedAt time.Time `bson:"first_readed_at" json:"first_readed_at"`
	LastReadedAt  time.Time `bson:"last_readed_at" json:"last_readed_at"`
}
