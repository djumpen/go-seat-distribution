package salon

import (
	"github.com/djumpen/go-seat-distribution/pkg/errors"
	"github.com/djumpen/go-seat-distribution/pkg/seat"
)

var DefaultOrder = []seat.SeatClass{seat.Aisle, seat.Window, seat.Middle}

// DefaultSeatAssigning assign seats according to classes order
func DefaultSeatAssign(salon *Salon) (seat.Seat, error) {
	for _, class := range DefaultOrder {
		s, err := assignSeatByClass(salon, class)
		if err == nil {
			return s, nil
		}
	}
	return seat.Seat{}, errors.ErrAllSeatsAssigned
}

func assignSeatByClass(salon *Salon, class seat.SeatClass) (seat.Seat, error) {
	for rowIndex, row := range salon.Rows {
		for blockIndex, block := range row.SeatBlocks {
			for seatIndex, s := range block.Seats {
				if !s.IsAssigned() && s.Class() == class {
					err := salon.Rows[rowIndex].SeatBlocks[blockIndex].Seats[seatIndex].Assign()
					return salon.Rows[rowIndex].SeatBlocks[blockIndex].Seats[seatIndex], err
				}
			}
		}
	}
	return seat.Seat{}, errors.ErrAllSeatsAssigned
}
