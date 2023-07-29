package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (r *redisRepo) GetValueByKey(ctx context.Context, key string) (value string, exists bool, err error) {
	cmd := r.redisClient.Get(ctx, key)
	value, err = cmd.Result()
	if err != nil {
		if err == redis.Nil {
			return "", false, nil
		}

		return "", false, err
	}

	exists = true
	return
}
