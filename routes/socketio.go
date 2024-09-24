package routes

import (
	"jsfraz/trek-server/handlers"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

// Returns new Socket.IO server.
//
//	@return *socketio.Server
func NewSocketIOServer() *socketio.Server {
	server := socketio.NewServer(nil)

	// connect
	server.OnConnect("/", handlers.SocketConnect)
	// send GNSS data event
	server.OnEvent("/", "sendCurrent", handlers.SendCurrentEvent)
	// Get current GNSS data event
	server.OnEvent("/", "getCurrent", handlers.GetCurrentEvent)
	// error
	server.OnError("/", func(s socketio.Conn, err error) {
		log.Printf("Error: %s",
			err.Error(),
		)
	})
	// disconnect
	server.OnDisconnect("/", handlers.SocketDisconnectEvent)

	return server
}
