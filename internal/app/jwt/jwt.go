package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kbiits/gikslab_assesment/internal/app/config"
	"github.com/kbiits/gikslab_assesment/internal/app/repository/redis"
)

type Jwt struct {
	cfg config.JwtConfig
	rdb redis.RedisRepository
}

var JwtAlg = jwt.SigningMethodHS256

func New(cfg config.JwtConfig, rdb redis.RedisRepository) *Jwt {
	return &Jwt{
		cfg: cfg,
		rdb: rdb,
	}
}

func (j *Jwt) IsValidMethod(t *jwt.Token) bool {
	return t.Method.Alg() == JwtAlg.Alg()
}

func (j *Jwt) CreateToken(claims Claims) (string, error) {
	now := time.Now()
	exp := now.Add(time.Second * time.Duration(j.cfg.ExpirationInSeconds))

	claims.ExpiredAt = exp.Unix()
	claims.IssuedAt = now.Unix()

	token := jwt.NewWithClaims(JwtAlg, claims.ToJwtMapClaims())
	return token.SignedString(j.GetSecretKey())
}

func (j *Jwt) GetSecretKey() []byte {
	return []byte(j.cfg.SecretKey)
}

const FormatRedisKeyBlacklisted = "BLACKLISTED_%v"

func (j *Jwt) CheckIsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	_, exists, err := j.rdb.GetValueByKey(ctx, fmt.Sprintf(FormatRedisKeyBlacklisted, token))
	if err != nil {
		return true, err
	}

	return exists, nil
}
