package routes

import (
	"jsfraz/trek-server/handlers"
	"jsfraz/trek-server/middlewares"
	"jsfraz/trek-server/utils"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

// Returns tracker route.
//
//	@param g
func TrackerRoute(g *fizz.RouterGroup) {
	// tracker router
	grp := g.Group("tracker", "Tracker", "Trackers")
	// auth middleware
	grp.Use(middlewares.Auth)

	// create tracker
	grp.POST("", utils.CreateOperationOption("Create tracker.", true), tonic.Handler(handlers.CreateTracker, 200))
	// regenerate tracker token
	grp.PATCH("token", utils.CreateOperationOption("Regenerate tracker token.", true), tonic.Handler(handlers.RegenerateTrackerToken, 200))
	// get all trackers
	grp.GET("all", utils.CreateOperationOption("Get all trackers.", true), tonic.Handler(handlers.GetAllTrackers, 200))
	// delete tracker
	grp.DELETE("", utils.CreateOperationOption("Delete tracker.", true), tonic.Handler(handlers.DeleteTracker, 204))
	// update tracker name
	grp.PATCH("name", utils.CreateOperationOption("Update tracker name.", true), tonic.Handler(handlers.UpdateTrackerName, 204))
}
