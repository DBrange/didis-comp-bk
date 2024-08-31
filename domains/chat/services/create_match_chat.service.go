package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ChatService) CreateMatchChat(ctx context.Context, matchID string, competitorIDs []string, userID string) error {
	err := s.chatQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		chatID, err := s.chatQuerier.CreateChat(ctx, matchID, models.CHAT_GROUP)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating chat")
		}

		if err := s.chatQuerier.CreateParticipantChats(ctx, chatID, competitorIDs, models.CHAT_GROUP); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating participantChat")
		}

		if err := s.chatQuerier.CreateParticipantChats(ctx, chatID, []string{userID}, models.CHAT_INDIVIDUAL); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating participantChat")
		}

		return nil
	})
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament location")
	}

	return nil
}
