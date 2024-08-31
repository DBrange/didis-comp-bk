package models

type SOCKET_EVENT string

const (
	SOCKET_EVENT_SEND_MESSAGE SOCKET_EVENT = "SEND_MESSAGE"
	SOCKET_EVENT_VOTE         SOCKET_EVENT = "VOTE"
)

func (g SOCKET_EVENT) IsValid() bool {
	switch g {
	case SOCKET_EVENT_SEND_MESSAGE, SOCKET_EVENT_VOTE:
		return true
	}
	return false
}
