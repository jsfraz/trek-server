package models

import "github.com/go-playground/validator/v10"

type GetCurrentGNSSDataInput struct {
	Id uint64 `json:"id" validate:"required"`
}

// Parse map into struct.
//
//	@param mapData
//	@return *GetCurrentGNSSDataInput
//	@return error
func ParseGetCurrentGNSSDataInput(mapData map[string]interface{}) (*GetCurrentGNSSDataInput, error) {
	var input GetCurrentGNSSDataInput
	if v, ok := mapData["id"].(float64); ok {
		input.Id = uint64(v)
	}
	// Validation
	validator := validator.New()
	err := validator.Struct(input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}
