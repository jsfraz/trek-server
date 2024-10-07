package models

type CreateUser struct {
	// TODO json instead of query
	Username string `query:"username" validate:"required,alphanum,min=2,max=32"`
	Password string `query:"password" validate:"required,min=8,max=64"`
}
