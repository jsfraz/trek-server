package utils

// https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
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
	// payload
	now := time.Now()
	claims := jwt.MapClaims{}
	claims["sub"] = id
	claims["type"] = "user"
	claims["exp"] = now.Add(time.Second * time.Duration(GetSingleton().Config.AccessTokenLifespan)).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// create and sign token
	return token.SignedString([]byte(GetSingleton().Config.AccessTokenSecret))
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
	claims["type"] = "tracker"
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// create and sign token
	return token.SignedString([]byte(GetSingleton().Config.TrackerTokenSecret))
}

// Return token claim "type".
//
//	@param token
//	@return string
//	@return error
func GetTokenType(token string) (string, error) {
	// Get parts
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid token format")
	}
	// Get payload
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}
	// Get claims
	var claims map[string]interface{}
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return "", err
	}
	// Check type
	if claims["type"] == nil {
		return "", errors.New("token type not set")
	}
	return claims["type"].(string), nil
}
