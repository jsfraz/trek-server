package utils

import (
	"log"
	"os"
	"regexp"

	"gorm.io/gorm/logger"
)

// Regex number pattern
const numberPattern string = `^\d+$`

// Checks environment variables for PostgreSQL connections.
// If an incorrect value is set, the program exits.
func CheckPostgresEnvs() {
	// env variables
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	server := os.Getenv("POSTGRES_SERVER")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_SERVER")

	// check
	ok := true
	// username
	if user == "" {
		ok = false
		log.Println("Empty username for PostgreSQL.")
	}
	// password
	if password == "" {
		ok = false
		log.Println("Empty PostgreSQL password.")
	}
	// mysql server
	if server == "" {
		ok = false
		log.Println("Empty PostgreSQL server address.")
	}
	// port
	matchPort, _ := regexp.MatchString(numberPattern, port)
	if !matchPort {
		ok = false
		log.Println("Empty or invalid PostgreSQL port.")
	}
	// database name
	if db == "" {
		ok = false
		log.Println("Empty PostgreSQL database.")
	}
	// result
	if !ok {
		log.Fatalln("Check the environment variables. Shutting down...")
	}
}

// Checks the environment variables for the superuser.
// If an incorrect value is set, the program exits.
func CheckSuperuserEnvs() {
	// env variables
	user := os.Getenv("SUPERUSER_USERNAME")
	password := os.Getenv("SUPERUSER_PASSWORD")

	// kontrola
	ok := true
	// uživatelské jméno
	if user == "" {
		ok = false
		log.Println("Empty username for superuser.")
	}
	// heslo
	if password == "" {
		ok = false
		log.Println("Blank superuser password.")
	}
	// result
	if !ok {
		log.Fatalln("Check the environment variables. Shutting down...")
	}
}

// Checks the environment variables for the access token.
// If an incorrect value is set, the program exits.
func CheckTokenEnvs() {
	// kontrola
	ok := true
	// access token secret
	if os.Getenv("ACCESS_TOKEN_SECRET") == "" {
		ok = false
		log.Println("Invalid access token secret.")
	}
	// životnost tokenu
	matchAccess, _ := regexp.MatchString(numberPattern, os.Getenv("ACCESS_TOKEN_LIFESPAN"))
	if !matchAccess {
		ok = false
		log.Println("Invalid access token lifetime.")
	}
	// tracker token
	if os.Getenv("TRACKER_TOKEN_SECRET") == "" {
		ok = false
		log.Println("Invalid tracker token secret.")
	}
	// result
	if !ok {
		log.Fatalln("Check the environment variables. Shutting down...")
	}
}

// Checks the environment variable for deployment.
// The default value can be empty.
// If an incorrect value is set, the program exits.
func CheckGinModeEnv() {
	match, _ := regexp.MatchString(`^(|debug|release)$`, os.Getenv("GIN_MODE"))
	if !match {
		log.Println("Invalid Gin mode.")
		log.Fatalln("Check the environment variables. Shutting down...")
	}
}

// Returns the Gorm log level according to the environment variable
//
//	@return logger.LogLevel
func GetGormLogLevel() logger.LogLevel {
	mode := os.Getenv("GIN_MODE")
	if mode == "release" {
		return logger.Error
	}
	return logger.Info
}
