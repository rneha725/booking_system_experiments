package mySql

import (
	"booking_system/config"
	"database/sql"
	"fmt"
	"log"
)

const createBookingsTable = "CREATE TABLE IF NOT EXISTS BookingsTable (seat_number INT PRIMARY KEY, user_id INT DEFAULT 0)"

const bookingsTableInsertQuery = "INSERT INTO BookingsTable (seat_number) VALUES (?)"
const selectSeat = "SELECT * FROM BookingsTable WHERE seat_number = %d LIMIT 1"

func SetupMysqlData() {
	createTables()
	createDummyData(bookingsTableInsertQuery)
	testGetFunctions()
}

func testGetFunctions() {
	booking, err := GetBooking(0)
	if err != nil {
		fmt.Println("Error in getting seat: {}", err)
	}

	fmt.Println(booking.SeatNumber)
}

func createTables() {
	_, err := config.DB.Query(createBookingsTable)
	if err != nil {
		log.Fatalf("error in creating Bookings table: %v", err)
	}
}

func createDummyData(query string) {
	stmt, err := config.DB.Prepare(query)

	if err != nil {
		log.Fatalf("Error in preparing insert statement: %v", err)
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Error in closing statement: {}", err)
		}
	}(stmt)

	for i := 0; i < 1000; i++ {
		_, err := stmt.Exec(i)
		if err != nil {
			log.Fatalf("Error inserting into Seat Table: %v", err)
		}
	}
}
