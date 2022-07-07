package ratelimit

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
	"github.com/labstack/echo/v4/middleware"
)

type RedisRateLimiter struct {
	Limit   redis_rate.Limit
	limiter *redis_rate.Limiter
}

var _ middleware.RateLimiterStore = &RedisRateLimiter{}

func (rr *RedisRateLimiter) Allow(identifier string) (bool, error) {
	ctx := context.Background()
	res, err := rr.limiter.Allow(ctx, identifier, rr.Limit)
	if err != nil {
		return false, err
	}

	return res.Allowed > 0, nil
}

func NewRedisRateLimiter(rdb *redis.Client, limit redis_rate.Limit) *RedisRateLimiter {
	return &RedisRateLimiter{
		limiter: redis_rate.NewLimiter(rdb),
		Limit:   limit,
	}
}
