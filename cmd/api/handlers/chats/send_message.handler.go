package chats

import (
	"context"
	"fmt"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/chat/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"
	socketio "github.com/googollee/go-socket.io"
)

func (h *Handler) SendMessage(s socketio.Conn, message *dto.CreateChatMessageDTOReq) {
	ctx := context.Background()

	msg, err := sendMessageBodyData(message)
	if err != nil {
		s.Emit("error", err)
	}
	
	// Llamar al servicio para enviar el mensaje
	messageID, err := h.chat.SendMessage(ctx, msg.ChatID, msg.SenderID, msg.Content)
	if err != nil {
		s.Emit("error", err.Error())
		return
	}

	// Emitir el mensaje a todos los participantes del chat
	h.socketServer.BroadcastToRoom("/", msg.ChatID, string(models.SOCKET_EVENT_SEND_MESSAGE), map[string]interface{}{
		"id":        messageID,
		"sender_id": msg.SenderID,
		"content":   msg.Content,
		"timestamp": time.Now().UTC(),
	})
}

func sendMessageBodyData(msg *dto.CreateChatMessageDTOReq) (*dto.CreateChatMessageDTOReq, error) {
	// Validar la estructura excepto el campo Location
	err := utils.Validate.Struct(msg)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error validation body"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	return msg, nil
}
