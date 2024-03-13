package models

import "github.com/IgorSteps/easypark/internal/domain/entities"

// UserCreationRequest represents the expected data for creating a new user.
type UserCreationRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func (s *UserCreationRequest) ToDomain() *entities.User {
	return &entities.User{
		FirstName: s.Firstname,
		LastName:  s.Lastname,
		Username:  s.Username,
		Password:  s.Password,
		Email:     s.Email,
	}
}
