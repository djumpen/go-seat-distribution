package seat

import (
	"github.com/djumpen/go-seat-distribution/pkg/errors"
)

// Seat represents seat in aircraft salon
type Seat struct {
	Assigned  bool
	Num       string
	SeatClass SeatClass
}

// IsAssigned returns status of seat Assigned
func (s *Seat) IsAssigned() bool {
	return s.Assigned
}

// Number returns seat number
func (s *Seat) Number() string {
	return s.Num
}

// Assign make seat Assigned
func (s *Seat) Assign() error {
	if s.Assigned == true {
		return errors.ErrAlreadyAssigned

	}
	s.Assigned = true
	return nil
}

// SetNumber sets seat number
func (s *Seat) SetNumber(num string) error {
	if num == "" {
		return errors.ErrInvalidNumber
	}
	s.Num = num
	return nil
}

// SetClass sets seat SeatClass
func (s *Seat) SetClass(class SeatClass) {
	s.SeatClass = class
}

func (s *Seat) Class() SeatClass {
	return s.SeatClass
}

// SeatClass represents seatClass according to it's location
type SeatClass string

// Available seat classes
const (
	Aisle  = SeatClass("aisle")
	Window = SeatClass("window")
	Middle = SeatClass("middle")
)
