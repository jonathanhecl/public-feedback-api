package models

type MessageObject struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

type RetryMessageRequest struct {
	MessageID string `json:"message_id"`
}

type HandleConfirmMessageRequest struct {
	MessageID        string `json:"message_id"`
	ConfirmationCode string `json:"confirmation_code"`
}
