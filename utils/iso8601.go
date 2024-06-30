package utils

import (
	"fmt"
	"time"
)

// Parse ISO8601 string.
//
//	@param timestamp
//	@return *time.Time
//	@return error
func ParseISO8601String(timestamp string) (*time.Time, error) {
	// Define a slice of layout strings representing various ISO 8601 formats
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05.999999999Z07:00", // Nanosecond precision
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05",
		"2006-01-02",
	}

	var parsedTime time.Time
	var err error

	// Iterate through each layout, attempting to parse the timestamp
	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, timestamp)
		if err == nil {
			// If parsing succeeds, return a pointer to the parsed time
			return &parsedTime, nil
		}
	}

	// If all parsing attempts fail, return nil and an error message
	return nil, fmt.Errorf("invalid ISO 8601 format: %s", timestamp)
}
