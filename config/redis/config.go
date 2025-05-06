package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

var (
	redisHost     = "REDIS_HOST"
	redisPort     = "REDIS_PORT"
	redisPassword = "REDIS_PASSWORD"
	redisDB       = "REDIS_DB"
	conn          *redis.Client
)

func InitRedis() error {
	host := os.Getenv(redisHost)
	port := os.Getenv(redisPort)
	password := os.Getenv(redisPassword)
	db := os.Getenv(redisDB)

	dbInt, err := strconv.Atoi(db)
	if err != nil {
		return err
	}
	conn = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       dbInt,
	})

	ctx := context.Background()
	_, err = conn.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func GetRedis() *redis.Client {
	return conn
}
