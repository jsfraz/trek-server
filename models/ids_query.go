package models

type Ids struct {
	Ids []uint64 `json:"ids" validate:"required"`
}
