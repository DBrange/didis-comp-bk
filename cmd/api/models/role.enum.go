package models

type ROLE string

const (
	ROLE_ADMIN ROLE = "ADMIN"
	ROLE_BASIC ROLE = "BASIC"
	ROLE_FREE  ROLE = "FREE"
)


type ROLE_TYPE string

const (
	ROLE_TYPE_USER ROLE_TYPE = "USER"
	ROLE_TYPE_TEAM ROLE_TYPE = "TEAM"
)


