package salon

import (
	"strconv"

	"github.com/djumpen/go-seat-distribution/pkg/util"
)

// DefaultSeatNaming names salon seats in default way
func DefaultSeatNaming(salon *Salon) {
	for rowIndex, row := range salon.Rows {
		seatLetter := 0
		for blockIndex, block := range row.SeatBlocks {
			for seatIndex := range block.Seats {
				salon.Rows[rowIndex].SeatBlocks[blockIndex].Seats[seatIndex].SetNumber(strconv.Itoa(rowIndex+1) + util.StringValueOf(seatLetter))
				seatLetter++
			}
		}
	}
}
