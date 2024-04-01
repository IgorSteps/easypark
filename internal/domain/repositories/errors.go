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

// NotFoundError represents an error for not finding a resource.
type NotFoundError struct {
	baseError
}

// NewNotFoundError creates a new NotFoundError.
func NewNotFoundError(id string) *NotFoundError {
	return &NotFoundError{
		baseError: baseError{msg: fmt.Sprintf("Resource '%s' not found", id)},
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
