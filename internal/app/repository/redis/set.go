package redis

import (
	"context"
	"time"
)

func (r *redisRepo) Set(ctx context.Context, key string, value interface{}, duration time.Duration) (err error) {
	return r.redisClient.SetEx(ctx, key, value, duration).Err()
}
