package models

type ROLE string

const (
	RoleAdmin  ROLE = "ADMIN"
	BasicAdmin ROLE = "BASIC"
	FreeAdmin  ROLE = "FREE"
)