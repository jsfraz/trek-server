package database

import (
	"jsfraz/trek-server/models"
	"jsfraz/trek-server/utils"
)

// Insert GNSS data into database.
//
//	@param data
//	@return error
func InsertGNSSData(data models.GNSSData) error {
	err := utils.GetSingleton().PostgresDb.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

// Check if GNSS data with timestamp exists.
//
//	@param timestamp
//	@return bool
//	@return error
func GNSSDataExists(trackerId uint64, timestamp string) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.GNSSData{}).Where("tracker_id = ? AND timestamp = ?", trackerId, timestamp).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}
