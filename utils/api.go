package utils

import (
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

// Returns array of Fizz Operation option with summary and error responses
//
//	@param summary
//	@param useSecurity
//	@return []fizz.OperationOption
func CreateOperationOption(summary string, useSecurity bool) []fizz.OperationOption {
	var option []fizz.OperationOption
	option = append(option, fizz.Summary(summary)) // append summary
	if useSecurity {
		option = append(option, fizz.Security(&openapi.SecurityRequirement{
			"bearerAuth": []string{},
		}))
	}
	return option
}
