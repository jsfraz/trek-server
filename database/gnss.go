package database

import (
	"jsfraz/trek-server/models"
	"jsfraz/trek-server/utils"
	"time"
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

// Get GNSS records by tracker ID.
//
//	@param trackerId
//	@return *[]models.GNSSData
//	@return error
func GetAllGNSSRecords(trackerId uint64) (*[]models.GNSSData, error) {
	var data []models.GNSSData = []models.GNSSData{}
	err := utils.GetSingleton().PostgresDb.Model(&models.GNSSData{}).Where("tracker_id = ?", trackerId).Order("timestamp ASC").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get GNSS records by tracker ID and timestamps.
//
//	@param trackerId
//	@param from
//	@param to
//	@return *[]models.GNSSData
//	@return error
func GetGNSSRecordsByTimestamps(trackerId uint64, from time.Time, to time.Time) (*[]models.GNSSData, error) {
	var data []models.GNSSData = []models.GNSSData{}
	err := utils.GetSingleton().PostgresDb.Model(&models.GNSSData{}).Where("tracker_id = ? AND timestamp >= ? AND timestamp <= ?", trackerId, from, to).Order("timestamp ASC").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
