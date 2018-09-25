package salon

import "github.com/djumpen/go-seat-distribution/pkg/seat"

// DefaultSeatClassifier classifies seats in default way
func DefaultSeatClassifier(salon *Salon) {
	for rowIndex, row := range salon.Rows {
		for blockIndex, block := range row.SeatBlocks {
			for seatIndex := range block.Seats {
				salon.Rows[rowIndex].SeatBlocks[blockIndex].Seats[seatIndex].SetClass(getSeatClass(blockIndex, seatIndex, len(row.SeatBlocks), len(block.Seats)))
			}
		}
	}
}

// getSeatClass returns seat class for given seat location
func getSeatClass(blockIndex, seatIndex, blocksCount, seatsCount int) seat.SeatClass {
	if isWindowSeat(blockIndex, seatIndex, blocksCount, seatsCount) {
		return seat.Window
	}
	if isAisleSeat(blockIndex, seatIndex, blocksCount, seatsCount) {
		return seat.Aisle
	}
	return seat.Middle
}

// isWindowSeat checks is seat with given location placed near window
func isWindowSeat(blockIndex, seatIndex, blocksCount, seatsCount int) bool {
	if isFirstBlock(blockIndex) && isFirstSeatInBlock(seatIndex) {
		return true
	}
	if isLastBlock(blockIndex, blocksCount) && isLastSeatInBlock(seatIndex, seatsCount) {
		return true
	}
	return false
}

// isAisleSeat checks is seat with given location placed near aisle
func isAisleSeat(blockIndex, seatIndex, blocksCount, seatsCount int) bool {
	if isFirstBlock(blockIndex) && isLastSeatInBlock(seatIndex, seatsCount) {
		return true
	}
	if isLastBlock(blockIndex, blocksCount) && isFirstSeatInBlock(seatIndex) {
		return true
	}
	if !isSideBlock(blockIndex, blocksCount) && (isFirstSeatInBlock(seatIndex) || isLastSeatInBlock(seatIndex, seatsCount)) {
		return true
	}
	return false
}

// isFirstBlock checks is given block is first in a row
func isFirstBlock(blockIndex int) bool {
	return blockIndex == 0
}

// isLastBlock checks is given block is first in a row
func isLastBlock(blockIndex, blocksCount int) bool {
	return blockIndex == blocksCount-1
}

// isSideBlock checks is given block is side
func isSideBlock(blockIndex, blocksCount int) bool {
	return isLastBlock(blockIndex, blocksCount) || isFirstBlock(blockIndex)
}

// isFirstSeatInBlock checks is given seat is first in block
func isFirstSeatInBlock(seatIndex int) bool {
	return seatIndex == 0
}

// isLastSeatInBlock checks is given seat is last in block
func isLastSeatInBlock(seatIndex, seatsCount int) bool {
	return seatIndex == seatsCount-1
}
