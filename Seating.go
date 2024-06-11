package main

type Seat struct {
	seatNumber int
}

type SeatLock struct {
	seatNumber int
	isLocked   bool
	isBooked   bool
}

var seating map[int]Seat
