package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisRepo struct {
	redisClient *redis.Client
}

func NewRedisRepository(client *redis.Client) (RedisRepository, error) {
	if err := testRedisConnectionBeforeInitialize(client); err != nil {
		return &redisRepo{}, err
	}

	return &redisRepo{
		redisClient: client,
	}, nil
}

func testRedisConnectionBeforeInitialize(rdb *redis.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pingCmd := rdb.Ping(ctx)
	val, err := pingCmd.Result()
	if err != nil {
		return fmt.Errorf("failed ping redis, err : %v", err)
	}

	if val != "PONG" {
		return errors.New("unknown redis ping result")
	}

	return nil
}

type RedisRepository interface {
	GetValueByKey(ctx context.Context, key string) (string, bool, error)
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
}
