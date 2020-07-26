package models

import "time"

type GroupObject struct {
	GroupID   string        `json:"group_id"` // Empty = new
	Title     string        `json:"title"`
	Members   []MemberGroup `json:"members"`
	Enabled   bool          `json:"enabled"`
	CreatedAt time.Time     `json:"created_at"` // Read only
	UpdatedAt time.Time     `json:"updated_at"` // Read only
}

type GroupSimpleObject struct {
	GroupID      string `json:"group_id"`
	Title        string `json:"title"`
	MembersCount int    `json:"members_count"`
}

type MemberGroup struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Enabled bool   `json:"enabled"`
}

type GetGroupsMessageResponse struct {
	Groups []GroupSimpleObject `json:"groups"`
}

type AdminSetGroupMessageRequest struct {
	Group GroupObject `json:"group"`
}

type AdminSetGroupMessageResponse struct {
	Group GroupObject `json:"group"`
}

type AdminDeleteGroupMessageRequest struct {
	GroupID string `json:"group_id"`
}

type AdminGetGroupsMessageResponse struct {
	Groups []GroupObject `json:"groups"`
}
