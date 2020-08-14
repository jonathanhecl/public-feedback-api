package models

type GroupObject struct {
	GroupID string        `json:"group_id"` // Empty = new
	Label   string        `json:"label"`
	Members []MemberGroup `json:"members"`
}

type MemberGroup struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GroupSimpleObject struct {
	GroupID      string `json:"group_id"`
	Label        string `json:"lebel"`
	MembersCount int    `json:"members_count"`
}

type GetGroupsMessageResponse struct {
	Groups []GroupSimpleObject `json:"groups"`
}
