package handlers

import (
	"errors"
	"jsfraz/trek-server/database"
	"jsfraz/trek-server/models"
	"jsfraz/trek-server/utils"

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

// Regenerate tracker token.
//
//	@param c
//	@param request
//	@return error
func RegenerateTrackerToken(c *gin.Context, request *models.Id) (*models.TrackerToken, error) {
	// check fo superuser
	u, _ := c.Get("user")
	user := u.(*models.User)
	if !user.Superuser {
		c.AbortWithStatus(401)
		return nil, errors.New("user is not superuser")
	}
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
	// update
	token, err := utils.GenerateTrackerToken(request.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	err = database.SetTrackerToken(request.Id, token)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return models.NewTrackerToken(token), err
}

// Get all trackers.
//
//	@param c
//	@return *[]models.Tracker
//	@return error
func GetAllTrackers(c *gin.Context) (*[]models.Tracker, error) {
	// get trackers
	trackers, err := database.GetAllTrackers()
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return trackers, nil
}

// Delete tracker.
//
//	@param c
//	@param id
//	@return error
func DeleteTracker(c *gin.Context, id *models.Id) error {
	// check for superuser
	u, _ := c.Get("user")
	user := u.(*models.User)
	if !user.Superuser {
		c.AbortWithStatus(401)
		return errors.New("user is not superuser")
	}
	// delete tracker
	err := database.DeleteTracker(id.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Update tracker name.
//
//	@param c
//	@param request
//	@return error
func UpdateTrackerName(c *gin.Context, request *models.UpdateTrackerName) error {
	// check fo superuser
	u, _ := c.Get("user")
	user := u.(*models.User)
	if !user.Superuser {
		c.AbortWithStatus(401)
		return errors.New("user is not superuser")
	}
	// check if exists
	exists, err := database.TrackerExistsById(request.Id)
	// error
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	// not found
	if !exists {
		c.AbortWithStatus(404)
		return errors.New("tracker does not exist")
	}
	// update
	err = database.SetTrackerName(request.Id, request.Name)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}
