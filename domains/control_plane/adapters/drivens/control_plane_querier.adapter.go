package drivens

import ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"

type ControlPlaneQuerierAdapter struct {
	adapter ports.ForManagingControlPlane
}

func NewControlPlaneQuerierAdapter(adapter ports.ForManagingControlPlane) *ControlPlaneQuerierAdapter {
	return &ControlPlaneQuerierAdapter{
		adapter: adapter,
	}
}
