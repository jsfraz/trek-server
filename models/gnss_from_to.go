package models

import (
	"errors"
	"jsfraz/trek-server/utils"
	"time"
)

type GNSSFromTo struct {
	Id     uint64 `query:"id" validate:"required"`
	From   string `query:"from" validate:"required"`
	To     string `query:"to" validate:"required"`
	Offset int    `query:"offset" validate:"min=1,max=300"`
}

// Validate timestamps.
//
//	@receiver g
//	@return error
func (g GNSSFromTo) ValidateTimestamps() (*time.Time, *time.Time, error) {
	// timestamp validation
	from, err := utils.ParseISO8601String(g.From)
	if err != nil {
		return nil, nil, errors.New("Invalid ISO 8601 timestamp: '" + g.From + "'")
	}
	to, err := utils.ParseISO8601String(g.To)
	if err != nil {
		return nil, nil, errors.New("Invalid ISO 8601 timestamp: '" + g.To + "'")
	}
	// check if from is after to
	toAfterFrom := to.After(*from)
	if !toAfterFrom {
		return nil, nil, errors.New("'to' is supposed to be after 'from'")
	}
	return from, to, nil
}
