package compose

import (
	"context"
	"didis-comp-bk/cmd/api/compose/dashboard"
	"didis-comp-bk/database"
	repo_adap_divers "didis-comp-bk/internal/repository/adapters/drivers"
	"didis-comp-bk/internal/repository/repository"
	user_adap_drivens "didis-comp-bk/internal/user/adapters/drivens"
	user_adap_drivers "didis-comp-bk/internal/user/adapters/drivers"
	"didis-comp-bk/internal/user/services"
	"errors"
)

func Compose() (dashboard.Dashboard, error) {
	ctx := context.Background()
	coll := database.GetCollection("users")
	if coll == nil {
		return nil, errors.New("failed to get collection")
	}

	// Create repository
	repository := repository.NewRepository(coll)

	// Create repository drivers
	userManagerProxyAdapter := repo_adap_divers.NewUserMangerProxyAdapter(ctx, repository)

	// Create repository drivens

	// Create user drivens
	userQueryerAdapter := user_adap_drivens.NewUserQueryerAdapter(ctx, userManagerProxyAdapter)

	// Create user service
	userServices := services.NewUserService(userQueryerAdapter)

	// Create user drivers
	userAdapter := user_adap_drivers.NewUserAdapter(ctx, userServices)

	// Create dashboard
	dashboard := dashboard.NewDashboardService(userAdapter)

	return dashboard, nil
}
