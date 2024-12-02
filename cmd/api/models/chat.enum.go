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

type CHAT_AVAILABILITY_STATUS string

const (
	CHAT_AVAILABILITY_STATUS_YES        CHAT_AVAILABILITY_STATUS = "YES"
	CHAT_AVAILABILITY_STATUS_NOT        CHAT_AVAILABILITY_STATUS = "NOT"
	CHAT_AVAILABILITY_STATUS_INDECISION CHAT_AVAILABILITY_STATUS = "-"
)

func (g CHAT_AVAILABILITY_STATUS) IsValid() bool {
	switch g {
	case CHAT_AVAILABILITY_STATUS_YES, CHAT_AVAILABILITY_STATUS_NOT, CHAT_AVAILABILITY_STATUS_INDECISION:
		return true
	}
	return false
}
