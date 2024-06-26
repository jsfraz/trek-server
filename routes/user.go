package routes

import (
	"jsfraz/trek-server/handlers"
	"jsfraz/trek-server/middlewares"
	"jsfraz/trek-server/utils"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

// Returns user route.
//
//	@param g
func UserRoute(g *fizz.RouterGroup) {
	// user router
	grp := g.Group("user", "User", "Users")
	// auth middleware
	grp.Use(middlewares.Auth)

	// create user
	grp.POST("", utils.CreateOperationOption("Create user.", true), tonic.Handler(handlers.CreateUser, 204))
	// get current user
	grp.GET("whoami", utils.CreateOperationOption("Get current user.", true), tonic.Handler(handlers.WhoAmI, 200))
	// get all users
	grp.GET("all", utils.CreateOperationOption("Get all users.", true), tonic.Handler(handlers.GetAllUsers, 200))
	// delete user
	grp.DELETE("", utils.CreateOperationOption("Delete user.", true), tonic.Handler(handlers.DeleteUser, 204))
	// update user
	grp.PATCH("", utils.CreateOperationOption("Update user.", true), tonic.Handler(handlers.UpdateUser, 204))

	// GNSS route
	GNSSRoute(grp)
}
