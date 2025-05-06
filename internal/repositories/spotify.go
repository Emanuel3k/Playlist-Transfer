package repositories

import (
	"context"
	"fmt"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
	"github.com/redis/go-redis/v9"
	"time"
)

type SpotifyRepository struct {
	rds *redis.Client
}

func NewSpotifyRepository(rds *redis.Client) *SpotifyRepository {
	return &SpotifyRepository{
		rds: rds,
	}
}

func (r *SpotifyRepository) SetState(userID string, scope string) *web.AppError {
	ctx := context.Background()
	_, err := r.rds.Set(ctx, userID, scope, 60*time.Second).Result()
	if err != nil {
		return web.InternalServerError(fmt.Errorf("failed to set state for user %s: %s", userID, err.Error()))
	}
	return nil
}
