package booking_system

type Seat struct {
	isBooked   bool
	seatNumber int
}

var seating map[int]Seat
