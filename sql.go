package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // import the MySQL driver
	"log"
)

var db *sql.DB

const seatTableCreateQuery = "CREATE TABLE IF NOT EXISTS SeatTable " +
	"(id INT AUTO_INCREMENT," +
	"seat_number VARCHAR(255)," +
	"PRIMARY KEY(id))"

const seatLockTableCreateQuery = "CREATE TABLE IF NOT EXISTS SeatLock " +
	"(id INT AUTO_INCREMENT," +
	"seat_number VARCHAR(255)," +
	"isLocked BOOL DEFAULT FALSE," +
	"isBooked BOOL DEFAULT FALSE," +
	"PRIMARY KEY(id))"

const seatTableDummyDataQuery = "INSERT INTO SeatTable (seat_number) VALUES (?)"
const seatLockTableDummyDataQuery = "INSERT INTO SeatLock (seat_number) VALUES (?)"
const selectSeat = "SELECT * FROM SeatTable WHERE seat_number = %d LIMIT 1"
const selectSeatLock = "SELECT * FROM SeatLock WHERE seat_number = %d LIMIT 1"

func initDb() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(localhost:3306)/BOOKING_SYSTEM")
	if err != nil {
		log.Fatalf("error in opening sql connection: %v", err)
	}

	db.SetMaxOpenConns(100)
}

func main() {
	initDb()
	//createTables()
	//createDummyDataInTable(seatTableDummyDataQuery)
	//createDummyDataInTable(seatLockTableDummyDataQuery)

	seat, err := getSeat(0)
	if err != nil {
		fmt.Println("Error in getting seat: {}", err)
	}

	fmt.Println(seat.seatNumber)

	seatLock, err := getSeatLock(1)
	if err != nil {
		fmt.Println("Error in getting seatLock: {}", err)
	}

	fmt.Println(seatLock.seatNumber)
}

func createTables() {
	_, err := db.Query(seatTableCreateQuery)
	if err != nil {
		log.Fatalf("error in creating Seat table: %v", err)
	}
	_, err = db.Query(seatLockTableCreateQuery)
	if err != nil {
		log.Fatalf("error in creating SeatLock table: %v", err)
	}
}

func createDummyDataInTable(query string) {
	stmt, err := db.Prepare(query)

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

func getSeat(seatNumber int) (*Seat, error) {
	row, err := db.Query(fmt.Sprintf(selectSeat, seatNumber))

	if err != nil {
		return nil, fmt.Errorf("error getting seat %d, error: %v", seatNumber, err)
	}

	var seat Seat
	var id int
	row.Next()

	// Scan the row into the Seat struct
	if err := row.Scan(&id, &seat.seatNumber); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("seat %d not found", seatNumber)
		}
		// If an error occurred while scanning, return it
		return nil, fmt.Errorf("error scanning row: %v", err)
	}
	return &seat, nil
}

func getSeatLock(seatNumber int) (*SeatLock, error) {
	row, err := db.Query(fmt.Sprintf(selectSeatLock, seatNumber))

	if err != nil {
		return nil, fmt.Errorf("error getting seat %d, error: %v", seatNumber, err)
	}

	var seatLock SeatLock
	var id int
	row.Next()

	// Scan the row into the Seat struct
	if err := row.Scan(&id, &seatLock.seatNumber, &seatLock.isLocked, &seatLock.isBooked); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("seat %d not found", seatNumber)
		}
		// If an error occurred while scanning, return it
		return nil, fmt.Errorf("error scanning row: %v", err)
	}
	return &seatLock, nil
}
