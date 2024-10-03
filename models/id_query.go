package models

type Id struct {
	// TODO json instead of query
	Id uint64 `query:"id" validate:"required"`
}
