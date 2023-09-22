package utils

import (
	"time"
)

const layout = "2006-01-02T15:04:05.999999"

// Parse ISO8601 string.
//
//	@param timestamp
//	@return *time.Time
//	@return error
func ParseISO8601String(timestamp string) (*time.Time, error) {
	t, err := time.Parse(layout, timestamp)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
