package models

type Id struct {
	Id uint64 `query:"id" validate:"required"`
}
