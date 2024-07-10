package compose

import (
	"context"
	// "errors"

	"github.com/DBrange/didis-comp-bk/compose/dashboard"
	"github.com/DBrange/didis-comp-bk/database"
	location_adap_drivers "github.com/DBrange/didis-comp-bk/domains/location/adapters/drivers"
	location_services "github.com/DBrange/didis-comp-bk/domains/location/services"
	repo_adap_divers "github.com/DBrange/didis-comp-bk/domains/repository/adapters/drivers"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
	user_adap_drivens "github.com/DBrange/didis-comp-bk/domains/user/adapters/drivens"
	location_adap_drivens "github.com/DBrange/didis-comp-bk/domains/location/adapters/drivens"
	user_adap_drivers "github.com/DBrange/didis-comp-bk/domains/user/adapters/drivers"
	user_services "github.com/DBrange/didis-comp-bk/domains/user/services"
)

func Compose() (dashboard.Dashboard, error) {
	ctx := context.Background()
	user_coll := database.GetCollection("users")
	location_coll := database.GetCollection("locations")

	// if coll == nil {
	// 	return nil, errors.New("failed to get collection")
	// }

	// Create repository
	repository := repository.NewRepository(user_coll, location_coll)

	// Create repository drivers
	userManagerProxyAdapter := repo_adap_divers.NewUserMangerProxyAdapter(repository)
	locationManagerProxyAdapter := repo_adap_divers.NewLocationMangerProxyAdapter(repository)

	// Create repository drivens

	// Create user drivens
	userQueryerAdapter := user_adap_drivens.NewUserQueryerAdapter(userManagerProxyAdapter)
	locationQueryerAdapter := location_adap_drivens.NewLocationQueryerAdapter(locationManagerProxyAdapter)

	// Create user service
	userServices := user_services.NewUserService(userQueryerAdapter)
	locationServices := location_services.NewLocationService(locationQueryerAdapter)

	// Create user drivers
	userProxyAdapter := user_adap_drivers.NewUserProxyAdapter(ctx, userServices)
	locationProxyAdapter := location_adap_drivers.NewLocationProxyAdapter(ctx, locationServices)

	// Create dashboard
	dashboard := dashboard.NewDashboardService(userProxyAdapter, locationProxyAdapter)

	return dashboard, nil
}
