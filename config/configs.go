package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// mysql
var DB *sql.DB

const mySqlDatasource = "root@tcp(localhost:3306)/BOOKING_SYSTEM"

func InitMySql() {
	var err error
	DB, err = sql.Open("mysql", mySqlDatasource)
	if err != nil {
		log.Fatalf("error in opening mySql connection: %v", err)
	}
	DB.SetMaxOpenConns(100)
}

// redis
const ExpiryDuration = 5 * time.Second

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
