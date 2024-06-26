package utils

// https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Generate and sign access token.
//
//	@param id
//	@return string
//	@return error
func GenerateAccessToken(id uint64) (string, error) {
	// token lifespan
	lifespan, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFESPAN"))
	// payload
	now := time.Now()
	claims := jwt.MapClaims{}
	claims["sub"] = id
	claims["exp"] = now.Add(time.Second * time.Duration(lifespan)).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// create and sign token
	return token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
}

// Check if the token is valid.
//
//	@param tokenStr
//	@param secret
//	@return uint64
//	@return error
func TokenValid(tokenStr string, secret string) (uint64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}
	// user ID
	claims, _ := token.Claims.(jwt.MapClaims)
	fId := claims["sub"].(float64)
	return uint64(fId), nil
}

// Extract token from Gin context.
//
//	@param c
//	@return string
func ExtractTokenFromContext(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// Generate and sign tracker token.
//
//	@param id
//	@return string
//	@return error
func GenerateTrackerToken(id uint64) (string, error) {
	// payload
	now := time.Now()
	claims := jwt.MapClaims{}
	claims["sub"] = id
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// create and sign token
	return token.SignedString([]byte(os.Getenv("TRACKER_TOKEN_SECRET")))
}
