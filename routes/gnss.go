package routes

import (
	"jsfraz/trek-server/handlers"
	"jsfraz/trek-server/utils"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

func GNSSRoute(g *fizz.RouterGroup) {
	// gnss route
	grp := g.Group("gnss", "GNSS data", "GNSS data")

	// get all GNSS records for tracker
	grp.GET("all", utils.CreateOperationOption("Get all GNSS records for tracker.", true), tonic.Handler(handlers.GetAllGNSSRecords, 200))
	// get GNSS records between two dates for tracker
	grp.GET("all/fromTo", utils.CreateOperationOption("Get GNSS records between two dates for tracker.", true), tonic.Handler(handlers.GetGNSSRecordsByTimestamps, 200))
}
