package services

import ports "github.com/DBrange/didis-comp-bk/domains/chat/ports/drivens"

type ChatService struct {
	chatQuerier ports.ForQueryingChat
}

func NewChatService(chatQuerier ports.ForQueryingChat) *ChatService {
	return &ChatService{
		chatQuerier: chatQuerier,
	}
}

