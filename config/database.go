package config

import (
	"fmt"
	"github.com/emanuel3k/playlist-transfer/config/postgres"
	"github.com/emanuel3k/playlist-transfer/config/redis"
)

func InitDB() error {
	if err := postgres.Config(); err != nil {
		return fmt.Errorf("error initializing postgres database: %w", err)
	}

	if err := redis.InitRedis(); err != nil {
		return fmt.Errorf("error initializing redis database: %w", err)
	}

	return nil
}
