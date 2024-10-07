package models

type UpdateUser struct {
	// TODO json instead of query
	Id       uint64 `query:"id" validate:"required"`
	Username string `query:"username" validate:"required,alphanum,min=2,max=32"`
	Password string `query:"password" validate:"required,min=8,max=64"`
}
