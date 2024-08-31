package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/chat/models/dto"
)

type EnterChat interface {
	EnterChat(ctx context.Context, chatID string) (*dto.GetMatchChatByIDDTORes, error)
}
