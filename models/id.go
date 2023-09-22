package models

type Id struct {
	Id uint64 `json:"id" validate:"required"`
}
