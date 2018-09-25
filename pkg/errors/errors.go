package errors

import "errors"

// Application errors
var (
	ErrAlreadyAssigned     = errors.New("seat already assigned")
	ErrInvalidNumber       = errors.New("invalid seat number")
	ErrAllSeatsAssigned    = errors.New("all seats are already assigned")
	ErrSeatIndexOutOfRange = errors.New("seat index out of range")
	ErrSalonNotFound       = errors.New("salon not found")
	ErrNotEnoughSeatCount  = errors.New("not enough seats count specified")
)
