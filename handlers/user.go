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
	accessToken, err := utils.GenerateAccessToken(user.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return models.NewLoginResponse(accessToken), nil
}

// Handler for creating a new user in the database. The following items are checked:
// 1) To create a user, the ID of the user who wants to make a request must be the same as the superuser ID (if not, return status 401).
// 2) The username must not be taken (if not, return status 409).
//
//	@param c
//	@param register
//	@return error
func CreateUser(c *gin.Context, register *models.CreateUser) error {
	// check for superuser
	u, _ := c.Get("user")
	user := u.(*models.User)
	if !user.Superuser {
		c.AbortWithStatus(401)
		return errors.New("user is not superuser")
	}
	// check if username is taken
	taken, err := database.UserExistsByUsername(register.Username)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	if taken {
		c.AbortWithStatus(409)
		return errors.New("user with given username already exists")
	}
	// initialize user
	newUser, err := models.NewUser(register.Username, register.Password, false)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	// insert
	err = database.InsertUser(*newUser)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Returns the user profile by user ID. The ID is obtained from the access token.
//
//	@param c
//	@return *models.User
//	@return error
func WhoAmI(c *gin.Context) (*models.User, error) {
	u, _ := c.Get("user")
	if u != nil {
		user := u.(*models.User)
		return user, nil
	} else {
		c.AbortWithStatus(500)
		return nil, errors.New("no user set in context")
	}
}

// Get all users.
//
//	@param c
//	@return *[]models.User
//	@return error
func GetAllUsers(c *gin.Context) (*[]models.User, error) {
	// check for superuser
	u, _ := c.Get("user")
	user := u.(*models.User)
	if !user.Superuser {
		c.AbortWithStatus(401)
		return nil, errors.New("user is not superuser")
	}
	// get users
	users, err := database.GetAllUsers()
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return users, nil
}

// Delete user.
//
//	@param c
//	@param id
//	@return error
func DeleteUser(c *gin.Context, id *models.Id) error {
	// check for superuser
	u, _ := c.Get("user")
	user := u.(*models.User)
	if !user.Superuser {
		c.AbortWithStatus(401)
		return errors.New("user is not superuser")
	}
	// check if id belong to root
	if id.Id == user.Id {
		c.AbortWithStatus(500)
		return errors.New("can not delete superuser")
	}
	// delete user
	err := database.DeleteUser(id.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Update user.
//
//	@param c
//	@param request
//	@return error
func UpdateUser(c *gin.Context, request *models.UpdateUser) error {
	// check fo superuser
	u, _ := c.Get("user")
	user := u.(*models.User)
	if !user.Superuser {
		c.AbortWithStatus(401)
		return errors.New("user is not superuser")
	}
	// check if exists
	exists, err := database.UserExistsById(request.Id)
	// error
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	// not found
	if !exists {
		c.AbortWithStatus(500)
		return nil
	}
	// update
	err = database.UpdateUser(request.Id, request.Username, request.Password)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return err
}
