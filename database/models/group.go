package models

import "time"

type GroupObject struct {
	GroupID   string        `json:"group_id" bson:"group_id"`
	Label     string        `json:"label" bson:"label"`
	Members   []MemberGroup `json:"members" bson:"members"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

type MemberGroup struct {
	Name  string `json:"name" bson:"name" `
	Email string `json:"email" bson:"email" `
}
