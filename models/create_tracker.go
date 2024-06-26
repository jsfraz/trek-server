package models

type CreateTracker struct {
	Name string `query:"name" min:"2" max:"32" validate:"required"`
}
