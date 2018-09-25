package salon

import (
	"github.com/djumpen/go-seat-distribution/pkg/errors"
	"github.com/djumpen/go-seat-distribution/pkg/seat"
)

type SalonFactory interface {
	NewSalon(rowsCount, blocksCount int, seatCountPerBlock []int) (*Salon, error)
}

type DefaultSalonFactory struct {
	classifierFunc SeatClassifierFunc
	namingFunc     SeatNamingFunc
}

func NewDefaultSalonFactory(classifierFunc SeatClassifierFunc, namingFunc SeatNamingFunc) *DefaultSalonFactory {
	return &DefaultSalonFactory{
		classifierFunc: classifierFunc,
		namingFunc:     namingFunc,
	}
}

func (f *DefaultSalonFactory) NewSalon(rowsCount, blocksCount int, seatCountPerBlock []int) (*Salon, error) {
	if len(seatCountPerBlock) < blocksCount {
		return nil, errors.ErrNotEnoughSeatCount
	}
	rows := make([]SeatRow, rowsCount)
	for i := range rows {
		rows[i] = newSeatRow(blocksCount, seatCountPerBlock)
	}
	salon := &Salon{
		Rows: rows[:],
	}
	f.classifierFunc(salon)
	f.namingFunc(salon)
	return salon, nil
}

func newSeatRow(blocksCount int, seatCountPerBlock []int) SeatRow {
	blocks := make([]SeatBlock, blocksCount)
	for i := range blocks {
		blocks[i] = newSeatBlock(seatCountPerBlock[i])
	}
	return SeatRow{
		SeatBlocks: blocks[:],
	}
}

func newSeatBlock(seatCount int) SeatBlock {
	seats := make([]seat.Seat, seatCount)
	return SeatBlock{
		Seats: seats,
	}
}
