package models

type MessageResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type TokenResponse struct {
	Status      string `json:"status"`
	AccessToken string `json:"access_token"`
}

type UserDataResponse struct {
	Status string        `json:"status"`
	Data   UserContainer `json:"data"`
}

type UserMeResponse struct {
	Status string        `json:"status"`
	Data   UserContainer `json:"data"`
}

type UserContainer struct {
	User UserResponse `json:"user"`
}
