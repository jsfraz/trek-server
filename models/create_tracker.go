package models

type CreateTracker struct {
	Name string `json:"name" min:"2" max:"32" validate:"required,alphanum"`
}
