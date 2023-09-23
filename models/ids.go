package models

type Ids struct {
	Ids []uint64 `query:"ids" validate:"required"`
}
