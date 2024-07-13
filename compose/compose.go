package compose

import (
	"fmt"

	"github.com/DBrange/didis-comp-bk/compose/dashboard"
	"github.com/DBrange/didis-comp-bk/database"
	location_adap_drivens "github.com/DBrange/didis-comp-bk/domains/location/adapters/drivens"
	location_adap_drivers "github.com/DBrange/didis-comp-bk/domains/location/adapters/drivers"
	location_services "github.com/DBrange/didis-comp-bk/domains/location/services"
	repo_adap_divers "github.com/DBrange/didis-comp-bk/domains/repository/adapters/drivers"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
	user_adap_drivens "github.com/DBrange/didis-comp-bk/domains/user/adapters/drivens"
	user_adap_drivers "github.com/DBrange/didis-comp-bk/domains/user/adapters/drivers"
	user_services "github.com/DBrange/didis-comp-bk/domains/user/services"
)

func Compose() (dashboard.Dashboard, error) {
	// List of all nesessary collections
	collections := []string{"users", "locations"}

	// Obtain collections
	collectionMap, err := database.GetCollections(collections)
	if err != nil {
		return nil, err
	}

	// Retrieve collections from the map
	userColl := collectionMap["users"]
	locationColl := collectionMap["locations"]

	// Create repository
	repository, err := repository.NewRepository(userColl, locationColl)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize repository: %v", err))
	}

	// Create repository drivers
	userManagerProxyAdapter := repo_adap_divers.NewUserManagerProxyAdapter(repository)
	locationManagerProxyAdapter := repo_adap_divers.NewLocationManagerProxyAdapter(repository)

	// Create repository drivens

	// Create user drivens
	userQueryerAdapter := user_adap_drivens.NewUserQueryerAdapter(userManagerProxyAdapter)
	// Create location drivens
	locationQueryerAdapter := location_adap_drivens.NewLocationQueryerAdapter(locationManagerProxyAdapter)

	// Create user service
	userServices := user_services.NewUserService(userQueryerAdapter)
	// Create location service
	locationServices := location_services.NewLocationService(locationQueryerAdapter)

	// Create user drivers
	userProxyAdapter := user_adap_drivers.NewUserProxyAdapter(userServices)
	// Create location drivers
	locationProxyAdapter := location_adap_drivers.NewLocationProxyAdapter(locationServices)

	// Create dashboard
	dashboard := dashboard.NewDashboardService(userProxyAdapter, locationProxyAdapter)

	return dashboard, nil
}
