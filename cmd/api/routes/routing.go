package routes

import (
	"time"

	// 	"context"
	// "errors"

	// "github.com/DBrange/didis-comp-bk/database"
	// repo_adap_divers "github.com/DBrange/didis-comp-bk/internal/repository/adapters/drivers"
	// "github.com/DBrange/didis-comp-bk/internal/repository/repository"
	// user_adap_drivens "github.com/DBrange/didis-comp-bk/internal/user/adapters/drivens"
	// user_adap_drivers "github.com/DBrange/didis-comp-bk/internal/user/adapters/drivers"
	// "github.com/DBrange/didis-comp-bk/internal/user/services"
	handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/users"
	"github.com/DBrange/didis-comp-bk/compose"
	"github.com/DBrange/didis-comp-bk/compose/dashboard"

	// user_ports_drivers "github.com/DBrange/didis-comp-bk/internal/user/ports/drivers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//	type Dashboard interface {
//		User() user_ports_drivers.ForUser
//	}
func NewRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))
	dashboard, _ := compose.Compose()
	RoutesHandler(router, dashboard)
	return router
}

func RoutesHandler(router *gin.Engine, dashboard dashboard.Dashboard) {
	userRoutes(router, handlers.NewHandlerUser(dashboard.User()))
}

// func Compose() (Dashboard, error) {
// 	ctx := context.Background()
// 	coll := database.GetCollection("users")
// 	if coll == nil {
// 		return nil, errors.New("failed to get collection")
// 	}

// 	// Create repository
// 	repository := repository.NewRepository(coll)

// 	// Create repository drivers
// 	userManagerProxyAdapter := repo_adap_divers.NewUserMangerProxyAdapter(ctx, repository)

// 	// Create repository drivens

// 	// Create user drivens
// 	userQueryerAdapter := user_adap_drivens.NewUserQueryerAdapter(ctx, userManagerProxyAdapter)

// 	// Create user service
// 	userServices := services.NewUserService(userQueryerAdapter)

// 	// Create user drivers
// 	userAdapter := user_adap_drivers.NewUserAdapter(ctx, userServices)

// 	// Create dashboard
// 	dashboard := NewDashboardService(userAdapter)

// 	return dashboard, nil
// }

// type DashboardService struct {
// 	ForUser user_ports_drivers.ForUser
// }

// func NewDashboardService(UserAdater *user_adap_drivers.UserAdapter) *DashboardService {
// 	return &DashboardService{
// 		ForUser: UserAdater,
// 	}
// }

// func (d *DashboardService) User() user_ports_drivers.ForUser {
// 	return d.ForUser
// }
