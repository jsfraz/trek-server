package handlers

import (
	"errors"
	"fmt"
	"jsfraz/trek-server/database"
	"jsfraz/trek-server/models"
	"jsfraz/trek-server/utils"
	"log"
	"net/url"

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
	if apiKey == "" {
		return errors.New("empty apiKey query parameter")
	}
	// Get token type
	tokenType, err := utils.GetTokenType(apiKey)
	switch tokenType {

	// Tracker
	case "tracker":
		trackerId, err := utils.TokenValid(apiKey, utils.GetSingleton().Config.TrackerTokenSecret)
		// invalid token
		if err != nil {
			return err
		}
		// check if tracker exists
		exists, err := database.TrackerExistsById(trackerId)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("tracker does not exist")
		}
		// get tracker
		tracker, err := database.GetTrackerById(trackerId)
		if err != nil {
			return err
		}
		// set to context
		s.SetContext(tracker)
		log.Println("Tracker '" + tracker.Name + "' connected.")
		break

	// User
	case "user":
		userId, err := utils.TokenValid(apiKey, utils.GetSingleton().Config.AccessTokenSecret)
		// Invalid token
		if err != nil {
			return err
		}
		// Check if user exists
		exists, err := database.UserExistsById(userId)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("user does not exist")
		}
		// Get user
		user, err := database.GetUserById(userId)
		if err != nil {
			return err
		}
		// Set to context
		s.SetContext(user)
		log.Println("User '" + user.Username + "' connected.")
		break

	// Default
	default:
		return fmt.Errorf("invalid token type '%s'", tokenType)

	}
	return nil
}

// Send current event.
//
//	@param s
//	@param msg
//	@return error
func SendCurrentEvent(s socketio.Conn, msg map[string]interface{}) error {
	t, okTracker := s.Context().(*models.Tracker)
	if !okTracker {
		return errors.New("only trackers are allowed for this event")
	}
	// parse map to struct
	dataInput, err := models.ParseMapToGnssDataInput(msg)
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

// Get current GNSS data.
//
//	@param s
//	@param msg
//	@return error
func GetCurrentEvent(s socketio.Conn, msg map[string]interface{}) error {
	_, okUser := s.Context().(*models.User)
	if !okUser {
		return errors.New("only users are allowed for this event")
	}
	// Parse map to struct
	dataInput, err := models.ParseGetCurrentGNSSDataInput(msg)
	if err != nil {
		log.Println(err)
		return err
	}
	// Get current data
	data, err := database.GetCurrentGNSSData(dataInput.Id)
	if err != nil {
		return err
	}
	s.Emit("getCurrent", data)
	return nil
}

// TODO fmt formatting
// Socket disconnect event.
//
//	@param s
//	@param reason
func SocketDisconnectEvent(s socketio.Conn, reason string) {
	obj := s.Context()
	tracker, okTracker := obj.(*models.Tracker)
	if okTracker {
		log.Println("Tracker '" + tracker.Name + "' disconnected (" + reason + ").")
	}
	user, okUser := obj.(*models.User)
	if okUser {
		log.Println("User '" + user.Username + "' disconnected (" + reason + ").")
	}
}
