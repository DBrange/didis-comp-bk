package port

import "github.com/DBrange/didis-comp-bk/domains/chat/ports/drivers/interfaces"

type ForChat interface {
	interfaces.CreateMatchChat
	interfaces.EnterChat
	interfaces.SendMessage
}
