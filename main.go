package main

import (
	"jsfraz/trek-server/database"
	"jsfraz/trek-server/models"
	"jsfraz/trek-server/routes"
	"jsfraz/trek-server/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// log settings
	log.SetPrefix("trek-server: ")
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds)

	waitDelay := 1
	log.Printf("Waiting %d seconds for the Postgres server...", waitDelay)
	time.Sleep(time.Second * time.Duration(waitDelay))

	// check Gin envs
	utils.CheckGinModeEnv()
	// check Postgres envs
	utils.CheckPostgresEnvs()
	// Postgres database
	connStr := "postgresql://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_SERVER") + ":" + os.Getenv("POSTGRES_PORT") + "/" + os.Getenv("POSTGRES_DB")
	postgres, err := gorm.Open(postgres.Open(connStr), &gorm.Config{Logger: logger.Default.LogMode(utils.GetGormLogLevel())})
	if err != nil {
		log.Fatal(err)
	}
	// database schema migration
	err = postgres.AutoMigrate(&models.User{}, &models.Tracker{}, &models.GNSSData{})
	if err != nil {
		log.Fatal(err)
	}
	// set database client in singleton
	singleton := utils.GetSingleton()
	singleton.PostgresDb = *postgres
	// superuser envs
	utils.CheckSuperuserEnvs()
	// check if superuser exists
	exists, _ := database.UserExistsByUsername(os.Getenv("SUPERUSER_USERNAME"))
	if !exists {
		// create superuser
		u, _ := models.NewUser(os.Getenv("SUPERUSER_USERNAME"), os.Getenv("SUPERUSER_PASSWORD"), true)
		err = database.CreateSuperuser(*u)
		if err != nil {
			log.Fatal(err)
		}
	}
	// access token envs
	utils.CheckTokenEnvs()

	// get Socket.IO instance
	socketio := routes.NewSocketIOServer()
	// get router
	router, err := routes.NewRouter()
	// log error if not nil
	if err != nil {
		log.Fatal(err)
	}
	// start Socket.IO
	go func() {
		if err := socketio.Serve(); err != nil {
			log.Fatalf("Socket.IO listen error: %s\n", err)
		}
	}()
	defer socketio.Close()
	router.GET("/socket.io/*any", nil, gin.WrapH(socketio))
	router.POST("/socket.io/*any", nil, gin.WrapH(socketio))
	// start HTTP server
	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	log.Println("Started server.")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Fizz listen error: %s\n", err)
	}
}
