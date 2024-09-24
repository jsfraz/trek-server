package database

import (
	"jsfraz/trek-server/models"
	"jsfraz/trek-server/utils"
)

// Insert tracker.
//
//	@param tracker
//	@return *models.TrackerToken
//	@return error
func InsertTracker(tracker models.Tracker) (*models.TrackerToken, error) {
	// transaction
	tx := utils.GetSingleton().PostgresDb.Begin()
	// insert to database
	if err := tx.Create(&tracker).Error; err != nil {
		tx.Rollback() // rollback the transaction if an error occurs
		return nil, err
	}
	// generate token
	token, err := utils.GenerateTrackerToken(tracker.Id)
	if err != nil {
		tx.Rollback() // rollback the transaction if an error occurs
		return nil, err
	}
	return models.NewTrackerToken(token), tx.Commit().Error
}

// Check if tracker with given name exists.
//
//	@param name
//	@return bool
//	@return error
func TrackerExistsByName(name string) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.Tracker{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// Check if tracker with given ID exists.
//
//	@param id
//	@return bool
//	@return error
func TrackerExistsById(id uint64) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.Tracker{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// Get tracker by ID.
//
//	@param id
//	@return *models.Tracker
//	@return error
func GetTrackerById(id uint64) (*models.Tracker, error) {
	var tracker models.Tracker
	err := utils.GetSingleton().PostgresDb.Model(&models.Tracker{}).Where("id = ?", id).First(&tracker).Error
	if err != nil {
		return nil, err
	}
	return &tracker, nil
}

// Get all trackers.
//
//	@return *[]models.Tracker
//	@return error
func GetAllTrackers() (*[]models.Tracker, error) {
	var trackers []models.Tracker = []models.Tracker{}
	err := utils.GetSingleton().PostgresDb.Model(&models.Tracker{}).Order("id ASC").Find(&trackers).Error
	if err != nil {
		return nil, err
	}
	return &trackers, nil
}

// Delete tracker with given ID.
//
//	@param ids
//	@return error
func DeleteTracker(id uint64) error {
	// transaction
	tx := utils.GetSingleton().PostgresDb.Begin()
	// delete GNSS data
	if err := tx.Where("tracker_id = ?", id).Delete(&models.GNSSData{}).Error; err != nil {
		tx.Rollback() // rollback the transaction if an error occurs
		return err
	}
	// delete tracker
	if err := tx.Where("id = ?", id).Delete(&models.Tracker{}).Error; err != nil {
		tx.Rollback() // rollback the transaction if an error occurs
		return err
	}
	return tx.Commit().Error
}

// Set tracker name.
//
//	@param trackerId
//	@param name
//	@return error
func SetTrackerName(trackerId uint64, name string) error {
	var tracker models.Tracker
	err := utils.GetSingleton().PostgresDb.Model(&models.Tracker{}).Where("id = ?", trackerId).First(&tracker).Error
	if err != nil {
		return err
	}
	// update
	tracker.Name = name
	return utils.GetSingleton().PostgresDb.Save(&tracker).Error
}
