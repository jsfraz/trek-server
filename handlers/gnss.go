package handlers

import (
	"errors"
	"jsfraz/trek-server/database"
	"jsfraz/trek-server/models"

	"github.com/gin-gonic/gin"
)

// Get GNSS records by tracker.
//
//	@param c
//	@param request
//	@return *models.GNSSDataSummary
//	@return error
func GetAllGNSSRecords(c *gin.Context, request *models.GNSSAll) (*models.GNSSDataSummary, error) {
	// check if exists
	exists, err := database.TrackerExistsById(request.Id)
	// error
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	// not found
	if !exists {
		c.AbortWithStatus(404)
		return nil, errors.New("tracker does not exist")
	}
	// get data
	data, err := database.GetAllGNSSRecords(request.Id, request.Offset)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return data, nil
}

// Get GNSS records between from and to timestamps.
//
//	@param c
//	@param request
//	@return *models.GNSSDataSummary
//	@return error
func GetGNSSRecordsByTimestamps(c *gin.Context, request *models.GNSSFromTo) (*models.GNSSDataSummary, error) {
	// check if exists
	exists, err := database.TrackerExistsById(request.Id)
	// error
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	// not found
	if !exists {
		c.AbortWithStatus(404)
		return nil, errors.New("tracker does not exist")
	}
	// validate timestamps
	from, to, err := request.ValidateTimestamps()
	if err != nil {
		c.AbortWithStatus(400)
		return nil, err
	}
	// get data
	data, err := database.GetGNSSRecordsByTimestamps(request.Id, *from, *to, request.Offset)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return data, nil
}
