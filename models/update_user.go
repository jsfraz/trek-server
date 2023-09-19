package models

type UpdateUser struct {
	Id       uint64 `json:"id" validate:"required"`
	Username string `json:"username" min:"2" max:"32" validate:"required,alphanum"`
	Password string `json:"password"`
}
