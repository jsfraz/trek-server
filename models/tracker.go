package models

type Tracker struct {
	Id       uint64     `json:"id" validate:"required" gorm:"primarykey"`
	Name     string     `json:"name" validate:"required"`
	GNSSData []GNSSData `json:"-"`
}

// Initialize new tracker instance.
//
//	@param name
//	@return *Tracker
func NewTracker(name string) *Tracker {
	s := new(Tracker)
	s.Name = name
	return s
}

// Default tracker .
//
//	@return *Tracker
func DefaultTracker() *Tracker {
	s := new(Tracker)
	s.Name = "unknown"
	return s
}
