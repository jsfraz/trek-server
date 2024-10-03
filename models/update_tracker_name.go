package models

type UpdateTrackerName struct {
	// TODO json instead of query
	Id   uint64 `query:"id" validate:"required"`
	Name string `query:"username" validate:"required,min=2,max=32"`
}
