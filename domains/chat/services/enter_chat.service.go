package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/chat/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ChatService) EnterChat(ctx context.Context, chatID string) (*dto.GetMatchChatByIDDTORes, error) {
	chatDTO, err := s.chatQuerier.GetMatchChatByID(ctx, chatID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting chat")
	}

	return chatDTO, nil
}

