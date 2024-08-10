package main

import (
	"fmt"
	"jsfraz/trek-server/database"
	"jsfraz/trek-server/models"
	"jsfraz/trek-server/routes"
	"jsfraz/trek-server/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// log settings
	log.SetPrefix("trek-server: ")
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds)

	// Load config
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	singleton := utils.GetSingleton()
	singleton.Config = config

	// Postgres database
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		singleton.Config.PostgresUser,
		singleton.Config.PostgresPassword,
		singleton.Config.PostgresServer,
		singleton.Config.PostgresPort,
		singleton.Config.PostgresDb,
	)
	postgres, err := gorm.Open(postgres.Open(connStr), &gorm.Config{Logger: logger.Default.LogMode(singleton.Config.GetGormLogLevel())})
	if err != nil {
		log.Fatal(err)
	}
	// database schema migration
	err = postgres.AutoMigrate(
		&models.User{},
		&models.Tracker{},
		&models.GNSSData{},
	)
	if err != nil {
		log.Fatal(err)
	}
	// set database client in singleton
	singleton.PostgresDb = *postgres
	// check if superuser exists
	exists, _ := database.UserExistsByUsername(singleton.Config.SuperuserUsername)
	if !exists {
		// create superuser
		u, _ := models.NewUser(singleton.Config.SuperuserUsername, singleton.Config.SuperuserPassword, true)
		err = database.CreateSuperuser(*u)
		if err != nil {
			log.Fatal(err)
		}
	}

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
