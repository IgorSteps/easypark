package repositories

import "fmt"

// baseError is our base error type.
type baseError struct {
	msg string
}

// Error implements the error interface for BaseError.
func (e *baseError) Error() string {
	return e.msg
}

// UserNotFoundError represents an error for not finding user.
type UserNotFoundError struct {
	baseError
}

// NewUserNotFoundError creates a new UserNotFoundError.
func NewUserNotFoundError(username string) *UserNotFoundError {
	return &UserNotFoundError{
		baseError: baseError{msg: fmt.Sprintf("User '%s' not found", username)},
	}
}

// UserAlreadyExistsError represents an error when user already exists.
type UserAlreadyExistsError struct {
	baseError
}

// NewUserAlreadyExistsError creates a new UserAlreadyExistsError.
func NewUserAlreadyExistsError(username, email string) *UserAlreadyExistsError {
	return &UserAlreadyExistsError{
		baseError: baseError{msg: fmt.Sprintf("User '%s'/'%s' already exists", username, email)},
	}
}

// InvalidCredentialsError represents an error for invalid credentials.
type InvalidCredentialsError struct {
	baseError
}

// NewInvalidCredentialsError creates a new InvalidCredentialsError.
func NewInvalidCredentialsError() *InvalidCredentialsError {
	return &InvalidCredentialsError{
		baseError: baseError{msg: "Invalid password provided"},
	}
}

// InternalError represents an unexpected internal error.
type InternalError struct {
	baseError
}

// NewInternalError creates a new InternalError with a custom message.
func NewInternalError(message string) *InternalError {
	return &InternalError{
		baseError: baseError{msg: fmt.Sprintf("Internal error: %s", message)},
	}
}
