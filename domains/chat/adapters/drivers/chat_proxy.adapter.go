package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/chat/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/chat/services"
)

type ChatProxyAdapter struct {
	chatService *services.ChatService
}

func NewChatProxyAdapter(chatService *services.ChatService) *ChatProxyAdapter {
	return &ChatProxyAdapter{
		chatService: chatService,
	}
}

func (a *ChatProxyAdapter) CreateMatchChat(ctx context.Context, matchID string, competitorIDs []string, userID string) error {
	return a.chatService.CreateMatchChat(ctx, matchID, competitorIDs, userID)
}

func (a *ChatProxyAdapter) EnterChat(ctx context.Context, chatID string) (*dto.GetMatchChatByIDDTORes, error) {
	return a.chatService.EnterChat(ctx, chatID)
}

func (a *ChatProxyAdapter) SendMessage(ctx context.Context, chatID, senderID, content string) (string, error) {
	return a.chatService.SendMessage(ctx, chatID, senderID, content)
}
