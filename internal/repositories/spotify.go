package repositories

import (
	"context"
	"errors"
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
		return web.InternalServerError(fmt.Errorf("failed to set state: %s", err.Error()))
	}
	return nil
}

func (r *SpotifyRepository) GetState(userID string) (string, *web.AppError) {
	ctx := context.Background()
	state, err := r.rds.Get(ctx, userID).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", web.NotFoundError("state not found for user")
		}
		return "", web.InternalServerError(fmt.Errorf("failed to get state: %s", err.Error()))
	}
	return state, nil
}

func (r *SpotifyRepository) SetAccessToken(accessToken string, ttl int) *web.AppError {
	ctx := context.Background()
	_, err := r.rds.Set(ctx, accessToken, accessToken, time.Duration(ttl)*time.Second).Result()
	if err != nil {
		return web.InternalServerError(fmt.Errorf("failed to set access token: %s", err.Error()))
	}
	return nil
}
