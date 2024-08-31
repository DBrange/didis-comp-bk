package chats

import (
	chat_ports "github.com/DBrange/didis-comp-bk/domains/chat/ports/drivers"
		socketio "github.com/googollee/go-socket.io"

)

type Handler struct {
	chat         chat_ports.ForChat
	socketServer *socketio.Server
}

func NewHandlerChat(chat chat_ports.ForChat,socketServer *socketio.Server) *Handler {
	return &Handler{
		chat: chat,
		socketServer:socketServer,
	}
}
