package models

type Ids struct {
	// TODO json instead of query
	Ids []uint64 `query:"ids" validate:"required"`
}
