package enitities

// UserRole defines the type for user roles within EasyPark.
type UserRole string

const (
	Admin  UserRole = "admin"
	Driver UserRole = "driver"
)
