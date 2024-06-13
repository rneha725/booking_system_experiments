package service

import (
	"booking_system/config"
	"booking_system/mySql"
	"booking_system/redis_mod"
	"fmt"
	"math/rand"
	"time"
)

/*
for 1 user: start a go routine
todo change the use case for locking
- if seat is booked or locked, return that we cannot book: failure
- start booking by updating the isLocked value
- redis_mod: isLocked
- setNx for atomic key set, we can set a TTL in redis_mod of payment service timeout
- eval to run scripts on redis_mod
*/

func StartBooking(userId int, seatNumber int) (bool, error) {
	timeOutChannel := time.After(config.ExpiryDuration)
	paymentChannel := make(chan bool)

	checkBooking(seatNumber)
	if err := redis_mod.SetSeat(seatNumber, userId); err != nil {
		return false, fmt.Errorf("cannot set the seat for seat and user: %d, %d", seatNumber, userId)
	}

	go func() {
		fmt.Printf("starting payment for seat and user: %d, %d\n", seatNumber, userId)
		paymentDuration := time.Duration(rand.Intn(11)) * time.Second
		fmt.Println("Payment took: {}", paymentDuration)
		time.Sleep(paymentDuration)
		paymentChannel <- true
	}()

	select {
	case <-paymentChannel:
		if err := mySql.SetBooking(userId, seatNumber); err != nil {
			fmt.Printf("Cannot book the seat for user: %d, %d %v. Unblocking the seat\n", seatNumber, userId, err)
			redis_mod.RemoveSeat(seatNumber, userId)
		} else {
			fmt.Printf("seat booked for user: %d, %d\n", seatNumber, userId)
		}
	case <-timeOutChannel:
		fmt.Printf("timeout: seat cannot be booked, user, seat number: %d, %d\n", userId, seatNumber)
	}

	return true, nil
}

func checkBooking(seatNumber int) {
	booking, err := mySql.GetBooking(seatNumber)

	if err != nil {
		panic(err)
	}

	if booking.UserId != 0 {
		panic(fmt.Errorf("already booked")) //do not panic, let the other goroutines run
	}
}

func RemoveBooking(userId int, seatNumber int) {
	//pass
}
