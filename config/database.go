package config

import (
	"fmt"
	"github.com/emanuel3k/playlist-transfer/config/postgres"
)

func InitDB() error {
	if _, err := postgres.Config(); err != nil {
		return fmt.Errorf("error initializing postgres database: %w", err)
	}

	return nil
}
