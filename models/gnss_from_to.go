package models

import (
	"errors"
	"jsfraz/trek-server/utils"
	"time"
)

type GNSSFromTo struct {
	Id      uint64 `query:"id" validate:"required"`
	FromUtc string `query:"fromUtc" validate:"required"`
	ToUtc   string `query:"toUtc" validate:"required"`
	Offset  int    `query:"offset" validate:"min=1,max=3600,required"`
}

// Validate timestamps.
//
//	@receiver g
//	@return error
func (g GNSSFromTo) ValidateTimestamps() (*time.Time, *time.Time, error) {
	// timestamp validation
	from, err := utils.ParseISO8601String(g.FromUtc)
	if err != nil {
		return nil, nil, errors.New("invalid ISO 8601 timestamp: '" + g.FromUtc + "'")
	}
	to, err := utils.ParseISO8601String(g.ToUtc)
	if err != nil {
		return nil, nil, errors.New("invalid ISO 8601 timestamp: '" + g.ToUtc + "'")
	}
	// check if from is after to
	toAfterFrom := to.After(*from)
	if !toAfterFrom {
		return nil, nil, errors.New("'to' is supposed to be after 'from'")
	}
	return from, to, nil
}
