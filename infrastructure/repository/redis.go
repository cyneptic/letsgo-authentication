package repositories

import (
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func createRedisConnection() *redis.Client {
	db, _ := strconv.Atoi(os.Getenv("redis_db"))
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_uri"),
		Password: os.Getenv("redis_password"),
		DB:       db,
	})
	return redisClient

}
