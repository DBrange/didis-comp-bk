package models

type NOTIFICATION_PRIORITY string

const (
	NOTIFICATION_PRIORITY_LOW      NOTIFICATION_PRIORITY = "LOW"
	NOTIFICATION_PRIORITY_MEDIUM   NOTIFICATION_PRIORITY = "MEDIUM"
	NOTIFICATION_PRIORITY_HIGH     NOTIFICATION_PRIORITY = "HIGH"
	NOTIFICATION_PRIORITY_CRITICAL NOTIFICATION_PRIORITY = "CRITICAL"
)

func (g NOTIFICATION_PRIORITY) IsValid() bool {
	switch g {
	case NOTIFICATION_PRIORITY_LOW,
		NOTIFICATION_PRIORITY_MEDIUM,
		NOTIFICATION_PRIORITY_HIGH,
		NOTIFICATION_PRIORITY_CRITICAL:
		return true
	}
	return false
}

type NOTIFICATION_TYPE string

const (
	NOTIFICATION_TYPE_UPDATE   NOTIFICATION_TYPE = "UPDATE"
	NOTIFICATION_TYPE_REMINDER NOTIFICATION_TYPE = "REMINDER"
	NOTIFICATION_TYPE_MESSAGE  NOTIFICATION_TYPE = "MESSAGE"
)

func (g NOTIFICATION_TYPE) IsValid() bool {
	switch g {
	case NOTIFICATION_TYPE_UPDATE,
		NOTIFICATION_TYPE_REMINDER,
		NOTIFICATION_TYPE_MESSAGE:
		return true
	}
	return false
}

type NOTIFICATION_STATE string

const (
	NOTIFICATION_STATE_READ   NOTIFICATION_STATE = "READ"
	NOTIFICATION_STATE_UNREAD NOTIFICATION_STATE = "UNREAD"
)

func (g NOTIFICATION_STATE) IsValid() bool {
	switch g {
	case NOTIFICATION_STATE_READ,
		NOTIFICATION_STATE_UNREAD:
		return true
	}
	return false
}
