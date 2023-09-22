package models

import (
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

const iso8601RegexPattern = `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?([+-]\d{2}:\d{2}|Z)?$`

type GNSSDataInput struct {
	Latitude  float64 `json:"latitude" validate:"latitude,required"`
	Longitude float64 `json:"longitude" validate:"longitude,required"`
	Speed     float64 `json:"speed" validate:"min=0"`
	Timestamp string  `json:"timestamp" validate:"required"`
}

// Checks ISO 8601 timestamp.
//
//	@receiver g
//	@param timestamp
//	@return bool
func (g GNSSDataInput) ValidateISO8601Timestamp() bool {
	regex := regexp.MustCompile(iso8601RegexPattern)
	return regex.MatchString(g.Timestamp)
}

// Parse map into struct.
//
//	@param mapData
//	@return *GNSSData
//	@return error
func ParseMap(mapData map[string]interface{}) (*GNSSDataInput, error) {
	var data GNSSDataInput
	errStr := "Invalid field: "
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
	timestampValid := data.ValidateISO8601Timestamp()
	if !timestampValid {
		return nil, errors.New("Invalid ISO 8601 timestamp: '" + data.Timestamp + "'")
	}
	return &data, nil
}

// Return GNSSDataDb.
//
//	@receiver g
//	@param trackerId
//	@return *GNSSData
func (g GNSSDataInput) ToDatabaseModel(trackerId uint64) *GNSSData {
	gDb := new(GNSSData)
	gDb.TrackerId = trackerId
	gDb.Latitude = g.Latitude
	gDb.Longitude = g.Longitude
	gDb.Speed = g.Speed
	gDb.Timestamp = g.Timestamp
	return gDb
}
