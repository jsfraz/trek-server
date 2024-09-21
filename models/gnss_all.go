package models

type GNSSAll struct {
	Id     uint64 `query:"id" validate:"required"`
	Offset int    `query:"offset" validate:"min=1,max=3600,required"`
}
