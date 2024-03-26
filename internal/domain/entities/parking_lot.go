package entities

import "github.com/google/uuid"

type ParkingLot struct {
	ID            uuid.UUID `gorm:"primary_key"`
	Name          string
	Location      string // We can make it a more complex type with latitude and longitude, but I don't think it matters for this prototype
	Capacity      int
	Available     int `gorm:"-"`
	Occupied      int `gorm:"-"`
	Reserved      int `gorm:"-"`
	Blocked       int `gorm:"-"`
	ParkingSpaces []ParkingSpace
}
