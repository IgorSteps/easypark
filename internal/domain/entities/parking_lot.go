package entities

import "github.com/google/uuid"

type ParkingLot struct {
	ID        uuid.UUID
	Name      string
	Location  string // We can make it a more complex type with latitude and longitude, but I don't think it matters for this prototype
	Capacity  int
	Available int
	Occupied  int
	Reserved  int
	Blocked   int
	Spaces    []ParkingSpace // Association to actual parking spaces
}
