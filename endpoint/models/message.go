package models

type MessageObject struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	GroupID string `json:"group_id"`
}

type RetryMessageRequest struct {
	MessageID string `json:"message_id"`
}

type ConfirmMessageRequest struct {
	MessageID        string `json:"message_id"`
	ConfirmationCode string `json:"confirmation_code"`
}

type GetMessageResponse struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	ToGroup   string `json:"to_group"`
	CreatedAt int64  `json:"created_at"`
	SendedAt  int64  `json:"sended_at"`
}
