package models

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           uint64 `json:"id" validate:"required" gorm:"primarykey"`
	Username     string `json:"username" validate:"required"`
	PasswordHash string `json:"-" validate:"required"`
	Superuser    bool   `json:"superuser" validate:"required"`
}

// Initialize new user.
//
//	@param username
//	@param password
//	@param superuser
//	@return *User
//	@return error
func NewUser(username string, password string, superuser bool) (*User, error) {
	u := new(User)
	u.Username = username
	err := u.SetPassword(password)
	if err != nil {
		return nil, err
	}
	u.Superuser = superuser
	return u, nil
}

// Set password or return error.
//
//	@receiver u
//	@param password
//	@return error
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = base64.StdEncoding.EncodeToString(bytes)
	return nil
}
