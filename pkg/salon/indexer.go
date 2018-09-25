package salon

import (
	"github.com/djumpen/go-seat-distribution/pkg/errors"
	"github.com/djumpen/go-seat-distribution/pkg/seat"
)

//GetSeatByIndex returns
func GetSeatByIndex(salon *Salon, index int) (seat.Seat, error) {
	var findIndex int
	for rowIndex, row := range salon.Rows {
		for blockIndex, block := range row.SeatBlocks {
			for seatIndex := range block.Seats {
				if findIndex == index {
					return salon.Rows[rowIndex].SeatBlocks[blockIndex].Seats[seatIndex], nil
				}
				findIndex++
			}
		}
	}
	return seat.Seat{}, errors.ErrSeatIndexOutOfRange
}
