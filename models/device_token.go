package models

type TrackerToken struct {
	Token string `json:"token" validate:"required"`
}

// Initializes new tracker token instance.
//
//	@param token
//	@return *TrackerToken
func NewTrackerToken(token string) *TrackerToken {
	d := new(TrackerToken)
	d.Token = token
	return d
}
