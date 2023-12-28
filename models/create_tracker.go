package models

type CreateTracker struct {
	Name string `query:"name" validate:"required,min=2,max=32"`
}
