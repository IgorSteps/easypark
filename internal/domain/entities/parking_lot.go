package entities

import (
	"fmt"

	"github.com/google/uuid"
)

type ParkingLot struct {
	ID            uuid.UUID `gorm:"primary_key"`
	Name          string
	Location      string // We can make it a more complex type with latitude and longitude, but I don't think it matters for this prototype
	Capacity      int
	ParkingSpaces []ParkingSpace
	// Below are a;; the stats we require for monitoring, they are not persisted
	Available int `gorm:"-"`
	Occupied  int `gorm:"-"`
	Reserved  int `gorm:"-"`
	Blocked   int `gorm:"-"`
}

func (s *ParkingLot) OnCreate(name, location string, capacity int) {
	s.ID = uuid.New()
	s.Name = name
	s.Location = location
	s.Capacity = capacity

	s.ParkingSpaces = make([]ParkingSpace, 0, capacity)

	for i := 0; i < capacity; i++ {
		space := ParkingSpace{}
		space.OnCreate(fmt.Sprintf("%s-%d", name, i+1), s.ID)
		s.ParkingSpaces = append(s.ParkingSpaces, space)
	}
}
