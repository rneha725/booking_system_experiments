package main

import (
	"booking_system/config"
	"booking_system/service"
)

func main() {
	config.InitMySql()
	//todo call with the input value
	//mySql.SetupMysqlData()

	config.InitRedis()

	service.StartBooking(1, 1)
}
