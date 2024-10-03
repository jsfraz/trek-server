package models

type CreateTracker struct {
	// TODO json instead of query
	Name string `query:"username" validate:"required,min=2,max=32"`
}
