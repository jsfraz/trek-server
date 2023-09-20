package handlers

import (
	"errors"
	"jsfraz/trek-server/database"
	"jsfraz/trek-server/models"

	"github.com/gin-gonic/gin"
)

// Create tracker.
//
//	@param c
//	@param request
//	@return *models.TrackerToken
//	@return error
func CreateTracker(c *gin.Context, request *models.CreateTracker) (*models.TrackerToken, error) {
	// check for superuser
	u, _ := c.Get("user")
	user := u.(*models.User)
	if !user.Superuser {
		c.AbortWithStatus(401)
		return nil, errors.New("user is not superuser")
	}
	// check if name is taken
	taken, err := database.TrackerExistsByName(request.Name)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	if taken {
		c.AbortWithStatus(409)
		return nil, errors.New("tracker with given name already exists")
	}
	// initialize tracker
	newTracker := models.NewTracker(request.Name)
	// insert
	token, err := database.InsertTracker(*newTracker)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return token, nil
}
