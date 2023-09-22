package models

type GNSSData struct {
	Id        uint64 `json:"id" validate:"required" gorm:"primarykey"`
	TrackerId uint64 `json:"trackerId" validate:"required"`
	GNSSDataInput
}
