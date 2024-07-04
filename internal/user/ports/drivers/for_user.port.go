package ports

import "didis-comp-bk/internal/user/ports/drivers/interfaces"

type ForUser interface {
	interfaces.CreateUser
	interfaces.GetUserByID
}