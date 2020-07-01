package models

type ErrorObject struct {
	Message string `json:"message"`
}

type SuccessObject struct {
	Success bool `json:"success"`
}
