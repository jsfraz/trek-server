package routes

import (
	"jsfraz/trek-server/handlers"
	"jsfraz/trek-server/utils"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

// Returns auth route.
//
//	@param g
func AuthRoute(g *fizz.RouterGroup) {
	// auth route
	grp := g.Group("auth", "Authentication", "Authentication")

	// login
	grp.POST("login", utils.CreateOperationOption("User login.", false), tonic.Handler(handlers.Login, 200))
}
