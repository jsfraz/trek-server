package models

type GNSSDataSummary struct {
	Data     []GNSSData `json:"data" validate:"required"`
	MinSpeed float64    `json:"minSpeed" validate:"required"`
	AvgSpeed float64    `json:"avgSpeed" validate:"required"`
	MaxSpeed float64    `json:"maxSpeed" validate:"required"`
}

// Initializes new GNSS data summary instance.
//
//	@param data
//	@param min
//	@param avg
//	@param max
//	@return *GNSSDataSummary
func NewGNSSDataSummary(data []GNSSData, min *float64, avg *float64, max *float64) *GNSSDataSummary {
	g := new(GNSSDataSummary)
	g.Data = data
	if min == nil {
		g.MinSpeed = 0
	} else {
		g.MinSpeed = *min
	}
	if avg == nil {
		g.AvgSpeed = 0
	} else {
		g.AvgSpeed = *avg
	}
	if max == nil {
		g.MaxSpeed = 0
	} else {
		g.MaxSpeed = *max
	}
	return g
}
