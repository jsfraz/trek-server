package models

type UpdateUser struct {
	Id       uint64 `query:"id" validate:"required"`
	Username string `query:"username" min:"2" max:"32" validate:"required,alphanum"`
	Password string `query:"password" min:"8" max:"64" validate:"required"`
}
