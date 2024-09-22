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
func NewGNSSDataSummary(data []GNSSData, min float64, avg float64, max float64) *GNSSDataSummary {
	g := new(GNSSDataSummary)
	g.Data = data
	g.MinSpeed = min
	g.AvgSpeed = avg
	g.MaxSpeed = max
	return g
}
