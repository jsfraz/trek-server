package models

type UpdateTrackerName struct {
	Id   uint64 `query:"id" validate:"required"`
	Name string `query:"name" validate:"required,min=2,max=32"`
}
