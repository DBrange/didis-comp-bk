package compose

import (
	"fmt"

	"github.com/DBrange/didis-comp-bk/compose/dashboard"
	"github.com/DBrange/didis-comp-bk/database"
	category_adap_drivens "github.com/DBrange/didis-comp-bk/domains/category/adapters/drivens"
	category_adap_drivers "github.com/DBrange/didis-comp-bk/domains/category/adapters/drivers"
	category_services "github.com/DBrange/didis-comp-bk/domains/category/services"
	control_plane_adap_drivens "github.com/DBrange/didis-comp-bk/domains/control_plane/adapters/drivens"
	control_plane_adap_drivers "github.com/DBrange/didis-comp-bk/domains/control_plane/adapters/drivers"
	control_plane_services "github.com/DBrange/didis-comp-bk/domains/control_plane/services"
	location_adap_drivens "github.com/DBrange/didis-comp-bk/domains/location/adapters/drivens"
	location_adap_drivers "github.com/DBrange/didis-comp-bk/domains/location/adapters/drivers"
	location_services "github.com/DBrange/didis-comp-bk/domains/location/services"
	profile_adap_drivens "github.com/DBrange/didis-comp-bk/domains/profile/adapters/drivens"
	profile_adap_drivers "github.com/DBrange/didis-comp-bk/domains/profile/adapters/drivers"
	profile_services "github.com/DBrange/didis-comp-bk/domains/profile/services"
	repo_adap_divers "github.com/DBrange/didis-comp-bk/domains/repository/adapters/drivers"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
	tournament_adap_drivens "github.com/DBrange/didis-comp-bk/domains/tournament/adapters/drivens"
	tournament_adap_drivers "github.com/DBrange/didis-comp-bk/domains/tournament/adapters/drivers"
	tournament_services "github.com/DBrange/didis-comp-bk/domains/tournament/services"
)

