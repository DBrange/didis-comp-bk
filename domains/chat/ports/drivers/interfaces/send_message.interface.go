package interfaces

import "context"

type SendMessage interface {
	SendMessage(ctx context.Context, chatID, senderID, content string) (string, error)
}