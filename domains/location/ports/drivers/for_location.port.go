package ports

import "github.com/DBrange/didis-comp-bk/domains/location/ports/drivers/interfaces"

type ForLocation interface {
	interfaces.CreateLocation
	interfaces.GetLocation
	interfaces.UpdateLocation
	interfaces.DeleteLocation
}