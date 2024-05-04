package models

import "github.com/IgorSteps/easypark/internal/domain/entities"

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	User  entities.User `json:"user"`
	Token string        `json:"token"`
}
