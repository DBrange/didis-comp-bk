package models

type CHAT string

const (
	CHAT_GROUP      CHAT = "GROUP"
	CHAT_INDIVIDUAL CHAT = "INDIVIDUAL"
)

func (g CHAT) IsValid() bool {
	switch g {
	case CHAT_GROUP, CHAT_INDIVIDUAL:
		return true
	}
	return false
}

type CHAT_STATUS string

const (
	CHAT_STATUS_YES        CHAT_STATUS = "YES"
	CHAT_STATUS_NOT        CHAT_STATUS = "NOT"
	CHAT_STATUS_INDECISION CHAT_STATUS = "-"
)

func (g CHAT_STATUS) IsValid() bool {
	switch g {
	case CHAT_STATUS_YES, CHAT_STATUS_NOT, CHAT_STATUS_INDECISION:
		return true
	}
	return false
}
