package handlers

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

// Socket connect event.
//
//	@param s
//	@return error
func SocketConnect(s socketio.Conn) error {
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
	// check if user exists
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
}

// Send current event.
//
//	@param s
//	@param msg
//	@return error
func SendCurrentEvent(s socketio.Conn, msg map[string]interface{}) error {
	tracker := s.Context()
	t := tracker.(*models.Tracker)
	// parse map to struct
	dataInput, err := models.ParseMap(msg)
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
	existsData, err := database.GNSSDataExists(t.Id, dataInput.Timestamp)
	if err != nil {
		return err
	}
	if existsData {
		return errors.New("duplicate GNSS data")
	}
	// insert to database
	gnssData, err := dataInput.ToDatabaseModel(t.Id)
	if err != nil {
		return err
	}
	return database.InsertGNSSData(*gnssData)
}

// Socket disconnect event.
//
//	@param s
//	@param reason
func SocketDisconnectEvent(s socketio.Conn, reason string) {
	tracker := s.Context()
	if tracker == nil {
		tracker = models.DefaultTracker()
	}
	log.Println("Tracker '" + tracker.(*models.Tracker).Name + "' disconnected (" + reason + ").")
}
