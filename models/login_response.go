package models

type LoginResponse struct {
	AccessToken string `json:"accessToken" validate:"required"`
}

// Initialize new login response.
//
//	@param accessToken
//	@return *LoginResponse
func NewLoginResponse(accessToken string) *LoginResponse {
	l := new(LoginResponse)
	l.AccessToken = accessToken
	return l
}
