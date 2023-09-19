package models

type CreateUser struct {
	Username string `json:"username" min:"2" max:"32" validate:"required,alphanum"`
	Password string `json:"password" min:"8" max:"64" validate:"required"`
}
