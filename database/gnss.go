package database

import (
	"fmt"
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
//	@param offset
//	@return *models.GNSSDataSummary
//	@return error
func GetAllGNSSRecords(trackerId uint64, offset int) (*models.GNSSDataSummary, error) {
	var data []models.GNSSData = []models.GNSSData{}
	var err error
	if offset > 1 {
		// offset is greater than 1
		query := fmt.Sprintf(`
				SELECT *
				FROM (
					SELECT *,
						   ROW_NUMBER() OVER (ORDER BY timestamp ASC) AS row_num
					FROM "gnss_data"
					WHERE tracker_id = %d
				) AS numbered_rows
				WHERE row_num %% %d = 1
				ORDER BY timestamp ASC;
			`, trackerId, offset)
		err = utils.GetSingleton().PostgresDb.Raw(query).Scan(&data).Error
	} else {
		// no offset
		err = utils.GetSingleton().PostgresDb.Model(&models.GNSSData{}).Where("tracker_id = ?", trackerId).Order("timestamp ASC").Find(&data).Error
	}
	// error
	if err != nil {
		return nil, err
	}
	// create summary
	// min
	min, err := getMinSpeed(trackerId)
	if err != nil {
		return nil, err
	}
	// avg
	avg, err := getAvgSpeed(trackerId)
	if err != nil {
		return nil, err
	}
	// max
	max, err := getMaxSpeed(trackerId)
	if err != nil {
		return nil, err
	}
	return models.NewGNSSDataSummary(trackerId, data, min, avg, max), nil
}

// Get GNSS records by tracker ID and timestamps.
//
//	@param trackerId
//	@param from
//	@param to
//	@param offset
//	@return *models.GNSSDataSummary
//	@return error
func GetGNSSRecordsByTimestamps(trackerId uint64, from time.Time, to time.Time, offset int) (*models.GNSSDataSummary, error) {
	var data []models.GNSSData = []models.GNSSData{}
	var err error
	if offset > 1 {
		// offset is greater than 1
		query := fmt.Sprintf(`
		SELECT *
		FROM (
			SELECT *,
				   ROW_NUMBER() OVER (ORDER BY timestamp ASC) AS row_num
			FROM "gnss_data"
			WHERE tracker_id = %d
				AND timestamp >= '%s'
				AND timestamp <= '%s'
		) AS numbered_rows
		WHERE row_num %% %d = 1
		ORDER BY timestamp ASC;
	`, trackerId, from.Format("2006-01-02 15:04:05.999"), to.Format("2006-01-02 15:04:05.999"), offset)
		err = utils.GetSingleton().PostgresDb.Raw(query).Scan(&data).Error
	} else {
		// no offset
		err = utils.GetSingleton().PostgresDb.Model(&models.GNSSData{}).Where("tracker_id = ? AND timestamp >= ? AND timestamp <= ?", trackerId, from, to).Order("timestamp ASC").Find(&data).Error
	}
	// error
	if err != nil {
		return nil, err
	}
	// create summary
	// min
	min, err := getMinSpeedFromTo(trackerId, from, to)
	if err != nil {
		return nil, err
	}
	// avg
	avg, err := getAvgSpeedFromTo(trackerId, from, to)
	if err != nil {
		return nil, err
	}
	// max
	max, err := getMaxSpeedFromTo(trackerId, from, to)
	if err != nil {
		return nil, err
	}
	return models.NewGNSSDataSummary(trackerId, data, min, avg, max), nil
}

// Get min speed.
//
//	@param trackerId
//	@return float64
//	@return error
func getMinSpeed(trackerId uint64) (float64, error) {
	var speed float64
	if err := utils.GetSingleton().PostgresDb.Model(&models.GNSSData{}).Select("MIN(speed)").Where("tracker_id = ?", trackerId).Scan(&speed).Error; err != nil {
		return 0, err
	}
	return speed, nil
}

// Get avg speed.
//
//	@param trackerId
//	@return float64
//	@return error
func getAvgSpeed(trackerId uint64) (float64, error) {
	var speed float64
	if err := utils.GetSingleton().PostgresDb.Model(&models.GNSSData{}).Select("AVG(speed)").Where("tracker_id = ?", trackerId).Scan(&speed).Error; err != nil {
		return 0, err
	}
	return speed, nil
}

// Get max speed.
//
//	@param trackerId
//	@return float64
//	@return error
func getMaxSpeed(trackerId uint64) (float64, error) {
	var speed float64
	if err := utils.GetSingleton().PostgresDb.Model(&models.GNSSData{}).Select("MAX(speed)").Where("tracker_id = ?", trackerId).Scan(&speed).Error; err != nil {
		return 0, err
	}
	return speed, nil
}

// Get min speed between two timestamps.
//
//	@param trackerId
//	@param from
//	@param to
//	@return float64
//	@return error
func getMinSpeedFromTo(trackerId uint64, from time.Time, to time.Time) (float64, error) {
	var speed float64
	if err := utils.GetSingleton().PostgresDb.Model(&models.GNSSData{}).Select("MIN(speed)").Where("tracker_id = ? AND timestamp >= ? AND timestamp <= ?", trackerId, from, to).Scan(&speed).Error; err != nil {
		return 0, err
	}
	return speed, nil
}

// Get avg speed between two timestamps.
//
//	@param trackerId
//	@param from
//	@param to
//	@return float64
//	@return error
func getAvgSpeedFromTo(trackerId uint64, from time.Time, to time.Time) (float64, error) {
	var speed float64
	if err := utils.GetSingleton().PostgresDb.Model(&models.GNSSData{}).Select("AVG(speed)").Where("tracker_id = ? AND timestamp >= ? AND timestamp <= ?", trackerId, from, to).Scan(&speed).Error; err != nil {
		return 0, err
	}
	return speed, nil
}

// Get max speed between two timestamps.
//
//	@param trackerId
//	@param from
//	@param to
//	@return float64
//	@return error
func getMaxSpeedFromTo(trackerId uint64, from time.Time, to time.Time) (float64, error) {
	var speed float64
	if err := utils.GetSingleton().PostgresDb.Model(&models.GNSSData{}).Select("MAX(speed)").Where("tracker_id = ? AND timestamp >= ? AND timestamp <= ?", trackerId, from, to).Scan(&speed).Error; err != nil {
		return 0, err
	}
	return speed, nil
}

// Get first GNSS record created max 1,25 seconds ago.
//
//	@param trackerId
//	@return *models.GNSSData
//	@return error
func GetCurrentGNSSData(trackerId uint64) (*models.GNSSData, error) {
	var data models.GNSSData
	duration := time.Duration(1.25 * float64(time.Second))
	err := utils.GetSingleton().PostgresDb.Model(&models.GNSSData{}).Where("timestamp > ? AND tracker_id = ?", time.Now().Add(-duration).Format("2006-01-02 15:04:05.999"), trackerId).First(&data).Error
	// Cancel error if record was not found so nil can be returned.
	if err.Error() == "record not found" {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &data, nil
}
