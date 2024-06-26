package middlewares

import (
	"jsfraz/trek-server/database"
	"jsfraz/trek-server/utils"
	"os"

	"github.com/gin-gonic/gin"
)

// Middleware for user authentication.
// If the user has a valid access token, it sets its ID in the context.
// If it is not valid, it returns a status of 401.
//
//	@param c Gin context
func Auth(c *gin.Context) {
	// get access token from context
	userId, err := utils.TokenValid(utils.ExtractTokenFromContext(c), os.Getenv("ACCESS_TOKEN_SECRET"))
	// invalid token
	if err != nil {
		c.AbortWithStatus(401)
		c.Error(err)
	}
	// check if user
	exists, err := database.UserExistsById(userId)
	// error
	if err != nil {
		c.AbortWithStatus(500)
		c.Error(err)
	}
	// not found
	if !exists {
		c.AbortWithStatus(401)
	}
	// get user
	user, err := database.GetUserById(userId)
	if err != nil {
		c.AbortWithStatus(500)
		c.Error(err)
	}
	// set user to context, continue
	c.Set("user", user)
	c.Next()
}
