package mySql

import (
	"booking_system/config"
	"booking_system/model"
	"database/sql"
	"errors"
	"fmt"
)

func GetBooking(seatNumber int) (*model.Booking, error) {
	row, err := config.DB.Query(fmt.Sprintf(selectSeat, seatNumber))

	if err != nil {
		return nil, fmt.Errorf("error getting seat %d, error: %v", seatNumber, err)
	}

	var booking model.Booking
	row.Next()

	// Scan the row into the Seat struct
	if err := row.Scan(&booking.SeatNumber, &booking.UserId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("seat %d not found", seatNumber)
		}
		// If an error occurred while scanning, return it
		return nil, fmt.Errorf("error scanning row: %v", err)
	}
	return &booking, nil
}

func SetBooking(userId int, seatNumber int) error {
	updateQuery := "UPDATE BookingsTable SET user_id = ? WHERE seat_number = ? AND user_id = 0"
	stmt, err := config.DB.Prepare(updateQuery)

	if err != nil {
		return fmt.Errorf("error in preparing insert statement seat, user: %d, %d, %d", seatNumber, userId, err)
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Error in closing statement: {}", err)
		}
	}(stmt)

	_, err = stmt.Exec(userId, seatNumber) //todo concurrency problems
	if err != nil {
		return fmt.Errorf("error in setting the seat, user: %d, %d, %d", seatNumber, userId, err)
	}

	return nil
}
