package handlers

import (
	"encoding/base64"
	"errors"
	"jsfraz/trek-server/database"
	"jsfraz/trek-server/models"
	"jsfraz/trek-server/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User login.
//
//	@param c
//	@param login
//	@return *models.LoginResponse
//	@return error
func Login(c *gin.Context, login *models.Login) (*models.LoginResponse, error) {
	// search user in database
	user, err := database.GetUserByUsername(login.Username)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	// user not found
	if user.Username == "" {
		c.AbortWithStatus(401)
		return nil, errors.New("user not found")
	}
	// password check
	hashBytes, _ := base64.StdEncoding.DecodeString(user.PasswordHash)
	err = bcrypt.CompareHashAndPassword(hashBytes, []byte(login.Password))
	if err != nil {
		// incorrect password
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.AbortWithStatus(401)
			return nil, err
		} else {
			// internal error
			c.AbortWithStatus(500)
			return nil, err
		}
	}
	// token
	accessToken, err := utils.GenerateToken(user.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return models.NewLoginResponse(accessToken), nil
}