func Compose() (dashboard.Dashboard, error) {
	// List of all nesessary collections
	collections := []string{
		"users",
		"locations",
		"availabilities",
		"roles",
		"organizers",
		"categories",
		"tournaments",
		"pots",
		"tournament_groups",
		"double_eliminations",
		"rounds",
		"competitors",
		"competitor_stats",
		"singles",
		"doubles",
		"teams",
		"matches",
		"chats",
		"chat_messages",
		"guest_users",
		"notifications",

		"competitor_matches",       // Intermediate table
		"competitor_users",         // Intermediate table
		"followers",                // Intermediate table
		"guest_competitors",        // Intermediate table
		"category_registrations",   // Intermediate table
		"opinions",                 // Intermediate table
		"tournament_registrations", // Intermediate table
		"user_chats",               // Intermediate table
	}

	// Obtain collections
	collectionMap, err := database.GetCollections(collections)
	if err != nil {
		return nil, err
	}

	// Retrieve collections from the map
	userColl := collectionMap["users"]
	locationColl := collectionMap["locations"]
	availabilityColl := collectionMap["availabilities"]
	roleColl := collectionMap["roles"]
	organizerColl := collectionMap["organizers"]
	categoryColl := collectionMap["categories"]
	tournamentColl := collectionMap["tournaments"]
	potColl := collectionMap["pots"]
	tournamentGroupColl := collectionMap["tournament_groups"]
	doubleEliminationColl := collectionMap["double_eliminations"]
	roundColl := collectionMap["rounds"]
	competitorColl := collectionMap["competitors"]
	competitorStatsColl := collectionMap["competitor_stats"]
	singleColl := collectionMap["singles"]
	doubleColl := collectionMap["doubles"]
	teamColl := collectionMap["teams"]
	matchColl := collectionMap["matches"]
	chatColl := collectionMap["chats"]
	chatMessageColl := collectionMap["chat_messages"]
	guestUserColl := collectionMap["guest_users"]
	notificationColl := collectionMap["notifications"]

	competitorMatchColl := collectionMap["competitor_matches"]              // Intermediate table
	competitorUserColl := collectionMap["competitor_users"]                 // Intermediate table
	followerColl := collectionMap["followers"]                              // Intermediate table
	guestCompetitorColl := collectionMap["guest_competitors"]               // Intermediate table
	categoryRegistrationColl := collectionMap["category_registrations"]     // Intermediate table
	opinionColl := collectionMap["opinions"]                                // Intermediate table
	tournamentRegistrationColl := collectionMap["tournament_registrations"] // Intermediate table
	userChatColl := collectionMap["user_chats"]                             // Intermediate table

	// Create repository
	repository, err := repository.NewRepository(
		database.MongoClient,
		userColl,
		locationColl,
		availabilityColl,
		roleColl,
		organizerColl,
		categoryColl,
		tournamentColl,
		potColl,
		tournamentGroupColl,
		doubleEliminationColl,
		roundColl,
		competitorColl,
		competitorStatsColl,
		singleColl,
		doubleColl,
		teamColl,
		matchColl,
		chatColl,
		chatMessageColl,
		guestUserColl,
		notificationColl,

		competitorMatchColl,        // Intermediate table
		competitorUserColl,         // Intermediate table
		followerColl,               // Intermediate table
		guestCompetitorColl,        // Intermediate table
		categoryRegistrationColl,   // Intermediate table
		opinionColl,                // Intermediate table
		tournamentRegistrationColl, // Intermediate table
		userChatColl,               // Intermediate table
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize repository: %v", err))
	}

	// Create repository drivers
	controlPlaneManagerProxyAdapter := repo_adap_divers.NewControlPlaneManagerProxyAdapter(repository)
	profileManagerProxyAdapter := repo_adap_divers.NewProfileManagerProxyAdapter(repository)
	locationManagerProxyAdapter := repo_adap_divers.NewLocationManagerProxyAdapter(repository)
	tournamentManagerProxyAdapter := repo_adap_divers.NewTournamentManagerProxyAdapter(repository)
	categoryManagerProxyAdapter := repo_adap_divers.NewCategoryManagerProxyAdapter(repository)

	// Create repository drivens

	// Create drivens
	controlPlaneQuerierAdapter := control_plane_adap_drivens.NewControlPlaneQuerierAdapter(controlPlaneManagerProxyAdapter)
	userQueryerAdapter := profile_adap_drivens.NewProfileQueryerAdapter(profileManagerProxyAdapter)
	locationQueryerAdapter := location_adap_drivens.NewLocationQueryerAdapter(locationManagerProxyAdapter)
	tournamentQueryerAdapter := tournament_adap_drivens.NewTournamentQueryerAdapter(tournamentManagerProxyAdapter)
	categoryQueryerAdapter := category_adap_drivens.NewCategoryQueryerAdapter(categoryManagerProxyAdapter)

	// Create services
	controlPlaneServices := control_plane_services.NewControlPlaneService(controlPlaneQuerierAdapter)
	userServices := profile_services.NewProfileService(userQueryerAdapter)
	locationServices := location_services.NewLocationService(locationQueryerAdapter)
	tournamentServices := tournament_services.NewTournamentService(tournamentQueryerAdapter)
	categoryServices := category_services.NewCategoryService(categoryQueryerAdapter)

	// Create  drivers
	controlPlaneProxyAdapter := control_plane_adap_drivers.NewControlPlaneProxyAdapter(controlPlaneServices)
	profileProxyAdapter := profile_adap_drivers.NewProfileProxyAdapter(userServices)
	locationProxyAdapter := location_adap_drivers.NewLocationProxyAdapter(locationServices)
	tournamentProxyAdapter := tournament_adap_drivers.NewTournamentProxyAdapter(tournamentServices)
	categoryProxyAdapter := category_adap_drivers.NewCategoryProxyAdapter(categoryServices)

	// Create dashboard
	dashboard := dashboard.NewDashboardService(controlPlaneProxyAdapter, profileProxyAdapter, locationProxyAdapter, tournamentProxyAdapter, categoryProxyAdapter)

	return dashboard, nil
}
