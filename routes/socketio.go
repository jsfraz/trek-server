package routes

import (
	"jsfraz/trek-server/models"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

// Returns new Socket.IO server.
//
//	@return *socketio.Server
func NewSocketIOServer() *socketio.Server {
	server := socketio.NewServer(nil)

	// TODO comments

	server.OnConnect("/", func(s socketio.Conn) error {
		// TODO get API key from query, check key, get device from database
		device := "test"
		s.SetContext(device)
		log.Println("Device " + device + " connected.")
		return nil
	})

	server.OnEvent("/", "sendCurrent", func(s socketio.Conn, msg map[string]interface{}) {
		// TODO upload data to database
		_, err := models.ParseMap(msg)
		if err != nil {
			panic(err)
		}
	})

	server.OnError("/", func(s socketio.Conn, err error) {
		device := s.Context()
		log.Println("Device " + device.(string) + " error (" + err.Error() + ").")
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		device := s.Context()
		log.Println("Device " + device.(string) + " disconnected (" + reason + ").")
	})

	return server
}
