package salon

import (
	"github.com/djumpen/go-seat-distribution/pkg/seat"
)

// Salon represents aircraft salon that contains range of seat rows
type Salon struct {
	Rows []SeatRow `json:"Rows"`
}

// SeatRow represents range of seat blocks in row
type SeatRow struct {
	SeatBlocks []SeatBlock `json:"Blocks"`
}

// SeatBlock represents group of seats
type SeatBlock struct {
	Seats []seat.Seat `json:"Seats"`
}

// SeatNamingFunc represents function, that names salon seats according to it's logic
type SeatNamingFunc func(*Salon)

// SeatAssigningFunc represents function, that choose seat to be assigned according to it's logic
type SeatAssigningFunc func(*Salon) (seat.Seat, error)

// SeatClassifierFunc represents function, that classifies salon seats
type SeatClassifierFunc func(*Salon)

// SeatByIndexFunc represents function, that returns seat by index
type SeatByIndexFunc func(*Salon, int) (seat.Seat, error)
