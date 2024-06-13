package redis_mod

import (
	"booking_system/config"
	"context"
	"fmt"
	"sync"
)

var mutex sync.Mutex

func GetSeat(seatNumber int) (string, error) {
	ctx := context.Background()
	val, err := config.RedisClient.Get(ctx, string(rune(seatNumber))).Result()

	if err != nil {
		return "", err
	}

	fmt.Println(seatNumber, val)
	return val, nil
}

func SetSeat(seatNumber int, user int) error {
	ctx := context.Background()
	_, err := config.RedisClient.SetNX(ctx, string(rune(seatNumber)), user, config.ExpiryDuration).Result()

	if err != nil {
		return err
	}

	return nil
}

func RemoveSeat(seatNumber int, userId int) {
	mutex.Lock()
	defer mutex.Unlock()

	ctx := context.Background()
	val, err := config.RedisClient.Get(ctx, string(rune(seatNumber))).Result()

	if err != nil {
		fmt.Println("No ongoing booking for seatNumber: {}", seatNumber)
	}

	if val == string(rune(userId)) {
		config.RedisClient.Del(ctx, string(rune(seatNumber)))
		fmt.Println("Released seatNumber: {}", seatNumber)
	}
}
