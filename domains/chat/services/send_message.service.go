package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/chat/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ChatService) SendMessage(ctx context.Context, chatID, senderID, content string) (string, error) {
	if err := s.chatQuerier.VerifyChatExists(ctx, chatID); err != nil {
		return "", customerrors.HandleErrMsg(err, "tournament", "error when verifying if chat exists")
	}

	if err := s.chatQuerier.VerifyUserExists(ctx, senderID); err != nil {
		return "", customerrors.HandleErrMsg(err, "tournament", "error when verifying if user exists")
	}

	messageDTO := &dto.CreateChatMessageDTOReq{
		ChatID:   chatID,
		SenderID: senderID,
		Content:  content,
	}

	messageID, err := s.chatQuerier.CreateChatMessage(ctx, messageDTO)
	if err != nil {
		return "", customerrors.HandleErrMsg(err, "tournament", "error when creating chatMessage")
	}

	return messageID, nil
}
