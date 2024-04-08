package entities

import (
	"fmt"

	"github.com/google/uuid"
)

// ParkingLot represents a parking lot.
type ParkingLot struct {
	ID            uuid.UUID `gorm:"primary_key"`
	Name          string    `gorm:"unique"`
	Capacity      int
	ParkingSpaces []ParkingSpace
	// Below are the stats we require for monitoring, they are not persisted.
	Available int `gorm:"-"`
	Occupied  int `gorm:"-"`
	Reserved  int `gorm:"-"`
	Blocked   int `gorm:"-"`
}

// OnCreate sets internally managed fields, name and capcaity and creates parking spaces and sets initial statistics.
func (s *ParkingLot) OnCreate(name string, capacity int) {
	s.ID = uuid.New()
	s.Name = name
	s.Capacity = capacity

	s.ParkingSpaces = make([]ParkingSpace, 0, capacity)
	for i := 0; i < capacity; i++ {
		space := ParkingSpace{}
		space.OnCreate(fmt.Sprintf("%s-%d", name, i+1), s.ID)
		s.ParkingSpaces = append(s.ParkingSpaces, space)
	}

	// On creation all spaces are available.
	s.Available = capacity
	s.Occupied = 0
	s.Reserved = 0
	s.Blocked = 0
}

// OnGet calculates statistics for this parking lot.
func (s *ParkingLot) OnGet() {
	// Reset.
	s.Available = 0
	s.Occupied = 0
	s.Reserved = 0
	s.Blocked = 0

	for _, space := range s.ParkingSpaces {
		switch space.Status {
		case StatusAvailable:
			s.Available++
		case StatusOccupied:
			s.Occupied++
		case StatusReserved:
			s.Reserved++
		case StatusBlocked:
			s.Blocked++
		}
	}
}
