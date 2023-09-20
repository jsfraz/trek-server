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
	// TODO regenerate API key
	// TODO get tracker(s)
	// TODO delete tracker(s)
}
