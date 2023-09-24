package cache

import (
	"os"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Cache *redis.Client
}

func New() *Cache {
	return &Cache{}
}

func (c *Cache) Configure() {
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")
	c.Cache = redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: REDIS_PASSWORD,
		DB:       0,
	})
}
