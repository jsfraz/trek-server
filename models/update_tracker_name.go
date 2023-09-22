package models

type UpdateTrackerName struct {
	Id   uint64 `json:"id" validate:"required"`
	Name string `json:"name" min:"2" max:"32" validate:"required"`
}
