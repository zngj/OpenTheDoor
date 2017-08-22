package redisx

import (
	"common/file"
	"github.com/go-redis/redis"
)

var Client *redis.Client

var config struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func init() {

	file.LoadJsonConfig("redis.json", &config)

	Client = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
}
