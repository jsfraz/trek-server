package models

type Tracker struct {
	Id   uint64 `json:"id" validate:"required" gorm:"primarykey"`
	Name string `json:"name" validate:"required"`
}

// Initialize new tracker.
//
//	@param name
//	@return *Device
func NewTracker(name string) *Tracker {
	s := new(Tracker)
	s.Name = name
	return s
}
