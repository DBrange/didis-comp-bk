package interfaces

import "github.com/DBrange/didis-comp-bk/domains/location/models/dto"

type CreateLocation interface {
	CreateLocation(location *dto.CreateLocationDTOReq) (string, error)
}
