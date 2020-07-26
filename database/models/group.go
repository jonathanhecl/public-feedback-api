package models

import "time"

type GroupObject struct {
	GroupID   string    `bson:"id" json:"id"`
	Title     string    `bson:"title" json:"title"`
	Members   string    `bson:"members" json:"members"` // JSON struct
	Enabled   bool      `bson:"enabled" json:"enabled"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type MemberGroupObject struct {
	Name    string `bson:"name" json:"name"`
	Email   string `bson:"email" json:"email"`
	Enabled bool   `bson:"enabled" json:"enabled"`
}
