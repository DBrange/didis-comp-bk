package dto

type CreateChatMessageDTOReq struct {
	ChatID   string `json:"chat_id" validate:"required"`
	SenderID string `json:"sender_id" validate:"required"` // userID
	Content  string `json:"content" validate:"required"`
}
