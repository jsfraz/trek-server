package models

type UpdateTrackerName struct {
	Id   uint64 `query:"id" validate:"required"`
	Name string `query:"name" min:"2" max:"32" validate:"required"`
}
