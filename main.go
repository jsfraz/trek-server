package main

import (
	"jsfraz/trek-server/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Fizz listen error: %s\n", err)
	}
}
