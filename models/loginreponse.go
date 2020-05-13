package models

//LoginResponse to connect the Jwt to the login
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}