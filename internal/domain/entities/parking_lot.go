package entities

import (
	"fmt"

	"github.com/google/uuid"
)

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
}
