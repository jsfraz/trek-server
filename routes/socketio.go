package routes

import (
	"errors"
	"jsfraz/trek-server/database"
	"jsfraz/trek-server/models"
	"jsfraz/trek-server/utils"
	"log"
	"net/url"
	"os"

	socketio "github.com/googollee/go-socket.io"
)

// Returns new Socket.IO server.
//
//	@return *socketio.Server
func NewSocketIOServer() *socketio.Server {
	server := socketio.NewServer(nil)

	// connect
	server.OnConnect("/", func(s socketio.Conn) error {
		// get access token from context
		queryValues, err := url.ParseQuery(s.URL().RawQuery)
		if err != nil {
			return err
		}
		// extract the apiKey value
		apiKey := queryValues.Get("apiKey")
		trackerId, err := utils.TokenValid(apiKey, os.Getenv("TRACKER_TOKEN_SECRET"))
		// token není platný
		if err != nil {
			return err
		}
		// check if user
		exists, err := database.TrackerExistsById(trackerId)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("invalid token or tracker does not exist")
		}
		// get tracker
		tracker, err := database.GetTrackerById(trackerId)
		if err != nil {
			return err
		}
		// set to context
		s.SetContext(tracker)
		log.Println("Tracker '" + tracker.Name + "' connected.")
		return nil
	})

	// send GNSS data event
	server.OnEvent("/", "sendCurrent", func(s socketio.Conn, msg map[string]interface{}) error {
		tracker := s.Context()
		t := tracker.(*models.Tracker)
		// parse map to struct
		data, err := models.ParseMap(msg)
		if err != nil {
			return err
		}
		// check if tracker existsTracker
		existsTracker, err := database.TrackerExistsById(t.Id)
		if err != nil {
			return err
		}
		if !existsTracker {
			return errors.New("tracker does not exist")
		}
		// check if GNSS record existsData
		existsData, err := database.GNSSDataExists(t.Id, data.Timestamp)
		if err != nil {
			return err
		}
		if existsData {
			return errors.New("duplicate GNSS data")
		}
		// insert to database
		return database.InsertGNSSData(*data.ToDatabaseModel(t.Id))
	})

	// error
	server.OnError("/", func(s socketio.Conn, err error) {
		log.Printf("%s",
			err.Error(),
		)
	})

	// disconnect
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		tracker := s.Context()
		if tracker == nil {
			tracker = models.DefaultTracker()
		}
		log.Println("Tracker '" + tracker.(*models.Tracker).Name + "' disconnected (" + reason + ").")
	})

	return server
}
