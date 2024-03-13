package models

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
