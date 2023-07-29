package config

import (
	"errors"
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Config struct {
	Redis    RedisConfig    `json:"redis_config"`
	DbConfig DatabaseConfig `json:"database_config"`
	Http     HttpConfig     `json:"http_config"`
	Jwt      JwtConfig      `json:"jwt"`
}

type RedisConfig struct {
	Url string `json:"url"`
}

type DatabaseConfig struct {
	Dsn string `json:"dsn"`
}

type HttpConfig struct {
	Port                  int    `json:"port"`
	TimeoutPerRequestInMs int    `json:"timeout_per_request_in_ms"`
	BasicAuthUname        string `json:"basic_auth_username"`
	BasicAuthPassword     string `json:"basic_auth_password"`
}

type JwtConfig struct {
	SecretKey           string `json:"secret_key"`
	ExpirationInSeconds int    `json:"expiration_in_seconds"`
}

func ParseConfig(cfgPath string) (cfg *Config, err error) {
	dataBytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return
	}

	cfg = &Config{}
	err = json.Unmarshal(dataBytes, cfg)
	if err != nil {
		return
	}

	if cfg.Jwt.SecretKey == "" {
		err = errors.New("jwt secret key config can't be empty")
		return
	}

	if cfg.Jwt.ExpirationInSeconds == 0 {
		cfg.Jwt.ExpirationInSeconds = 3600
		return
	}

	if cfg.Http.Port == 0 {
		cfg.Http.Port = 3000
		return
	}
	if cfg.Http.TimeoutPerRequestInMs == 0 {
		cfg.Http.TimeoutPerRequestInMs = 3_000
	}
	if cfg.Http.BasicAuthUname == "" {
		err = errors.New("http basic_auth_username config can't be empty")
		return
	}
	if cfg.Http.BasicAuthPassword == "" {
		err = errors.New("http basic_auth_password config can't be empty")
		return
	}

	if cfg.DbConfig.Dsn == "" {
		err = errors.New("database dsn config can't be empty")
		return
	}

	if cfg.Redis.Url == "" {
		err = errors.New("redis url config can't be empty")
		return
	}

	return
}
