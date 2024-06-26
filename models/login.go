package models

type Login struct {
	// TODO json instead of query
	Username string `query:"username" validate:"required,alphanum"`
	Password string `query:"password" validate:"required"`
}
