package models

type Login struct {
	Username string `query:"username" validate:"required,alphanum"`
	Password string `query:"password" validate:"required"`
}
