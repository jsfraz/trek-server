package models

import (
	"errors"
	"jsfraz/trek-server/utils"

	"github.com/go-playground/validator/v10"
)

type GNSSDataInput struct {
	Latitude  float64 `json:"latitude" validate:"latitude"`
	Longitude float64 `json:"longitude" validate:"longitude"`
	Speed     float64 `json:"speed" validate:"min=0"`
	Timestamp string  `json:"timestamp" validate:"required"`
}

// Parse map into struct.
//
//	@param mapData
//	@return *GNSSData
//	@return error
func ParseMap(mapData map[string]interface{}) (*GNSSDataInput, error) {
	var data GNSSDataInput
	errStr := "invalid field: "
	// iterate trough values
	for key, value := range mapData {
		switch key {
		case "latitude":
			if v, ok := value.(float64); ok {
				data.Latitude = v
			} else {
				return nil, errors.New(errStr + "latitude")
			}
		case "longitude":
			if v, ok := value.(float64); ok {
				data.Longitude = v
			} else {
				return nil, errors.New(errStr + "longitude")
			}
		case "speed":
			if v, ok := value.(float64); ok {
				data.Speed = v
			} else {
				return nil, errors.New(errStr + "speed")
			}
		case "timestamp":
			if v, ok := value.(string); ok {
				data.Timestamp = v
			} else {
				return nil, errors.New(errStr + "timestamp")
			}
		}
	}
	// validation
	validator := validator.New()
	err := validator.Struct(data)
	if err != nil {
		return nil, err
	}
	// timestamp validation
	_, err = utils.ParseISO8601String(data.Timestamp)
	if err != nil {
		return nil, errors.New("invalid ISO 8601 timestamp: '" + data.Timestamp + "'")
	}
	return &data, nil
}

// Return GNSSDataDb.
//
//	@receiver g
//	@param trackerId
//	@return *GNSSData
func (g GNSSDataInput) ToDatabaseModel(trackerId uint64) (*GNSSData, error) {
	gDb := new(GNSSData)
	gDb.TrackerId = trackerId
	gDb.Latitude = g.Latitude
	gDb.Longitude = g.Longitude
	gDb.Speed = g.Speed
	timestamp, err := utils.ParseISO8601String(g.Timestamp)
	if err != nil {
		return nil, err
	}
	gDb.Timestamp = *timestamp
	return gDb, nil
}
