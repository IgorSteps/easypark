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

// ResourceAlreadyExistsError represents an error when user already exists.
type ResourceAlreadyExistsError struct {
	baseError
}

// NewResourceAlreadyExistsError creates a new ResourceAlreadyExistsError.
func NewResourceAlreadyExistsError(name string) *ResourceAlreadyExistsError {
	return &ResourceAlreadyExistsError{
		baseError: baseError{msg: fmt.Sprintf("Resource '%s' already exists", name)},
	}
}

// InvalidInputError represents an error for invalid user input.
type InvalidInputError struct {
	baseError
}

// NewInvalidInputError creates a new instance of InvalidInputError.
func NewInvalidInputError(msg string) *InvalidInputError {
	return &InvalidInputError{
		baseError: baseError{msg: msg},
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
