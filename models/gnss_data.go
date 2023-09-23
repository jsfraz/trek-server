package models

import "time"

type GNSSData struct {
	Id        uint64    `json:"id" validate:"required" gorm:"primarykey"`
	TrackerId uint64    `json:"trackerId" validate:"required"`
	Latitude  float64   `json:"latitude" validate:"latitude,required"`
	Longitude float64   `json:"longitude" validate:"longitude,required"`
	Speed     float64   `json:"speed" validate:"min=0"`
	Timestamp time.Time `json:"timestamp" validate:"required" gorm:"type:timestamp"`
}
