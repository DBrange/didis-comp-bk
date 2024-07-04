package ports

import "github.com/DBrange/didis-comp-bk/internal/user/ports/drivers/interfaces"

type ForUser interface {
	interfaces.CreateUser
	interfaces.GetUserByID
}
